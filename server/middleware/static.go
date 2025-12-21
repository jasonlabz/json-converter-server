// Package controller
//
//   _ __ ___   __ _ _ __  _   _| |_
//  | '_ ` _ \ / _` | '_ \| | | | __|
//  | | | | | | (_| | | | | |_| | |_
//  |_| |_| |_|\__,_|_| |_|\__,_|\__|
//
//  Buddha bless, no bugs forever!
//
//  Author:    lucas
//  Email:     1783022886@qq.com
//  Created:   2025/12/21 23:19
//  Version:   v1.0.0

package middleware

import (
	"embed"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/jasonlabz/json-converter-server/common/consts"
	"github.com/jasonlabz/json-converter-server/common/ginx"
)

// 嵌入所有前端文件
var frontendFS embed.FS

// 文件类型映射
var mimeTypes = map[string]string{
	".html":  "text/html; charset=utf-8",
	".htm":   "text/html; charset=utf-8",
	".css":   "text/css; charset=utf-8",
	".js":    "application/javascript; charset=utf-8",
	".json":  "application/json; charset=utf-8",
	".png":   "image/png",
	".jpg":   "image/jpeg",
	".jpeg":  "image/jpeg",
	".gif":   "image/gif",
	".svg":   "image/svg+xml",
	".ico":   "image/x-icon",
	".woff":  "font/woff",
	".woff2": "font/woff2",
	".ttf":   "font/ttf",
	".eot":   "application/vnd.ms-fontobject",
	".otf":   "font/otf",
	".txt":   "text/plain; charset=utf-8",
	".xml":   "application/xml",
	".pdf":   "application/pdf",
	".zip":   "application/zip",
	".tar":   "application/x-tar",
	".gz":    "application/gzip",
	".mp3":   "audio/mpeg",
	".mp4":   "video/mp4",
	".webm":  "video/webm",
	".webp":  "image/webp",
	".wasm":  "application/wasm",
}

// 获取文件的MIME类型
func getContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	if mime, ok := mimeTypes[ext]; ok {
		return mime
	}
	// 默认返回二进制流
	return "application/octet-stream"
}

// HandleStaticFile 处理静态文件请求
func HandleStaticFile(c *gin.Context) {
	// 获取请求路径
	path := c.Request.URL.Path

	// API请求直接返回
	if strings.HasPrefix(path, "/api/") {
		c.Next()
		return
	}

	// 默认首页
	if path == "/" || path == "" {
		path = "/index.html"
	}

	// 移除开头的斜杠
	filename := strings.TrimPrefix(path, "/")

	switch filename {
	case "index":
		filename = "index.html"
	case "favicon":
		filename = "favicon.ico"
	case "style":
		filename = "style.css"
	case "app":
		filename = "app.js"
	case "logo":
		filename = "logo.png"
	case "icon-go":
		filename = "icon-go.png"
	case "icon-java":
		filename = "icon-java.png"
	case "icon-ts":
		filename = "icon-ts.png"
	case "icon-python":
		filename = "icon-python.png"
	}
	// 设置Content-Type
	contentType := getContentType(filename)
	c.Header("Content-Type", contentType)

	config := &ginx.FileDownloadConfig{
		Filename:    filename,
		Preview:     true,
		ContentType: contentType,
		Header:      make(map[string]string),
	}
	// 读取嵌入的文件
	data, err := frontendFS.ReadFile("static/" + filename)
	if err != nil {
		// 文件不存在，尝试读取index.html
		if filename != "index.html" {
			data, err = frontendFS.ReadFile("static/index.html")
			config.Content = data
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
				return
			}
			ginx.FileResult(c, consts.APIVersionV1, config)
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	config.Content = data
	// 设置缓存头（静态文件可以缓存）
	if strings.HasSuffix(filename, ".html") {
		config.Header["Cache-Control"] = "no-cache, no-store, must-revalidate"
	} else {
		// CSS、JS、图片等静态资源可以缓存
		config.Header["Cache-Control"] = "public, max-age=31536000"
		config.Header["Expires"] = time.Now().Add(365 * 24 * time.Hour).Format(http.TimeFormat)
	}
	ginx.FileResult(c, consts.APIVersionV1, config)
}
