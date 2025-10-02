import asyncio
from typing import Any
from fastmcp import FastMCP
from starlette.requests import Request
from starlette.responses import PlainTextResponse, HTMLResponse
import grpc
from grpc_proto import invoice_pb2
from grpc_proto import invoice_pb2_grpc

# Create MCP instance
mcp = FastMCP("invoice-mcp-gateway")

# ---------------------------
# gRPC Stub Helper
# ---------------------------
async def get_grpc_stub():
    channel = grpc.aio.insecure_channel("localhost:50051")  # Adjust host/port if needed
    stub = invoice_pb2_grpc.InvoiceServiceStub(channel)
    return stub

# ---------------------------
# Tools
# ---------------------------

@mcp.tool()
async def get_all_invoices() -> dict[str, Any]:
    """
    Get all invoices from database via gRPC.
    """
    stub = await get_grpc_stub()
    resp = await stub.GetAllInvoices(invoice_pb2.GetAllInvoicesRequest())
    
    # Convert the repeated invoice field to a list of dicts
    invoices = []
    for invoice in resp.invoices:
        invoices.append({
            "id": invoice.id,
            "name": invoice.name,
            "status": invoice.status,
            "method": invoice.method,
            "amount": invoice.amount
        })
    
    return {
        "invoices": invoices,
        "count": len(invoices),
        "next_suggestion": "Try fetching a specific invoice by ID using get_invoice_id(id='your_invoice_id')."
    }

@mcp.tool()
async def get_invoice_id(id: str) -> dict[str, Any]:
    """
    Get a specific invoice by ID via gRPC.
    """
    stub = await get_grpc_stub()
    resp = await stub.GetInvoiceByID(invoice_pb2.GetInvoiceRequest(id=id))
    return {
        "id": resp.invoice.id,
        "name": resp.invoice.name,
        "status": resp.invoice.status,
        "method": resp.invoice.method,
        "amount": resp.invoice.amount,
        "next_suggestion": "You can update this invoice using update_invoice or delete it using delete_invoice."
    }


@mcp.tool()
async def create_invoice(name: str, status: str, method: str, amount: float
) -> dict[str, Any]:
    """
    Create a new invoice via gRPC.
    """
    stub = await get_grpc_stub()
    resp = await stub.CreateInvoice(invoice_pb2.CreateInvoiceRequest(
        name=name,
        status=status,
        method=method,
        amount=amount
    ))
    return {
        "id": resp.invoice.id,
        "name": resp.invoice.name,
        "status": resp.invoice.status,
        "method": resp.invoice.method,
        "amount": resp.invoice.amount,
        "next_suggestion": f"Invoice created with ID {resp.invoice.id}. You can fetch it using get_invoice_id(id='{resp.invoice.id}')."
    }

@mcp.tool()
async def update_invoice(id: str, name: str = "", status: str = "", method: str = "", amount: float = 0.0
) -> dict[str, Any]:
    """
    Update an existing invoice via gRPC.
    """
    stub = await get_grpc_stub()
    resp = await stub.UpdateInvoice(invoice_pb2.UpdateInvoiceRequest(
        id=id,
        name=name,
        status=status,
        method=method,
        amount=amount
    ))
    return {
        "id": resp.invoice.id,
        "name": resp.invoice.name,
        "status": resp.invoice.status,
        "method": resp.invoice.method,
        "amount": resp.invoice.amount,
        "next_suggestion": f"Invoice with ID {resp.invoice.id} updated. You can fetch it using get_invoice_id(id='{resp.invoice.id}')."
    }

@mcp.tool()
async def delete_invoice(id: str) -> dict[str, Any]:
    """
    Delete an invoice by ID via gRPC.
    """
    stub = await get_grpc_stub()
    resp = await stub.DeleteInvoice(invoice_pb2.DeleteInvoiceRequest(id=id))
    return {
        "id": resp.invoice.id,
        "name": resp.invoice.name,
        "status": resp.invoice.status,
        "method": resp.invoice.method,
        "amount": resp.invoice.amount,
        "next_suggestion": "You can create a new invoice using create_invoice."
    }


# ---------------------------
# Root HTML page 
# ---------------------------
@mcp.custom_route("/", methods=["GET"])
async def root(request: Request) -> HTMLResponse:
    """Return a tiny HTML dashboard for the demo service."""
    html = """
    <!doctype html>
    <html>
    <head>
      <meta charset="utf-8"/>
      <title>Invoice MCP Demo</title>
      <style>
        body { font-family: Arial, sans-serif; max-width: 900px; margin: 40px auto; line-height:1.6; padding: 0 20px; }
        h1 { color: #0b5; }
        a { color: #0366d6; }
        code { background:#f6f8fa; padding:2px 6px; border-radius:4px; }
        .card { padding:12px; border:1px solid #eee; border-radius:8px; margin-bottom:12px; box-shadow: 0 1px 2px rgba(0,0,0,.02); }
      </style>
    </head>
    <body>
      <h1>Invoice — Invoice Gateway (Demo)</h1>
      <p>This is a lightweight demo of the MCP tools.</p>

      <div class="card">
        <strong>Health check</strong><br/>
        <a href="/health" target="_blank">GET /health</a> — returns <code>OK</code>
      </div>

      <div class="card">
        <strong>Notes</strong>
        <p>This page is intentionally simple. Use the MCP client (`fastmcp`) to call the tools declared in this process.</p>
      </div>

      <footer>
        <small>Running on <code>127.0.0.1:9000</code> if started with the default <code>main()</code>.</small>
      </footer>
    </body>
    </html>
    """
    return HTMLResponse(content=html)

# ---------------------------
# Health Check
# ---------------------------
@mcp.custom_route("/health", methods=["GET"])
async def health_check(request: Request) -> PlainTextResponse:
    return PlainTextResponse("OK")


# ---------------------------
# Middlewares
# ---------------------------
# Add middleware to log requests
@mcp.custom_route("/mcp", methods=["POST"])
async def mcp_debug(request: Request):
    """Debug middleware to log incoming requests"""
    body = await request.body()
    logger.debug(f"Received request body: {body.decode('utf-8', errors='ignore')}")
    logger.debug(f"Content-Type: {request.headers.get('content-type')}")
    # Let FastMCP handle the request normally
    return None

# ---------------------------
# Main
# ---------------------------
def main():
    print("Starting Invoice MCP Gateway (gRPC mode)...")
    mcp.run(transport="http", host="127.0.0.1", port=9000, path="/mcp")

if __name__ == "__main__":
    main()