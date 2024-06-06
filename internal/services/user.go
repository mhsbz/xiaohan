package services

import "github.com/mhsbz/xiaohan/internal/schemas"

type IUserLogic interface {
	CreateOrGetUser(*schemas.User, error)
}

var rank = 0

func (s *Service) CreateOrGetUser(memberID string) (*schemas.User, error) {
	user, err := s.dataStore.GetUser(memberID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		user = schemas.NewUser(memberID)
		user.Rank = rank
		if err1 := s.dataStore.CreateUser(user); err1 != nil {
			return nil, err1
		}
		rank++
	}
	return user, nil
}
