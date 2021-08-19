package tests

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"

	sea "simple-crud/serviceexampleapi"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func TestName(t *testing.T) {



}



func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := sea.NewServiceExampleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user, err := c.CreateUser(ctx, &sea.CreateUserRequest{
		Name: "Vasya",
		Age: 20,
		UserType: sea.UserType(0),

	})
	if err != nil {
		log.Fatalf("CreateUser error ", err)
	}

	if user.Id =="" {
		log.Fatalf("User wasnt created", err)
	}
}
