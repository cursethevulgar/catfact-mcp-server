package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"crypto/tls"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"Cat Fact MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	tool := mcp.NewTool("cat_fact",
		mcp.WithDescription("Provides a random cat fact summarized in a single sentence."),
	)

	s.AddTool(tool, catfactHandler)

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func catfactHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
    	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cat fact: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch cat fact: %s", resp.Status)
	}
	var result struct {
		Fact string `json:"fact"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode cat fact: %w", err)
	}

	return mcp.NewToolResultText(result.Fact), nil
}
