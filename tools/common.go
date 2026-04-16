package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/mark3labs/mcp-go/mcp"

	"prosearchmcp/prosearch"
)

// GetString returns the string value for the given name.
func GetString(name string, input map[string]any) string {
	var value string
	if v, ok := input[name]; ok {
		value = v.(string)
	}
	return value
}

// GetFloat64 从 input 中 value 解析 float64类型，注意，json 中数字都是 float64，只能用这个
func GetFloat64(name string, input map[string]any) float64 {
	var value float64
	if v, ok := input[name]; ok {
		switch vv := v.(type) {
		case int64:
			return float64(vv)
		case float64:
			// 处理JSON解析出的浮点数整型值
			return vv
		case int:
			return float64(vv)
		}
	}
	return value
}

// GetBool returns the string value for the given name.
func GetBool(name string, input map[string]any) bool {
	var value bool
	if v, ok := input[name]; ok {
		if vv, vok := v.(bool); vok {
			value = vv
		}
	}
	return value
}

// GetStringArr 从 input 中 value 解析 []string类型，注意：json 中数组会解析成 []interface{}
func GetStringArr(name string, input map[string]any) []string {
	value := []string{}
	if v, ok := input[name]; ok {
		switch vv := v.(type) {
		case []interface{}:
			for _, one := range vv {
				if oneStr, ok := one.(string); ok {
					value = append(value, oneStr)
				}
			}
			return value
		case []string:
			return vv
		}
	}
	return value
}

// StringToTimestamp 根据输入的时间字符串转化为时间戳
// timeStr length为10且为数字, 直接返回
// tmeStr为YYYY-MM-DD或者YYYY-MM-DD HH:MM:SS, 转化为时间戳
// 其他返回-1
func StringToTimestamp(timeStr string) int64 {
	if len(timeStr) == 10 {
		// 如果时间戳字符串的长度为10，且其格式为Unix时间戳（以秒为单位）
		if timestamp, err := strconv.ParseInt(timeStr, 10, 64); err == nil {
			return timestamp
		}
	}
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return -1
	}
	var layout string
	if len(timeStr) == 10 {
		// 如果时间戳字符串的长度为10，且其格式为"YYYY-MM-DD"
		layout = "2006-01-02"
	} else if len(timeStr) == 19 {
		// 如果时间戳字符串的长度为19，且其格式为"YYYY-MM-DD HH:MM:SS"
		layout = "2006-01-02 15:04:05"
	} else {
		return -1
	}

	// 将timeStr转化为time.Time
	parsedTime, err := time.ParseInLocation(layout, timeStr, location)
	if err != nil {
		return -1
	}
	// 转换为秒
	timeStamp := parsedTime.Local().Unix()

	return timeStamp
}

// DoFailResponse fail response process
func DoFailResponse(ctx context.Context, req mcp.CallToolRequest, creq *prosearch.CloudAPIRequest,
	err error) (*mcp.CallToolResult, error) {
	if err == nil {
		return nil, nil
	}
	jsonData, _ := json.Marshal(req.Params.Arguments)
	log.Printf("toolanme:%v, input:%s, DoFailResponse reason:%v", req.Params.Name, string(jsonData), err)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("Error: tool name=%s, input=%s, fail upstream request", req.Params.Name,
					string(jsonData)),
			},
		},
		IsError: true,
	}, nil
}

// DoResponse success response process
func DoResponse(ctx context.Context, req mcp.CallToolRequest, rspStr string) (*mcp.CallToolResult, error) {
	input := req.Params.Arguments
	jsonData, _ := json.Marshal(input)
	log.Printf("toolanme:%v, input:%s, DoResponse rspStr:%v", req.Params.Name, string(jsonData), rspStr)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Type: "text",
				Text: rspStr,
			},
		},
		IsError: false,
	}, nil
}
