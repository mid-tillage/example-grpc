package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v4"
	pb "github.com/sys-internals/example-grpc/proto/productonsale"
	"google.golang.org/grpc"
)

const (
	port = ":3001"
)

func productOnSaleServer() *ProductOnSaleServer {
	return &ProductOnSaleServer{}
}

type ProductOnSaleServer struct {
	conn                *pgx.Conn
	first_user_creation bool
	pb.UnimplementedProductOnSaleServer
}

func (server *ProductOnSaleServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductOnSaleServer(s, server)
	log.Printf("server listening at %v", lis.Addr())

	return s.Serve(lis)
}

func (server *ProductOnSaleServer) CreateProductOnSale(ctx context.Context, in *pb.ProductOnSale) (*pb.User, error) {
	createSql := `
	create table if not exists catalog(
		idProductCatalog SERIAL PRIMARY KEY,
		name varchar2
	);
	create table if not exists Enterprise(
		idEnterprise SERIAL PRIMARY KEY,
		name varchar2
	);
	create table if not exists Product(
		idProduct SERIAL PRIMARY KEY,
		name varchar2,
		description varchar2,
		idEnterprise int8
	)
	create table if not exists productOnSale(
		idProductOnSale SERIAL PRIMARY KEY,
		title varchar2,
		idProduct int8,
		price int8,
		saleStartDatetime string,
		saleEndDatetime string,
		idCatalog int8
	);
	`

	_, err := server.conn.Exec(context.Background(), createSql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Table creation failed: %v\n", err)
		os.Exit(1)
	}

	server.first_user_creation = false

	log.Printf("Received: %v", in.GetName())

	createdProductOnSale := &pb.ProductOnSale{
		IdProductOnSale:   in.GetIdProductOnSale(),
		Title:             in.GetTitle(),
		Product:           in.GetProduct(),
		Price:             in.GetPrice(),
		SaleStartDatetime: in.GetSaleStartDatetime(),
		SaleEndDatetime:   in.GetSaleEndDatetime(),
		Catalog:           in.GetCatalog(),
	}

	_, err = tx.Exec(context.Background(), "insert into catalog(name) values ($1)",
		createdProductOnSale.Catalog.Name)
	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}

	_, err = tx.Exec(context.Background(), "insert into enterprise(name) values ($1)",
		createdProductOnSale.Product.Enterprise.Name)
	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}

	_, err = tx.Exec(context.Background(), "insert into Product(name, description, enterprise) values ($1,$2)",
		createdProductOnSale.Name, createdProductOnSale.Age)
	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}

	_, err = tx.Exec(context.Background(), "insert into users(name, age) values ($1,$2)",
		createdProductOnSale.Name, createdProductOnSale.Age)
	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}

	tx.Commit(context.Background())
	return createdProductOnSale, nil
}
