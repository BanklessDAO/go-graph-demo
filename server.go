package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"go-graph-demo/app/generated"
	"go-graph-demo/app/infrastructure/database"
	"go-graph-demo/app/infrastructure/persistence"
	"go-graph-demo/app/interfaces"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func init() {
	// loads values from .env into the system
	err := godotenv.Load()
	if err != nil {
		log.Warnln("Skipped loading .env file because it was not found...")
	}
	// Setup logrus formatter
	log.SetFormatter(
		&log.JSONFormatter{
			PrettyPrint: false,
		},
	)

	log.SetReportCaller(false)

	// Setup log level
	switch os.Getenv("LOG_LEVEL") {
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
	default:
		log.Warnln("No valid log level was specified, defaulting to info")
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Instantiate services
	db, err := database.OpenDB()
	if err != nil {
		log.Panicf("caught error initializing repositories, err %+v", err)
	}
	defer database.Close(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		Repositories: persistence.Repositories{
			Bounties: persistence.NewBountyRepository(db),
			Users:    persistence.NewUserRepository(db),
		},
	}}))

	// Set the catch error function
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		return err
	})

	// Set the recovery function
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		log.Errorf("caught error, %+v", err)
		return fmt.Errorf("%+v", err)
	})

	http.Handle("/", playground.Handler("BanklessDAO GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
