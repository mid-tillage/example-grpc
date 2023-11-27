package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/sys-internals/example-grpc/catalog"
)

const (
	serverAddress = "localhost:50051"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewCatalogServiceClient(conn)

	// Call CreateCatalog
	createCatalogResponse, err := client.CreateCatalog(context.Background(), &pb.CreateCatalogRequest{Name: "ExampleCatalog"})
	if err != nil {
		log.Fatalf("CreateCatalog failed: %v", err)
	}
	fmt.Printf("Catalog created with ID: %d\n", createCatalogResponse.Id)

	// Call GetCatalog
	getCatalogResponse, err := client.GetCatalog(context.Background(), &pb.GetCatalogRequest{})
	if err != nil {
		log.Fatalf("GetCatalog failed: %v", err)
	}

	// Print the retrieved catalogs
	fmt.Println("Catalogs:")
	for _, catalog := range getCatalogResponse.Catalog {
		fmt.Printf("ID: %d, Name: %s\n", catalog.Id, catalog.Name)
	}
}
