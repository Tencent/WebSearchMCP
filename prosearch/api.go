// Package prosearch 增强搜索
package prosearch

import (
	"context"

	"prosearchmcp/constants"
)

// Search 增强搜索
type Search interface {
	// CloudApi 云api 联网搜索
	CloudApi(ctx context.Context, req *CloudAPIRequest) (string, error)
}

// NewSearch 创建gaokao实例
func NewSearch(toolName constants.ToolName) Search {
	return newSearchImpl(toolName)
}
