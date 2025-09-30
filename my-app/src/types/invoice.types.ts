export interface Invoice {
  ID: string;
  name: string;
  status: string;
  method: string;
  amount: number;
}

export interface CreateInvoiceRequest {
  name: string;
  status: string;
  method: string;
  amount: number;
}
