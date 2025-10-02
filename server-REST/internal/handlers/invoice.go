package handlers

import (
	model "github.com/ChanatpakornS/inventory-demo/REST/internal/models"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type InvoiceHandler interface {
	GetAllInvoices(ctx fiber.Ctx, db *gorm.DB) error
	GetInvoiceByID(ctx fiber.Ctx, db *gorm.DB) error
	CreateInvoice(ctx fiber.Ctx, db *gorm.DB) error
	UpdateInvoice(ctx fiber.Ctx, db *gorm.DB) error
	DeleteInvoice(ctx fiber.Ctx, db *gorm.DB) error
}

type invoiceHandler struct{}

func NewInvoiceHandler() *invoiceHandler {
	return &invoiceHandler{}
}

func (i *invoiceHandler) GetAllInvoices(ctx fiber.Ctx, db *gorm.DB) error {
	var Invoices []model.Invoice

	result := db.Find(&Invoices)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Invoice not found",
		})
	}

	ctx.JSON(Invoices)

	return ctx.Status(fiber.StatusOK).JSON(Invoices)
}
func (i *invoiceHandler) GetInvoiceByID(ctx fiber.Ctx, db *gorm.DB) error {
	var (
		Invoice model.Invoice
		param   = struct {
			ID uint `uri:"id"`
		}{}
	)

	if err := ctx.Bind().URI(&param); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID parameter",
		})
	}
	id := param.ID

	result := db.First(&Invoice, id)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Invoice not found",
		})
	}

	ctx.JSON(Invoice)

	return ctx.Status(fiber.StatusOK).JSON(Invoice)
}
func (i *invoiceHandler) CreateInvoice(ctx fiber.Ctx, db *gorm.DB) error {
	var Invoice model.Invoice
	if err := ctx.Bind().Body(&Invoice); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := db.Create(&Invoice).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create Invoice",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(Invoice)
}
func (i *invoiceHandler) UpdateInvoice(ctx fiber.Ctx, db *gorm.DB) error {
	var (
		Invoice model.Invoice
		param   = struct {
			ID uint `uri:"id"`
		}{}
	)
	if err := ctx.Bind().URI(&param); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID parameter",
		})
	}
	id := param.ID

	if err := ctx.Bind().Body(&Invoice); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := db.Where("id = ?", id).Updates(&Invoice).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update Invoice",
		})
	}

	result := db.First(&Invoice, id)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Invoice not found",
		})
	}

	ctx.JSON(Invoice)

	return ctx.Status(fiber.StatusOK).JSON(Invoice)
}
func (i *invoiceHandler) DeleteInvoice(ctx fiber.Ctx, db *gorm.DB) error {
	var (
		Invoice model.Invoice
		param   = struct {
			ID uint `uri:"id"`
		}{}
	)

	if err := ctx.Bind().URI(&param); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID parameter",
		})
	}
	id := param.ID

	result := db.First(&Invoice, id)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Invoice not found",
		})
	}

	if err := db.Delete(&Invoice, id).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not delete Invoice",
		})
	}

	ctx.JSON(result)

	return ctx.Status(fiber.StatusOK).JSON(result)
}
