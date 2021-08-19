Нужно реализовать CRUD grpc server, proto file ниже

Не использовать ORM (like gorm, etc)

Для миграций можно использовать:
https://github.com/rubenv/sql-migrate

Рекомендации по структуре проекта можно посмотреть тут:
https://github.com/golang-standards/project-layout

Database: postgres

```bigquery
syntax = "proto3";

package example.service;

option go_package = "serviceexampleapi";

import "google/protobuf/timestamp.proto";
import "google/api/annotations.proto";

service ServiceExampleService {
    rpc CreateUser(CreateUserRequest) returns (User) {
        option (google.api.http) = {
            post: "/service-example/v1/user"
            body: "*"
        };
}

    rpc UpdateUser(UpdateUserRequest) returns (User) {
        option (google.api.http) = {
            post: "/service-example/v1/user"
            body: "*"
        };
}

    rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse) {
        option (google.api.http) = {
delete : "/service-example/v1/user/{id}"
    };
}

    rpc ListUser(ListUserRequest) returns (ListUserResponse) {
        option (google.api.http) = {
            get: "/service-example/v1/user"
        };
}

    rpc GetUser(GetUserRequest) returns (User) {
        option (google.api.http) = {
            get: "/service-example/v1/user/{id}"
        };
}

    rpc CreateItem(CreateItemRequest) returns (Item) {
        option (google.api.http) = {
            post: "/service-example/v1/item"
            body: "*"
        };
}

    rpc UpdateItem(UpdateItemRequest) returns (Item) {
        option (google.api.http) = {
            put: "/service-example/v1/item/{id}"
            body: "*"
        };
}
}

// rpc
message CreateUserRequest {
    string name = 1;
int32 age = 2;
UserType user_type = 3;
repeated  CreateItemRequest items = 4;
optional int32 test = 5;
}

message CreateItemRequest {
    string name = 1;
string user_id = 2;
}

message UpdateUserRequest {
    string id = 1;
string name = 2;
int32 age = 3;
UserType user_type = 4;
repeated UpdateItemRequest items = 5;
}

message UpdateItemRequest {
    string id = 1;
string name = 2;
}

message DeleteUserRequest {
    string id = 1;
}

message DeleteUserResponse {}

message ListUserRequest {
    PageFilter page_filter = 1;
}

message ListUserResponse {
    repeated User users = 1;
}

message GetUserRequest {
    string id = 1;
}

// meta
message PageFilter {
    uint32 limit = 1;
uint32 page = 2;
}

enum UserType {
    INVALID_USER_TYPE = 0;
EMPLOYEE_USER_TYPE = 1;
CUSTOMER_USER_TYPE = 2;
}

message User {
    string id = 1;
string name = 2;
int32 age = 3;
UserType user_type = 4;
repeated Item items = 5;
google.protobuf.Timestamp created_at = 6;
google.protobuf.Timestamp updated_at = 7;
}

message Item {
    string id = 1;
string name = 2;
string user_id = 3;
google.protobuf.Timestamp created_at = 5;
google.protobuf.Timestamp updated_at = 6;
}

```