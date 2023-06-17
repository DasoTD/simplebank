package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	_ "net/http"
	"time"

	"github.com/dasotd/simplebank/gapi"
	"google.golang.org/grpc"

	// "github.com/techschool/simplebank/mail"
	"github.com/dasotd/simplebank/api"
	db "github.com/dasotd/simplebank/db/sqlc"
	"github.com/dasotd/simplebank/pb"
	"github.com/dasotd/simplebank/util"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/reflection"
)





func main(){
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config")
	}
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	  }))

	con, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(con)
	runGrpcServer(config, store)
	// server, err := api.NewServer(config, store)
	// if err != nil {
	// 	log.Fatal("can not start server", err)
	// }
	// err = server.Start(config.ServerAddress)
	// if err != nil {
	// 	log.Fatal("can not start server", err)
	// }
	
}
func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err) //.Err(err).Msg("cannot create server")
	}

	gprcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(gprcLogger)
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create server", err ) //.Err(err).Msg("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String()) //.Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start Grpc server", err) //.Err(err).Msg("cannot start gRPC server")
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)//.Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err) //.Err(err).Msg("cannot start server")
	}
}