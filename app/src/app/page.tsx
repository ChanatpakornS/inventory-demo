import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { DialogWithForm } from "@/components/dialog-form";
import { getAllInvoices } from "./api/invoices/route";

export default async function Home() {
  const invoiceList = await getAllInvoices();

  return (
    <main className="flex min-h-screen flex-col items-center p-24">
      <h1 className="text-2xl font-bold">Invoices</h1>
      <h2>The invoice management</h2>
      <DialogWithForm />
      <Table className="size-full mt-4 border">
        <TableCaption>A list of your recent invoices.</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">ID</TableHead>
            <TableHead>Name</TableHead>
            <TableHead>Status</TableHead>
            <TableHead>Method</TableHead>
            <TableHead className="text-right">Amount</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {invoiceList.map((invoice) => (
            <TableRow key={invoice.ID}>
              <TableCell className="font-medium">{invoice.ID}</TableCell>
              <TableCell>{invoice.name}</TableCell>
              <TableCell>{invoice.status}</TableCell>
              <TableCell>{invoice.method}</TableCell>
              <TableCell className="text-right">${invoice.amount}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </main>
  );
}
