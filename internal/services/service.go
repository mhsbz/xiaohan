package services

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/mhsbz/xiaohan/api/gen/xiaohan/server/operations"
	"github.com/mhsbz/xiaohan/internal/repository"
	"strings"
)

// 本文档为指令系统
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
	//所有指令信息
	/*case "重入仙途":
	user, err := s.CreateOrGetUser(params.MemberID)
	if err != nil {
		return operations.NewActionInternalServerError().WithPayload("Internal Server Error")
	}


	user.MM = "GenerateRandommingmai()"
	responseStr = "nide mingmai =" + user.MM*/
	case "踏入仙途":
		fmt.Println(params.MemberID)
		responseStr = fmt.Sprintf("阁下是踏入仙途的第1位道友，道号：")
		user, err := s.CreateOrGetUser(params.MemberID)
		if err != nil {
			return operations.NewActionInternalServerError().WithPayload("Internal Server Error")
		}
		responseStr = fmt.Sprintf("阁下是踏入仙途的第%d位道友，道号：%s", user.Rank, user.Nickname)
	case "":
		responseStr = "当你召唤我的时候，你的路就只有一条，加入赛博修仙界,输入“踏入仙途”加入au修仙1.0"
	case "修炼":
		responseStr = "道友当前选择的地点为x，将在此地进行云游修炼，预计时长：x分钟"
		//return operations.NewActionOK().WithPayload(responseStr)
		//user, err := s.CreateOrGetUser(params.MemberID)
		//if err != nil {
		//	return operations.NewActionInternalServerError().WithPayload("Internal Server Error")
		//}
		//if user.XStatus {
		//	responseStr = "niyijingzaixiulian"
		//}
		//responseStr = "已经开始修炼"
	case "修仙帮助":
		responseStr = "目前为测试版暂时只有文字菜单，可输入踏入仙途进行创建初始角色"
	case "修仙信息":
		responseStr = "idx，名称x，你好，您的当前境界为：x，当前修为：x"
	case "领取内测专属奖励":
		responseStr = "恭喜道友获得由453411753内测群发出的内测专属奖励，内测专属称号：AU仙人，持有效果：幸运值+1"
	case "1":
		responseStr = "1你吗"

	}

	return operations.NewActionOK().WithPayload(responseStr)
}
