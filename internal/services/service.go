package services

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/mhsbz/xiaohan/api/gen/xiaohan/server/operations"
)

type Service struct {
	IUser IUserLogic
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Action(params operations.ActionParams) middleware.Responder {
	var responseStr string

	switch params.Action {
	case "踏入仙途":
		user := s.NewUser(params.MemberID)
		responseStr = fmt.Sprintf("阁下是踏入仙途的第%d位道友，道号：%s", user.Rank, user.Nickname)
	}

	return operations.NewActionOK().WithPayload(&operations.ActionOKBody{Result: responseStr})
}
