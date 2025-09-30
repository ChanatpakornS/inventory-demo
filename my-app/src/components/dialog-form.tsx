"use client";
import * as React from "react";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { createInvoice } from "@/app/api/invoices/route";

const formSchema = z.object({
  invoiceName: z.string().min(2, {
    message: "Name must be at least 2 characters.",
  }),
  status: z.string().min(2, {
    message: "Status must be at least 2 characters.",
  }),
  method: z.string().min(2, {
    message: "Method must be at least 2 characters.",
  }),
  amount: z.number().gt(0, {
    message: "Amount must be a positive number.",
  }),
});

export function DialogWithForm() {
  const [open, setOpen] = React.useState(false);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      invoiceName: "",
      status: "",
      method: "",
      amount: 0,
    },
  });

  const handleSubmit = async (values: z.infer<typeof formSchema>) => {
    // Handle form submission
    console.log("Form submitted");
    console.log(values);
    await createInvoice({
      name: values.invoiceName,
      status: values.status,
      method: values.method,
      amount: values.amount,
    });
    setOpen(false);
  };

  React.useEffect(() => {
    if (!open) {
      form.reset();
    }
  }, [open, form]);

  return (
    <div className="w-full p-6 flex justify-center">
      <Form {...form}>
        <Dialog open={open} onOpenChange={setOpen}>
          <DialogTrigger asChild>
            <Button variant="outline">Open Form Dialog</Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[425px]">
            <form
              onSubmit={form.handleSubmit(handleSubmit)}
              className="space-y-4"
            >
              <DialogHeader>
                <DialogTitle>Add the Items!</DialogTitle>
                <DialogDescription>
                  Make change to your invoices here
                </DialogDescription>
              </DialogHeader>
              {/* FORM STARTED HERE */}
              <FormField
                control={form.control}
                name="invoiceName"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Invoice name</FormLabel>
                    <FormControl>
                      <Input placeholder="shadcn" {...field} />
                    </FormControl>
                    <FormDescription>
                      This is your invoice display name.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="status"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Status</FormLabel>
                    <FormControl>
                      <Input placeholder="shadcn" {...field} />
                    </FormControl>
                    <FormDescription>
                      This is your invoice's status.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="method"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Method</FormLabel>
                    <FormControl>
                      <Input placeholder="shadcn" {...field} />
                    </FormControl>
                    <FormDescription>
                      This is your payment method.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="amount"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Amount</FormLabel>
                    <FormControl>
                      <Input
                        type="number"
                        placeholder="0.00"
                        {...field}
                        onChange={(e) =>
                          field.onChange(parseFloat(e.target.value) || 0)
                        }
                      />
                    </FormControl>
                    <FormDescription>
                      This is your amount in invoice.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* FORM ENDED HERE */}
              <DialogFooter>
                <DialogClose asChild>
                  <Button type="button" variant="outline">
                    Cancel
                  </Button>
                </DialogClose>
                <Button type="submit">Submit</Button>
              </DialogFooter>
            </form>
          </DialogContent>
        </Dialog>
      </Form>
    </div>
  );
}
