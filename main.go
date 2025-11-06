package main

import (
	"encoding/json"
	"flag"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"prosearchmcp/constants"
	"prosearchmcp/tools/prosearchtool"
)

func main() {
	// Create a new MCP server
	secretID, secretKey := "", ""
	flag.StringVar(&secretID, "TENCENTCLOUD_SECRET_ID", "", "Secret ID")
	flag.StringVar(&secretKey, "TENCENTCLOUD_SECRET_KEY", "", "Secret Key")
	flag.Parse()
	s := server.NewMCPServer("prosearch", "1.0.0", server.WithToolCapabilities(true))
	tool := mcp.NewToolWithRawSchema("prosearch", constants.ToolDesc, json.RawMessage(constants.InputSchema))
	// Add tool handler
	handler := prosearchtool.NewProSearch(secretID, secretKey)
	s.AddTool(tool, handler.ToolHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		log.Printf("Server error: %v\n", err)
	}
}
