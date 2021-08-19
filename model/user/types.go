package user

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"simple-crud/model/item"
	sea "simple-crud/serviceexampleapi"
)

type User struct {
	Id        int64                  `json:"id,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Age       int32                  `json:"age,omitempty"`
	UserType  sea.UserType           `json:"user_type,omitempty"`
	Items     []*item.Item           `json:"items,omitempty"`
	CreatedAt *timestamppb.Timestamp `json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `json:"updated_at,omitempty"`
}
