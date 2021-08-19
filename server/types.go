package server

import (
	"database/sql"
	"simple-crud/config"
	sea "simple-crud/serviceexampleapi"
)

const pageSize = 10

type Obj struct {
	sea.UnimplementedServiceExampleServiceServer
	CFG config.Config
	DB *sql.DB
}