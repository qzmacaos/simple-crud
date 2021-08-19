package tests

import (
	"context"
	"google.golang.org/grpc"
	"log"
	sea "simple-crud/serviceexampleapi"
	"testing"
	"time"
)

const (
	address     = "localhost:50000"
)
//
//func TestCreate(t *testing.T) {
//	// Set up a connection to the server.
//	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
//	if err != nil {
//		t.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	c := sea.NewServiceExampleServiceClient(conn)
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//	user, err := c.CreateUser(ctx, &sea.CreateUserRequest{
//		Name: "Vasya",
//		Age: 20,
//		UserType: sea.UserType(0),
//
//	})
//	if err != nil {
//		t.Error("CreateUser error ", err)
//	}
//
//	if user.Id =="" {
//		t.Error("User wasnt created", err)
//	}
//}
//
//func TestUpdate(t *testing.T) {
//	// Set up a connection to the server.
//	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
//	if err != nil {
//		t.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	c := sea.NewServiceExampleServiceClient(conn)
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//	_, err = c.UpdateUser(ctx, &sea.UpdateUserRequest{
//		Id:"2",
//		Name: "Petya",
//		Age: 30,
//		UserType: sea.UserType(1),
//
//	})
//
//	if err != nil {
//		t.Error("UpdateUser error ", err)
//	}
//}
//
//func TestDelete(t *testing.T) {
//	// Set up a connection to the server.
//	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
//	if err != nil {
//		t.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	c := sea.NewServiceExampleServiceClient(conn)
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//	_, err = c.DeleteUser(ctx, &sea.DeleteUserRequest{
//		Id:"1",
//
//	})
//
//	if err != nil {
//		t.Error("DeleteUser error ", err)
//	}
//}
//
//func TestItemCreate(t *testing.T) {
//	// Set up a connection to the server.
//	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
//	if err != nil {
//		t.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	c := sea.NewServiceExampleServiceClient(conn)
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//	_, err = c.CreateItem(ctx, &sea.CreateItemRequest{
//		Name: "test item",
//		UserId: "2",
//	})
//
//	if err != nil {
//		t.Error("CreateItem error ", err)
//	}
//}
//
//
//func TestItemUpdate(t *testing.T) {
//	// Set up a connection to the server.
//	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
//	if err != nil {
//		t.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	c := sea.NewServiceExampleServiceClient(conn)
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//	_, err = c.UpdateItem(ctx, &sea.UpdateItemRequest{
//		Id:"1",
//		Name:"Another test item",
//	})
//
//	if err != nil {
//		t.Error("UpdateItem error ", err)
//	}
//}
//
//func TestGet(t *testing.T) {
//	// Set up a connection to the server.
//	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
//	if err != nil {
//		t.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	c := sea.NewServiceExampleServiceClient(conn)
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//	user, err := c.GetUser(ctx, &sea.GetUserRequest{
//		Id:"2",
//	})
//
//	log.Printf("%#v", user)
//
//	if err != nil {
//		t.Error("GetUser error ", err)
//	}
//}

func TestList(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := sea.NewServiceExampleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.ListUser(ctx, &sea.ListUserRequest{
		PageFilter:&sea.PageFilter{
			Limit: 1,
		},
	})

	if err != nil {
		t.Error("GetUser error ", err)
	}

	users := resp.GetUsers()

	if len(users) >1{
		t.Error("Too much")
	}

	if len(users) <1{
		t.Error("Not enough")
	}

	log.Printf("%#v", *users[0])
}