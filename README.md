# Invoice playground

TBH, I don't know what is invoices are, just used the words from table from shadCN and still develop until now!

> [!IMPORTANT]
> Why not turbo repo? IDK what it is and how to use it. K?

# Service

This repository consisted of 
 - server gRPC (connected with the MCP-server)
 - server REST (connected to the app interface)
 - app (stupid table only for 'GET', 'POST' method)
 - grpc-proto (collection of proto)
 - mcp-server (connected to Claude Desktop or etc.)

# Demo
 1. Installing [Claude Desktop](https://claude.ai/download)
 2. Init PostgreSQL DB from docker in `server-REST` --> `make compose`
 3. Activate all server (app `NextJS`, server-gRPC - `Go`, server-REST - `Go` , mcp-server - `Python + uv`)
 4. Configure the following to `claude_desktop_config`
 ```json
{
  "mcpServers": {
    "invoice-mcp-gateway": {
      "command": "npx",
      "args": ["mcp-remote", "http://127.0.0.1:9000/mcp"],
      "transport": "sse"
    }
  }
}
 ```
 5. Enjoy!

> [!NOTE] 
> Thanks [Tin's playground](https://github.com/tin2003tin/mcp-playground)
