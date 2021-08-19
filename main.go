package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"simple-crud/init"
	"simple-crud/server"
	sea "simple-crud/serviceexampleapi"
)

func main() {

	s := grpc.NewServer()

	cfg, err := initialization.InitCfg("config/config.yml")
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	log.Println("config initiated")

	lis, err := net.Listen("tcp", ":" + cfg.Http.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("listening to port:", cfg.Http.Port)

	db, err:= initialization.InitDB(cfg.CfgDB)
	defer db.Close()

	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	log.Println("db connention initiated")

	sea.RegisterServiceExampleServiceServer(s, &server.Obj{
		CFG:cfg,
		DB:db,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("server started")
}
