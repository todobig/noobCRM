package main

import (
	"context"
	"log"
	"os" // Add this line to import the os package

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"1billion/config"
	"1billion/handlers"
	"1billion/repositories"
	"1billion/services"
)

func main() {
	// Load application configuration
	cfg := config.NewConfig()

	// Initialize MongoDB client
	clientOptions := options.Client().ApplyURI(cfg.MongoDBURI)
	mongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(context.Background())

	// Create MongoDB database instance
	mongoDB := mongoClient.Database(cfg.MongoDBName)

	// Initialize repositories
	leadRepository := repositories.NewMongoDBLeadRepository(mongoDB)

	// Initialize services
	leadService := services.NewLeadService(leadRepository)

	// Initialize Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.Default())

	// Initialize API endpoints
	leadHandler := handlers.NewLeadHandler(leadService)
	router.POST("/leads", func(c *gin.Context) { leadHandler.AddLead(c.Writer, c.Request) })
	router.GET("/leads", func(c *gin.Context) { leadHandler.GetLead(c.Writer, c.Request) })
	router.PUT("/leads", func(c *gin.Context) { leadHandler.UpdateLead(c.Writer, c.Request) })
	router.DELETE("/leads", func(c *gin.Context) { leadHandler.DeleteLead(c.Writer, c.Request) })

	// Retrieve certificate and key file paths from environment variables
	certFile := os.Getenv("CERT_FILE_PATH")
	keyFile := os.Getenv("KEY_FILE_PATH")

	// Check if the environment variables are set
	if certFile == "" || keyFile == "" {
		log.Fatal("SSL certificate or key file path environment variables are not set")
	}

	// Run the server with TLS
	port := ":8080"
	log.Printf("Server is running on port %s with SSL", port)
	log.Fatal(router.RunTLS(port, certFile, keyFile))
}
