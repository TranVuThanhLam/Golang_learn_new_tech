package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var backendServers = []string{
	"http://localhost:8080",
	"http://localhost:8081",
}

func main() {
	r := gin.Default()

	// Proxy API requests
	r.Any("/api/*path", func(c *gin.Context) {
		target := pickBackend()
		fullURL := target + c.Param("path")

		req, err := http.NewRequest(c.Request.Method, fullURL, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Tạo request thất bại"})
			return
		}
		req.Header = c.Request.Header

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Không thể kết nối backend"})
			return
		}
		defer resp.Body.Close()

		for k, v := range resp.Header {
			c.Header(k, v[0])
		}
		c.Status(resp.StatusCode)
		io.Copy(c.Writer, resp.Body)
	})

	// Redirect frontend (tùy chọn)
	r.NoRoute(func(c *gin.Context) {
		target := "http://localhost:3000" + c.Request.URL.Path
		c.Redirect(http.StatusTemporaryRedirect, target)
	})

	log.Println("Proxy server đang chạy tại http://localhost:4000")
	r.Run(":4000")
}

func pickBackend() string {
	rand.Seed(time.Now().UnixNano())
	return backendServers[rand.Intn(len(backendServers))]
}
