package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/x893675/graphql-sample/app/domain/repository/answer"
	"github.com/x893675/graphql-sample/app/domain/repository/question"
	"github.com/x893675/graphql-sample/app/domain/repository/question_option"
	"github.com/x893675/graphql-sample/app/generated"
	"github.com/x893675/graphql-sample/app/infrastructure/db"
	"github.com/x893675/graphql-sample/app/infrastructure/persistence"
	"github.com/x893675/graphql-sample/app/interfaces"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	var (
		defaultPort      = "8080"
		databaseUser     = os.Getenv("DATABASE_USER")
		databaseName     = os.Getenv("DATABASE_NAME")
		databaseHost     = os.Getenv("DATABASE_HOST")
		databasePort     = os.Getenv("DATABASE_PORT")
		databasePassword = os.Getenv("DATABASE_PASSWORD")
		debug            = os.Getenv("DEBUG")
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, databasePort, databaseUser, databaseName, databasePassword)
	conn := db.OpenDB(dbConn)
	if err := db.AutoMigrate(conn); err != nil {
		panic(err)
	}

	var ansService answer.AnsService
	var questionService question.QuesService
	var questionOptService question_option.OptService

	ansService = persistence.NewAnswer(conn)
	questionService = persistence.NewQuestion(conn)
	questionOptService = persistence.NewQuestionOption(conn)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		AnsService:            ansService,
		QuestionService:       questionService,
		QuestionOptionService: questionOptService,
	}}))

	if strings.ToLower(debug) == "true" {
		http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	}
	http.Handle("/graphql", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
