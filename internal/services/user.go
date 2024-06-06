package services

import "github.com/mhsbz/xiaohan/internal/schemas"

type IUserLogic interface {
	NewUser(memberID string) *schemas.User
}

func (s *Service) NewUser(memberID string) *schemas.User {
	return schemas.NewUser(memberID)
}
