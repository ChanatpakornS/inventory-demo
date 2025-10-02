package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ChanatpakornS/inventory-demo/REST/internal/handlers"
	model "github.com/ChanatpakornS/inventory-demo/REST/internal/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "idb"
	password = "idb"
	dbname   = "idb"
)

func main() {
	app := fiber.New()
	db := setUpDatabase(host, port, user, password, dbname)

	invoiceHandler := handlers.NewInvoiceHandler()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin, Content-Type, Accept"},
	}))
	app.Get(healthcheck.LivenessEndpoint, healthcheck.New())
	app.Get("/invoices", func(c fiber.Ctx) error {
		return invoiceHandler.GetAllInvoices(c, db)
	})
	app.Get("/invoices/:id", func(c fiber.Ctx) error {
		return invoiceHandler.GetInvoiceByID(c, db)
	})
	app.Post("/invoices", func(c fiber.Ctx) error {
		return invoiceHandler.CreateInvoice(c, db)
	})
	app.Put("/invoices/:id", func(c fiber.Ctx) error {
		return invoiceHandler.UpdateInvoice(c, db)
	})
	app.Delete("/invoices/:id", func(c fiber.Ctx) error {
		return invoiceHandler.DeleteInvoice(c, db)
	})

	log.Fatal(app.Listen(":8080"))
}

func setUpDatabase(host string, port int32, user string, password string, dbname string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&model.Invoice{})

	return db
}
