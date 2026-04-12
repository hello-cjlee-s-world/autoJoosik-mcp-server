# auto-joosik-market-data-fetcher
   
# Build
# linux build / amd64
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -ldflags="-s -w" -trimpath -o .\build\autoJoosik-mcp-server .\cmd\main.go

# window build / amd64
$env:GOOS="windows"; $env:GOARCH="amd64"; $env:CGO_ENABLED="0"; go build -ldflags="-s -w" -trimpath -o .\build\autoJoosik-mcp-server.exe .\cmd\main.go 


#Claude 코드 설정
{
    "mcpServers": {
        "stock-server": {
        "command": "D:\\Project\\autoJoosik-mcp-server\\build\\autoJoosik-mcp-server.exe",
        "args": []
        }
    }
}

#codex cli 코드 설정(config.toml)
[mcp_servers.stock]
command = "D:\\Project\\autoJoosik-mcp-server\\build\\autoJoosik-mcp-server.exe"