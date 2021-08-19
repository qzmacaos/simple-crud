package server

import (
	"database/sql"
	"simple-crud/config"
	sea "simple-crud/serviceexampleapi"
)

type Obj struct {
	sea.UnimplementedServiceExampleServiceServer
	CFG config.Config
	DB *sql.DB
}