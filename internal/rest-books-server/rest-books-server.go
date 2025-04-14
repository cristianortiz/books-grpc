package restbooksserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cristianortiz/books-grpc/internal/pkg/configs"
	"github.com/cristianortiz/books-grpc/internal/pkg/db"
	"github.com/cristianortiz/books-grpc/internal/pkg/db/migrations"
	"github.com/cristianortiz/books-grpc/internal/pkg/repository"
	"github.com/cristianortiz/books-grpc/internal/pkg/service"
	"gorm.io/gorm"
)

type App struct {
	dbConn *gorm.DB
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	appConfig, err := configs.ProvideAppConfig()
	if err != nil {
		log.Fatal(err)
	}
	dbConn, err := db.ProvideDBConn(&appConfig.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	a.dbConn = dbConn
	migrator, err := migrations.ProvideMigrator(appConfig.DBConfig, dbConn)
	if err != nil {
		log.Fatal(err)
	}
	migrator.RunMigrations()

	booksRepo := repository.NewBookRepository(dbConn)
	booksSrv := service.NewBooksService(booksRepo)
	r := ProvideRouter(booksSrv)
	srv := http.Server{
		Addr:         fmt.Sprintf("%s:%d", appConfig.ServerConfig.Host, appConfig.DBConfig.Port),
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Starting Server")
	log.Fatal((srv.ListenAndServe()))
}

func (a *App) Shutdown() {
	dbIntance, _ := a.dbConn.DB()
	_ = dbIntance.Close()
}
