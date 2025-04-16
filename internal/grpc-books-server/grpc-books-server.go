package grpcbooksserver

import (
	"fmt"
	"log"
	"net"

	"github.com/cristianortiz/books-grpc/internal/pkg/configs"
	"github.com/cristianortiz/books-grpc/internal/pkg/db"
	"github.com/cristianortiz/books-grpc/internal/pkg/db/migrations"
	"github.com/cristianortiz/books-grpc/internal/pkg/proto"
	"github.com/cristianortiz/books-grpc/internal/pkg/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

// used to encapsulate the grpc server functionality
type App struct {
	//Automatically generated interface by the protoc compiler, used to ensure type safety and avoid compilation errors in the application
	proto.UnimplementedBookServiceServer
	dbConn   *gorm.DB
	bookRepo *repository.BookRepository
}

// creates and returns a new instance of App struct
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
	a.bookRepo = repository.NewBookRepository(a.dbConn)
	migrator, err := migrations.ProvideMigrator(appConfig.DBConfig, dbConn)
	if err != nil {
		log.Fatal(err)
	}
	migrator.RunMigrations()
	servAddr := fmt.Sprintf("0.0.0.0:%d", appConfig.ServerConfig.Port)
	fmt.Println("starting books gRPC server at", servAddr)

	lis, err := net.Listen("tcp", servAddr)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	opts := []grpc.ServerOption{}
	//creates new grpc server
	s := grpc.NewServer(opts...)
	//register the gRPC server implementation 's' wich is the App struct instance
	// and has the book server interface implemented and bound to it with the server using proto.RegisterBookServiceServer
	proto.RegisterBookServiceServer(s, a)
	//register gRPC server reflection, to provide info about grpc services publicily on the server
	//and assists clients at runtime to construct rpc requests and responses without precompiled service info
	//or for debuggin and testing
	reflection.Register(s)
	//start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// gracefully shut down the grpc server, it closes the db conn
func (a *App) Shutdown() {
	dbInstance, _ := a.dbConn.DB()
	_ = dbInstance.Close()
}
