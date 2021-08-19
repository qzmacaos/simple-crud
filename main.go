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

	cfg, err := initialization.InitCfg()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	lis, err := net.Listen("tcp", ":" + cfg.Http.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err:= initialization.InitDB(cfg.CfgDB)
	defer db.Close()

	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	sea.RegisterServiceExampleServiceServer(s, &server.Obj{
		CFG:cfg,
		DB:db,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
