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
//  Created:   2025/12/21 23:49
//  Version:   v1.0.0

package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type reqBody struct {
	JSON         string   `json:"json"`
	Language     string   `json:"language"`
	Tags         []string `json:"tags"`
	StructName   string   `json:"struct_name"`
	PackageName  string   `json:"package_name"`
	MakePublic   bool     `json:"make_public"`
	AddComments  bool     `json:"add_comments"`
	UsePointer   bool     `json:"use_pointer"`
	TimeAsString bool     `json:"time_as_string"`
}

// HandleConvert API处理函数
func HandleConvert(c *gin.Context) {
	req := &reqBody{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "无效的请求格式",
		})
		return
	}

	// 这里应该调用转换器逻辑
	// 暂时返回示例
	exampleCode := ""
	switch req.Language {
	case "golang":
		exampleCode = generateGoExample(req)
	case "java":
		exampleCode = generateJavaExample(req)
	case "typescript":
		exampleCode = generateTypeScriptExample(req)
	case "python":
		exampleCode = generatePythonExample(req)
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "不支持的语言类型",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    exampleCode,
	})
}

// 生成Go示例代码
func generateGoExample(req *reqBody) string {
	tags := ""
	if len(req.Tags) > 0 {
		tagStr := make([]string, len(req.Tags))
		for i, tag := range req.Tags {
			tagStr[i] = fmt.Sprintf(`%s:"-"`, tag)
		}
		tags = " `" + strings.Join(tagStr, " ") + "`"
	}

	return fmt.Sprintf(`package %s

// %s 自动生成的Go结构体
type %s struct {
	ID        int64  %s
	Name      string %s
	Email     string %s
	CreatedAt string %s
	IsActive  bool   %s
}`,
		req.PackageName,
		req.StructName,
		req.StructName,
		tags, tags, tags, tags, tags)
}

// 生成Java示例代码
func generateJavaExample(req *reqBody) string {
	return fmt.Sprintf(`import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

@Data
public class %s {
    @JsonProperty("id")
    private Long id;
    
    @JsonProperty("name")
    private String name;
    
    @JsonProperty("email")
    private String email;
}`, req.StructName)
}

// 生成TypeScript示例代码
func generateTypeScriptExample(req *reqBody) string {
	return fmt.Sprintf(`interface %s {
    id: number;
    name: string;
    email: string;
    createdAt: string;
    isActive: boolean;
}`, req.StructName)
}

// 生成Python示例代码
func generatePythonExample(req *reqBody) string {
	return fmt.Sprintf(`from dataclasses import dataclass
from typing import Optional

@dataclass
class %s:
    id: int
    name: str
    email: str
    created_at: str
    is_active: bool`, req.StructName)
}

func HandleHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"time":    time.Now().Format(time.RFC3339),
		"service": "json-to-struct",
		"version": "1.0.0",
	})
}

func HandleGetLanguages(c *gin.Context) {
	languages := []gin.H{
		{"value": "golang", "label": "Go", "icon": "language-go"},
		{"value": "java", "label": "Java", "icon": "language-java"},
		{"value": "typescript", "label": "TypeScript", "icon": "language-typescript"},
		{"value": "python", "label": "Python", "icon": "language-python"},
	}
	c.JSON(http.StatusOK, languages)
}

func HandleGetTags(c *gin.Context) {
	tags := []gin.H{
		{"value": "json", "label": "JSON", "description": "JSON序列化标签"},
		{"value": "mapstructure", "label": "MapStructure", "description": "用于mapstructure库"},
		{"value": "gorm", "label": "GORM", "description": "GORM数据库标签"},
		{"value": "bson", "label": "BSON", "description": "MongoDB BSON标签"},
		{"value": "yaml", "label": "YAML", "description": "YAML序列化标签"},
		{"value": "xml", "label": "XML", "description": "XML序列化标签"},
		{"value": "validate", "label": "Validate", "description": "验证标签"},
		{"value": "form", "label": "Form", "description": "表单绑定标签"},
	}
	c.JSON(http.StatusOK, tags)
}

// HandleDemoImage 演示接口：返回图片数据
func HandleDemoImage(c *gin.Context) {
	// 创建一个简单的SVG图标作为示例
	svgData := `<?xml version="1.0" encoding="UTF-8"?>
<svg width="100" height="100" viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
  <circle cx="50" cy="50" r="40" fill="#3498db"/>
  <text x="50" y="55" text-anchor="middle" fill="white" font-family="Arial" font-size="16">
    JSON
  </text>
</svg>`

	c.Header("Content-Type", "image/svg+xml")
	c.String(http.StatusOK, svgData)
}

// HandleDemoJSON 演示接口：返回JSON数据
func HandleDemoJSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "这是一个JSON响应示例",
		"timestamp": time.Now().Unix(),
		"data": map[string]any{
			"id":    1,
			"name":  "示例数据",
			"items": []string{"item1", "item2", "item3"},
		},
	})
}

// HandleDemoText 演示接口：返回纯文本数据
func HandleDemoText(c *gin.Context) {
	c.Header("Content-Type", "text/plain; charset=utf-8")
	c.String(http.StatusOK, "这是一个纯文本响应示例\n时间: %s\n服务: JSON转结构体工具",
		time.Now().Format("2006-01-02 15:04:05"))
}
