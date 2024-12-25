package app

import (
	"Hackathon-Management-System/src/graph"
	configuration "Hackathon-Management-System/src/internal/config"
	"Hackathon-Management-System/src/internal/models"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var router = gin.New()

func StartApplication(appConf *configuration.AppConfig) {
	router.Use(gin.Recovery())

	// Initialize database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", appConf.DBConfig.DBUSER, appConf.DBConfig.DBPASSWORD, appConf.DBConfig.DBHOST, appConf.DBConfig.DBPORT, appConf.DBConfig.DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	appConf.DB = db
	err = db.AutoMigrate(models.User{}, models.Hackathon{}, models.Category{})
	if err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}

	// Initialize GraphQL resolver
	resolver := graph.NewResolver(appConf)

	// GraphQL endpoint
	router.POST("/query", func(c *gin.Context) {
		// Extract the Authorization token from the header
		token := c.Request.Header.Get("Authorization")

		// Pass the Authorization token to the context
		ctx := context.WithValue(c.Request.Context(), "Authorization", token)
		ctx = context.WithValue(ctx, "GinContext", c)
		srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
		srv.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	})

	// GraphQL playground endpoint
	router.GET("/", playgroundHandler)

	// REST endpoint
	mapURLs(appConf)

	// Start the Gin server
	log.Printf("Connect to http://localhost:8080/ for GraphQL playground")
	log.Printf("REST server is running on port 8080")

	// Start the Gin server in a separate goroutine
	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf(err.Error())
		}
	}()

	// The main goroutine will terminate after executing all lines of code.
	// To keep the session alive, we use signal handling to listen for interrupt or termination signals.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	// Block until we receive a signal
	<-sig
	fmt.Println("Shutting down server...")

}

// Handler for the GraphQL playground
func playgroundHandler(c *gin.Context) {
	playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
}
