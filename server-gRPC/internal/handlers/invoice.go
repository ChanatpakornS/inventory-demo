package handlers

import (
	"context"
	"errors"
	"strconv"

	invoicegRPC "github.com/ChanatpakornS/inventory-demo/gRPC/grpc-proto/invoice"
	model "github.com/ChanatpakornS/inventory-demo/gRPC/internal/models"
	"gorm.io/gorm"
)

type InvoiceHandler interface {
	GetAllInvoices(ctx context.Context, req *invoicegRPC.GetAllInvoicesRequest) (*invoicegRPC.GetAllInvoicesResponse, error)
	GetInvoiceByID(ctx context.Context, req *invoicegRPC.GetInvoiceRequest) (*invoicegRPC.GetInvoiceRequest, error)
	CreateInvoice(ctx context.Context, req *invoicegRPC.CreateInvoiceRequest) (*invoicegRPC.CreateInvoiceResponse, error)
	UpdateInvoice(ctx context.Context, req *invoicegRPC.UpdateInvoiceRequest) (*invoicegRPC.UpdateInvoiceResponse, error)
	DeleteInvoice(ctx context.Context, req *invoicegRPC.DeleteInvoiceRequest) (*invoicegRPC.DeleteInvoiceResponse, error)
}

type invoiceHandler struct {
	invoicegRPC.UnimplementedInvoiceServiceServer
	db *gorm.DB
}

func NewInvoiceHandler(db *gorm.DB) *invoiceHandler {
	return &invoiceHandler{
		db: db,
	}
}

func (i *invoiceHandler) GetAllInvoices(ctx context.Context, req *invoicegRPC.GetAllInvoicesRequest) (*invoicegRPC.GetAllInvoicesResponse, error) {
	var Invoices []model.Invoice

	result := i.db.Find(&Invoices)
	if result.Error != nil {
		return nil, errors.New("Fail to return invoices: " + result.Error.Error())
	}
	var invoiceList []*invoicegRPC.Invoice
	for _, invoice := range Invoices {
		invoiceList = append(invoiceList, &invoicegRPC.Invoice{
			Id:     strconv.FormatUint(uint64(invoice.ID), 10),
			Name:   invoice.Name,
			Status: invoice.Status,
			Method: invoice.Method,
			Amount: invoice.Amount,
		})
	}

	return &invoicegRPC.GetAllInvoicesResponse{
		Invoices: invoiceList,
	}, nil
}

func (i *invoiceHandler) GetInvoiceByID(ctx context.Context, req *invoicegRPC.GetInvoiceRequest) (*invoicegRPC.GetInvoiceResponse, error) {
	var Invoice model.Invoice

	result := i.db.First(&Invoice, req.Id)
	if result.Error != nil {
		return nil, errors.New("Fail to return invoice: " + result.Error.Error())
	}

	return &invoicegRPC.GetInvoiceResponse{
		Invoice: &invoicegRPC.Invoice{
			Id:     strconv.FormatUint(uint64(Invoice.ID), 10),
			Name:   Invoice.Name,
			Status: Invoice.Status,
			Method: Invoice.Method,
			Amount: Invoice.Amount,
		},
	}, nil
}

func (i *invoiceHandler) CreateInvoice(ctx context.Context, req *invoicegRPC.CreateInvoiceRequest) (*invoicegRPC.CreateInvoiceResponse, error) {
	var Invoice model.Invoice
	Invoice.Name = req.Name
	Invoice.Status = req.Status
	Invoice.Method = req.Method
	Invoice.Amount = req.Amount

	result := i.db.Create(&Invoice)
	if result.Error != nil {
		return nil, errors.New("Fail to create invoice: " + result.Error.Error())
	}

	return &invoicegRPC.CreateInvoiceResponse{
		Invoice: &invoicegRPC.Invoice{
			Id:     strconv.FormatUint(uint64(Invoice.ID), 10),
			Name:   Invoice.Name,
			Status: Invoice.Status,
			Method: Invoice.Method,
			Amount: Invoice.Amount,
		},
	}, nil
}

func (i *invoiceHandler) UpdateInvoice(ctx context.Context, req *invoicegRPC.UpdateInvoiceRequest) (*invoicegRPC.UpdateInvoiceResponse, error) {
	var Invoice model.Invoice

	result := i.db.First(&Invoice, req.Id)
	if result.Error != nil {
		return nil, errors.New("Fail to find invoice: " + result.Error.Error())
	}

	if req.Name != nil {
		Invoice.Name = req.GetName()
	}
	if req.Status != nil {
		Invoice.Status = req.GetStatus()
	}
	if req.Method != nil {
		Invoice.Method = req.GetMethod()
	}
	if req.Amount != nil {
		Invoice.Amount = req.GetAmount()
	}

	saveResult := i.db.Save(&Invoice)
	if saveResult.Error != nil {
		return nil, errors.New("Fail to update invoice: " + saveResult.Error.Error())
	}

	return &invoicegRPC.UpdateInvoiceResponse{
		Invoice: &invoicegRPC.Invoice{
			Id:     strconv.FormatUint(uint64(Invoice.ID), 10),
			Name:   Invoice.Name,
			Status: Invoice.Status,
			Method: Invoice.Method,
			Amount: Invoice.Amount,
		},
	}, nil
}

func (i *invoiceHandler) DeleteInvoice(ctx context.Context, req *invoicegRPC.DeleteInvoiceRequest) (*invoicegRPC.DeleteInvoiceResponse, error) {
	var Invoice model.Invoice

	result := i.db.First(&Invoice, req.Id)
	if result.Error != nil {
		return nil, errors.New("Fail to find invoice: " + result.Error.Error())
	}

	deleteResult := i.db.Delete(&Invoice)
	if deleteResult.Error != nil {
		return nil, errors.New("Fail to delete invoice: " + deleteResult.Error.Error())
	}

	return &invoicegRPC.DeleteInvoiceResponse{
		Invoice: &invoicegRPC.Invoice{
			Id:     strconv.FormatUint(uint64(Invoice.ID), 10),
			Name:   Invoice.Name,
			Status: Invoice.Status,
			Method: Invoice.Method,
			Amount: Invoice.Amount,
		},
	}, nil
}
