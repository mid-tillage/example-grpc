package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	catalog "github.com/sys-internals/example-grpc/proto/catalog"
	"google.golang.org/grpc"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "1234"
	dbName     = "postgres"
)

// YourCatalogServiceServer implements the CatalogServiceServer interface
type YourCatalogServiceServer struct {
	db *sql.DB
}

// NewCatalogServiceServer creates a new instance of YourCatalogServiceServer
func NewCatalogServiceServer() (*YourCatalogServiceServer, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	return &YourCatalogServiceServer{db: db}, nil
}

func (s *YourCatalogServiceServer) CreateCatalog(ctx context.Context, req *catalog.CreateCatalogRequest) (*catalog.Catalog, error) {
	// createSql := `
	// create table if not exists catalog(
	// 	id SERIAL PRIMARY KEY,
	// 	name varchar2
	// );
	// `
	// _, err := s.db.BeginTx(context.Background(), createSql)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Table creation failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// Assuming "catalogs" table has "id" and "name" columns
	const createCatalogQuery = "INSERT INTO catalogs (name) VALUES ($1) RETURNING id"

	var catalogID int32
	err := s.db.QueryRowContext(ctx, createCatalogQuery, req.Name).Scan(&catalogID)
	if err != nil {
		log.Printf("Failed to create catalog: %v", err)
		return nil, err
	}

	// Return the created catalog
	return &catalog.Catalog{
		Id:   catalogID,
		Name: req.Name,
	}, nil
}

func (s *YourCatalogServiceServer) GetCatalog(ctx context.Context, req *catalog.GetCatalogRequest) (*catalog.GetCatalogResponse, error) {
	// Assuming "catalogs" table has "id" and "name" columns
	const getCatalogQuery = "SELECT id, name FROM catalogs"

	rows, err := s.db.QueryContext(ctx, getCatalogQuery)
	if err != nil {
		log.Printf("Failed to fetch catalogs: %v", err)
		return nil, err
	}
	defer rows.Close()

	var catalogs []*catalog.Catalog
	for rows.Next() {
		var catalog catalog.Catalog
		err := rows.Scan(&catalog.Id, &catalog.Name)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			return nil, err
		}
		catalogs = append(catalogs, &catalog)
	}

	// Return the list of catalogs
	return &catalog.GetCatalogResponse{
		Catalog: catalogs,
	}, nil
}

// Initialize gRPC server
func main() {
	server, err := NewCatalogServiceServer()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
		return
	}

	grpcServer := grpc.NewServer()
	catalog.RegisterCatalogServiceServer(grpcServer, server)

	log.Println("gRPC server is running...")
	if err := grpcServer.Serve(server); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
