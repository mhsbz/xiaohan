package services

// 本文档为指令系统
import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/mhsbz/xiaohan/api/gen/xiaohan/server/operations"
	"github.com/mhsbz/xiaohan/internal/repository"
	"github.com/mhsbz/xiaohan/internal/schemas"
	"math/rand"
	"strings"
	"time"
)

type Service struct {
	IUser     IUserLogic
	dataStore *repository.MongoClient
	a         string
	b         string
	c         string
}

func NewService() *Service {
	rand.New(rand.NewSource(time.Now().UnixNano()))
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
	//case "踏入仙途":
	//	fmt.Println(params.MemberID)
	//	responseStr = fmt.Sprintf("阁下是踏入仙途的第1位道友，道号：")
	//	user, err := s.CreateOrGetUser(params.MemberID)
	//	if err != nil {
	//		return operations.NewActionInternalServerError().WithPayload("Internal Server Error")
	//	}
	//	responseStr = fmt.Sprintf("阁下是踏入仙途的第%d位道友，道号：%s", user.Rank, user.Nickname)
	case "":
		responseStr = "当你召唤我的时候，你的路就只有一条，加入赛博修仙界,输入“加入异世界修仙”进入游戏"
	case "加入异世界修仙":
		responseStr = "那是一个风雨交加的晚上，你在公司加完班迎着雨滴走在马路上，一辆刹车打滑的大卡车创向了你，一串串回马灯般的画面闪入你的脑海中，但是你并没有死，你降临到了一个名为AU的大陆，这片大陆本身是一个异世界大陆，几百年前，世界上最顶尖的巫术师聚集在了一起，不知是何原因，创造了一个超大的魔法阵，将另一个位面的大陆拖拽合并了进来，然而这片大陆上生存的，居然是修仙界的人类......\n请选择你的阵营：\n1. 加入修仙界\n2. 加入异世界"
	case "加入修仙界":
		s.a = "修仙界" // 用户选择了加入修仙界
		responseStr = "请选择术修方向或者剑修方向"
	case "术修":
		s.b = "术修" // 用户在修仙界中选择了术修方向
		responseStr = "您已选择术修方向，输入生成角色获取您的初始信息"
	case "剑修":
		s.b = "剑修" // 用户在修仙界中选择了剑修方向
		responseStr = "您已选择剑修方向，输入生成角色获取您的初始信息"
	case "加入异世界":
		s.a = "异世界" // 用户选择了加入异世界
		responseStr = "请选择魔法路线和剑士路线"
	case "魔法":
		s.c = "魔法师" // 用户在异世界中选择了魔法路线
		responseStr = "您已选择魔法路线，输入生成角色获取您的初始信息"
	case "剑士":
		s.c = "剑士" // 用户在异世界中选择了剑士路线
		responseStr = "您已选择剑士路线，输入生成角色获取您的初始信息"
	case "生成角色":
		// 使用schemas.User来创建新角色
		newUser := schemas.NewUser("username")
		// 根据之前的选择生成角色信息
		initialMeridian := generateMeridian()
		responseStr = fmt.Sprintf("创建角色成功，您是第%d位进入AU界的玩家，您的角色名称为%s，是%s%s，诞生于公元%s年，你从母亲怀中降生之日，AU大陆的光芒赐福于您，获得了初始命脉：%s",
			newUser.Uid, newUser.Nickname, s.a, combineBCValue(s.b, s.c), time.Now().Format("2006"), initialMeridian)
	case "领取内测专属奖励":
		responseStr = "恭喜道友获得由453411753内测群发出的内测专属奖励，内测专属称号：AU仙人，持有效果：幸运值+1"
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

	}

	return operations.NewActionOK().WithPayload(responseStr)
}

// 生成角色时的假设的辅助函数来组合b和c的值，根据实际需要实现
func combineBCValue(b string, c string) string {
	return fmt.Sprintf("%s的%s", b, c)
}

// 命脉及其概率定义
var meridians = []struct {
	Name        string
	Probability float64
}{
	{"轮回觉醒者", 0.005},
	{"怪盗基德", 0.01},
	// 添加剩余八个命脉，这里假设它们的概率相等
	{"命脉3", 0.0985 / 8},
	{"命脉4", 0.0985 / 8},
	{"命脉5", 0.0985 / 8},
	{"命脉6", 0.0985 / 8},
	{"命脉7", 0.0985 / 8},
	{"命脉8", 0.0985 / 8},
	{"命脉9", 0.0985 / 8},
	{"命脉10", 0.0985 / 8},
}

// 命脉生成随机数规则
func generateMeridian() string {
	// 计算所有命脉的总概率
	totalProbability := 0.0
	for _, v := range meridians {
		totalProbability += v.Probability
	}

	// 生成一个介于0到总概率之间的随机数
	randomNum := rand.Float64() * totalProbability

	// 遍历命脉列表，累加概率直到找到对应的命脉
	accumulatedProb := 0.0
	for _, v := range meridians {
		accumulatedProb += v.Probability
		if randomNum <= accumulatedProb {
			return v.Name
		}
	}

	// 理论上不会走到这里，但作为一个安全措施，如果没有匹配到任何命脉，则返回一个默认值
	return "未知命脉"
}
