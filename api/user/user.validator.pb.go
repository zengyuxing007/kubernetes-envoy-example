// Code generated by protoc-gen-gogo.
// source: user/user.proto
// DO NOT EDIT!

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user/user.proto

It has these top-level messages:
	User
	CreateUserRequest
	GetUserRequest
	ListUsersRequest
	ListUsersResponse
	DeleteUserRequest
*/
package user

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *User) Validate() error {
	return nil
}
func (this *CreateUserRequest) Validate() error {
	return nil
}
func (this *GetUserRequest) Validate() error {
	return nil
}
func (this *ListUsersRequest) Validate() error {
	return nil
}
func (this *ListUsersResponse) Validate() error {
	for _, item := range this.Users {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
			}
		}
	}
	return nil
}
func (this *DeleteUserRequest) Validate() error {
	return nil
}