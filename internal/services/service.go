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
	case "我要修仙":
		responseStr = "当你召唤我的时候，你的路就只有一条，加入赛博修仙界,输入“踏入仙途”加入au修仙1.0"

	case "重入仙途":
		user, err := s.CreateOrGetUser(params.MemberID)
		if err != nil {
			return operations.NewActionInternalServerError().WithPayload("Internal Server Error")
		}

		user.MM = "GenerateRandommingmai()"
		responseStr = "nide mingmai =" + user.MM

	case "修炼":
		responseStr = "已经开始修炼"

		//return operations.NewActionOK().WithPayload(responseStr)
		//user, err := s.CreateOrGetUser(params.MemberID)
		//if err != nil {
		//	return operations.NewActionInternalServerError().WithPayload("Internal Server Error")
		//}
		//if user.XStatus {
		//	responseStr = "niyijingzaixiulian"
		//}
		//responseStr = "已经开始修炼"

	case "踏入仙途":
		fmt.Println(params.MemberID)
		responseStr = fmt.Sprintf("阁下是踏入仙途的第1位道友，道号：")
		user, err := s.CreateOrGetUser(params.MemberID)
		if err != nil {
			return operations.NewActionInternalServerError().WithPayload("Internal Server Error")
		}
		responseStr = fmt.Sprintf("阁下是踏入仙途的第%d位道友，道号：%s", user.Rank, user.Nickname)
	}

	return operations.NewActionOK().WithPayload(responseStr)
}
