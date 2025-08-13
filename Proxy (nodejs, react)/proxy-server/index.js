const express = require("express");
const { createProxyMiddleware } = require("http-proxy-middleware");

const app = express();

// ✅ Khai báo đầu tiên
app.use((req, res, next) => {
  res.setHeader(
    "Content-Security-Policy",
    [
      "default-src 'self'",
      "script-src 'self' 'unsafe-inline'",
      "style-src 'self' 'unsafe-inline'",
      "img-src 'self' data:",
      // ⚠️ Cho phép websocket đến domain React dev server đang chạy
      "connect-src 'self' ws://localhost:3000 ws://gmo-h110m-h.tail04954f.ts.net:3000",
    ].join("; ")
  );
  next();
});

app.get("/favicon.ico", (req, res) => res.status(204).end());

// Proxy cho API
app.use(
  "/api",
  createProxyMiddleware({
    target: "http://localhost:5000", // backend port
    changeOrigin: true,
  })
);

// Proxy cho frontend (React build)
app.use(
  "/",
  createProxyMiddleware({
    target: "http://localhost:3000", // frontend port (dev)
    changeOrigin: true,
  })
);

app.listen(4000, () => {
  console.log("Proxy server running on http://localhost:4000");
});
