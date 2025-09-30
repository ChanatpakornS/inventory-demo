import { CreateInvoiceRequest, Invoice } from "@/types/invoice.types";

export async function getAllInvoices() {
  const res = await fetch("http://localhost:8080/invoices", {
    method: "GET",
    cache: "no-store",
  });
  const invoiceList = await res.json();
  return invoiceList as Invoice[];
}

export async function createInvoice(data: CreateInvoiceRequest) {
  const res = await fetch("http://localhost:8080/invoices", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  const newInvoice = await res.json();
  return newInvoice as Invoice;
}
