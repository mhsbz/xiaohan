package services

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/mhsbz/xiaohan/api/gen/xiaohan/server/operations"
	"github.com/mhsbz/xiaohan/internal/repository"
	"strings"
)

type Service struct {
	IUser     IUserLogic
	dataStore *repository.MongoClient
}

func NewService() *Service {
	return &Service{
		dataStore: repository.NewMongoClient(),
	}
}

func (s *Service) Action(params operations.ActionParams) middleware.Responder {
	var responseStr string
	action := strings.TrimSpace(params.Action)

	switch action {
	case "踏入仙途":
		fmt.Println(params.MemberID)
		user, err := s.CreateOrGetUser(params.MemberID)
		if err != nil {
			return operations.NewActionInternalServerError().WithPayload("Internal Server Error")
		}
		responseStr = fmt.Sprintf("阁下是踏入仙途的第%d位道友，道号：%s", user.Rank, user.Nickname)
	}

	return operations.NewActionOK().WithPayload(responseStr)
}
