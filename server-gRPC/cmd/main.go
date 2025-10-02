package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ChanatpakornS/inventory-demo/gRPC/internal/handlers"
	model "github.com/ChanatpakornS/inventory-demo/gRPC/internal/models"

	invoicegRPC "github.com/ChanatpakornS/inventory-demo/gRPC/grpc-proto/invoice"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "idb"
	password = "idb"
	dbname   = "idb"
)

func main() {
	db := setUpDatabase(host, port, user, password, dbname)

	invoiceHandler := handlers.NewInvoiceHandler(db)

	grpcServer := grpc.NewServer()
	invoicegRPC.RegisterInvoiceServiceServer(grpcServer, invoiceHandler)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
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
