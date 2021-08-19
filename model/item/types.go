package item

import "google.golang.org/protobuf/types/known/timestamppb"

type Item struct {
	Id        int64                  `json:"id,omitempty"`
	Name      string                 `json:"name,omitempty"`
	UserId    int64                  `json:"user_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `json:"updated_at,omitempty"`
}
