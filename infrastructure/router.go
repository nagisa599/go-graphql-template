package infrastructure

import (
	"context"
	"errors"
	"log" // Standard log package
	"log/slog"
	"net/http"
	"os"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nagisa599/go-graphql-template/graphql"
	"github.com/nagisa599/go-graphql-template/graphql/resolvers"
	"github.com/nagisa599/go-graphql-template/internal/domain/repository"
	"github.com/nagisa599/go-graphql-template/internal/handler"
	"github.com/nagisa599/go-graphql-template/internal/usecase"
	"github.com/nagisa599/go-graphql-template/middleware"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

func Router() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set.")
	}

	databaseHandler := NewDatabaseHandler()
    

	userHandler := handler.NewUserHandler(usecase.NewUserUsecase(
		repository.NewUserRepository(databaseHandler),
	))
	srv := graphql_handler.NewDefaultServer(
		graphql.NewExecutableSchema(
			graphql.Config{
				Resolvers: resolvers.NewResolver(*userHandler),
			},
		),
	)
	srv.SetErrorPresenter(handleError)
  // TODO: JWTからuser_idを取得するmiddlewareを実行
	srv.AroundOperations(middleware.JwtAuthenticateMiddleware)
	// ログを出力するmiddlewareを実行
	srv.AroundOperations(middleware.LoggerMiddleware)

	// CORS settings
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // TODO: Change this for production
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
	})
	handler := c.Handler(srv)

	if os.Getenv("ENV") == "development" {
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}
	http.Handle("/query", handler)

	defer databaseHandler.Close()

	// Add log statement before starting the server
	log.Printf("Server starting on port %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func handleError(ctx context.Context, err error) *gqlerror.Error {
	slog.ErrorContext(ctx, err.Error())
	
	if os.Getenv("ENV") == "development" {
		// Return detailed error in development environment
		return gqlerror.Errorf(err.Error())
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return gqlerror.Errorf("record not found")
	} else {
		// TODO: Send error to Sentry
		return gqlerror.Errorf("internal server error")
	}
}
