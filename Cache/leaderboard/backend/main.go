package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"
)

var db *sql.DB
var rdb *redis.Client

func main() {
	// Kết nối SQLite
	var err error
	db, err = sql.Open("sqlite3", "./leaderboard.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Tạo bảng nếu chưa tồn tại
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS scores (
			user_id TEXT PRIMARY KEY,
			score INTEGER NOT NULL
		)
	`)
	if err != nil {
		panic(err)
	}

	// Kết nối Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis chạy local
		Password: "",               // no password
		DB:       0,                // use default DB
	})

	// Khởi tạo Gin
	r := gin.Default()
	// --- Cấu hình CORS ở đây ---
	// Cho phép tất cả các tên miền truy cập trong quá trình phát triển.
	// Trong production, bạn nên chỉ định tên miền cụ thể.
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	// API Endpoints
	r.POST("/api/scores", updateScore)
	r.GET("/api/leaderboard", getLeaderboard)

	// Chạy server
	r.Run(":8080")
}

func updateScore(c *gin.Context) {
	var data struct {
		UserID string `json:"user_id"`
		Score  string `json:"score"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Lưu vào SQLite
	_, err := db.Exec(`
		INSERT INTO scores (user_id, score) 
		VALUES (?, ?) 
		ON CONFLICT(user_id) DO UPDATE SET score = excluded.score
	`, data.UserID, data.Score)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 2. Cập nhật Redis
	score, err := strconv.ParseFloat(data.Score, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}
	err = rdb.ZAdd(c, "leaderboard", redis.Z{
		Score:  score,
		Member: data.UserID,
	}).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func getLeaderboard(c *gin.Context) {
	top := c.DefaultQuery("top", "10")
	n, _ := strconv.Atoi(top)

	// Lấy từ Redis
	result, err := rdb.ZRevRangeWithScores(c, "leaderboard", 0, int64(n-1)).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Format response
	leaderboard := make([]map[string]interface{}, len(result))
	for i, item := range result {
		leaderboard[i] = map[string]interface{}{
			"user_id": item.Member,
			"score":   item.Score,
		}
	}

	c.JSON(http.StatusOK, leaderboard)
}
