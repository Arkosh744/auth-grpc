package user

import (
	"strings"

	"github.com/Arkosh744/auth-grpc/internal/model"
	desc "github.com/Arkosh744/auth-grpc/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUser(user *desc.CreateRequest) *model.User {
	return &model.User{
		Username: user.GetUser().GetUsername(),
		Email:    strings.ToLower(strings.TrimSpace(user.GetUser().GetEmail())),
		Password: user.GetUser().GetUsername(),
		Role:     model.StringToRole(user.GetUser().GetRole()),
	}
}

func ToUpdateUser(req *desc.UpdateRequest) *model.UpdateUser {
	user := &model.UpdateUser{}

	if req.GetNewRole() != nil {
		user.Role.String = model.StringToRole(req.GetNewRole().GetValue()).String()
		user.Role.Valid = true
	}

	if req.GetNewUsername() != nil {
		user.Username.String = req.GetNewUsername().GetValue()
		user.Username.Valid = true
	}

	if req.GetNewEmail() != nil {
		user.Email.String = req.GetNewEmail().GetValue()
		user.Email.Valid = true
	}

	if req.GetNewPassword() != nil {
		user.Password.String = req.GetNewPassword().GetValue()
		user.Password.Valid = true
	}

	return user
}

func ToGetResponse(user *model.User) *desc.GetResponse {
	return &desc.GetResponse{
		User: &desc.UserInfo{
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role.String()},
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
