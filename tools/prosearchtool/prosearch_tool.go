package prosearchtool

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"prosearchmcp/constants"
	"prosearchmcp/prosearch"
	"prosearchmcp/tools"
)

// ProSearch is the prosearch tool.
type ProSearch struct {
	searchClient prosearch.Search
	secretID     string
	secretKey    string
}

// NewProSearch returns a new ProSearch tool.
func NewProSearch(secretID string, secretKey string) *ProSearch {
	return &ProSearch{secretID: secretID, secretKey: secretKey, searchClient: prosearch.NewSearch(constants.ProSearch)}
}

// ToolHandler is the handler for the prosearch tool.
func (ps *ProSearch) ToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	input := request.GetArguments()
	query := tools.GetString("query", input)
	mode := tools.GetFloat64("mode", input)
	site := tools.GetString("site", input)
	fromTime := tools.GetString("from_time", input)
	toTime := tools.GetString("to_time", input)
	ft := tools.StringToTimestamp(fromTime)
	tt := tools.StringToTimestamp(toTime)
	req := &prosearch.CloudAPIRequest{Query: query, Mode: int(mode), Site: site,
		SecretKey: ps.secretKey, SecretID: ps.secretID, FromMcp: true}
	if ft > -1 {
		req.FromTime = ft
	}
	if tt > -1 {
		req.ToTime = tt
	}
	rsp, err := ps.searchClient.CloudApi(ctx, req)
	if err != nil {
		return tools.DoFailResponse(ctx, request, req, err)
	}
	return tools.DoResponse(ctx, request, rsp)
}
