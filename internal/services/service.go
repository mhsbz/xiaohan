package services

// 本文档为指令系统
import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/mhsbz/xiaohan/api/gen/xiaohan/server/operations"
	"github.com/mhsbz/xiaohan/internal/schemas"
	"github.com/mhsbz/xiaohan/pkg/utils"
	"strings"
	"time"
)

type Service struct {
	IUser        *UserService
	ITraining    *TrainingService
	IMenu        *MenuService
	IDungeon     *DungeonService
	IFight       *FightService
	IDescription *DescriptionService
}

func NewService() *Service {
	return &Service{
		IUser:        &UserService{},
		ITraining:    &TrainingService{},
		IMenu:        &MenuService{},
		IDungeon:     &DungeonService{},
		IFight:       &FightService{},
		IDescription: &DescriptionService{},
	}
}

func (s *Service) Action(params operations.ActionParams) middleware.Responder {
	var responseStr string
	action := strings.TrimSpace(params.Action)

	switch {
	case action == "菜单":
		responseStr = s.IMenu.MenuList()
	case action == "":
		responseStr = "当你召唤我的时候，你的路就只有一条，加入赛博修仙界,输入“加入异世界修仙”进入游戏"
	case action == "加入异世界修仙":
		responseStr = "那是一个风雨交加的晚上，你在公司加完班迎着雨滴走在马路上，一辆刹车打滑的大卡车创向了你，一串串回马灯般的画面闪入你的脑海中，但是你并没有死，你降临到了一个名为AU的大陆，这片大陆本身是一个异世界大陆，几百年前，世界上最顶尖的巫术师聚集在了一起，不知是何原因，创造了一个超大的魔法阵，将另一个位面的大陆拖拽合并了进来，然而这片大陆上生存的，居然是修仙界的人类......\n请选择你的阵营：\n1. 加入修仙界\n2. 加入异世界"
		responseStr += "\n请注意，地区是不可更换的"
	case action == "加入修仙界":
		user := schemas.NewUser("username")
		user.Location = "修仙界" // 用户选择了加入修仙界
		responseStr = "\n请选择术修方向或者剑修方向\n法系和物理系的攻击和战斗方式有很大区别，如需具体了解请发送战斗相关"
	case action == "术修":
		user := schemas.NewUser("username")
		user.Vocation = "术修" // 用户在修仙界中选择了术修方向
		responseStr = "您已选择术修方向，输入生成角色获取您的初始信息"
	case action == "剑修":
		user := schemas.NewUser("username")
		user.Vocation = "剑修"
		responseStr = "您已选择剑修方向，输入生成角色获取您的初始信息"
	case action == "加入异世界":
		user := schemas.NewUser("username")
		user.Location = "异世界" // 用户选择了加入修仙界
		responseStr = "\n请选择魔法路线和剑士路线\n法系和物理系的攻击和战斗方式有很大区别，如需具体了解请发送战斗相关"
	case action == "魔法":
		user := schemas.NewUser("username")
		user.Vocation = "魔法师" // 用户在异世界中选择了魔法路线
		responseStr = "您已选择魔法路线，输入生成角色获取您的初始信息"
	case action == "剑士":
		user := schemas.NewUser("username")
		user.Vocation = "剑士" // 用户在异世界中选择了剑士路线
		responseStr = "您已选择剑士路线，输入生成角色获取您的初始信息"
	case action == "生成角色":
		// 使用schemas.User来创建新角色
		newUser := schemas.NewUser("username")
		// 根据之前的选择生成角色信息
		responseStr = fmt.Sprintf("创建角色成功，您是第%d位进入AU界的玩家，您的角色名称为%s，是%s的%s，诞生于公元%s年，你从母亲怀中降生之日，AU大陆的光芒赐福于您，获得了初始命脉：%s，输入菜单进入游戏主界面",
			newUser.Uid, newUser.Nickname, newUser.Location, newUser.Vocation, time.Now().Format("2006"), newUser.Meridian)
	case action == "个人信息":
		responseStr = "\n地区：\n职业：\n名称：\n战力：\n等级： \npower/修为：\n力量/真气：\n敏捷/灵气：\n防御/元气：\n武器：\n防具：\n项链/护符： \n心法： \n技能列表：\n \n \n金币: \n店铺id："
	case action == "战斗相关":
		responseStr = s.IDescription.FightDescription()
	case action == "签到":
		responseStr = s.IDate.EnterDate()
	case utils.InArray(action, []string{"领取内测专属奖励", "内测奖励", "领取内测奖励"}):
		responseStr = "恭喜道友获得由453411753内测群发出的内测专属奖励，内测专属称号：AU仙人，持有效果：幸运值+1"
	case action == "闭关":
		newUser := schemas.NewUser("username")

		var cultivationLocation string
		if newUser.Location == "修仙界" {
			cultivationLocation = "丹塔"
		} else if newUser.Location == "异世界" {
			cultivationLocation = "天冠山"
		} else {
			responseStr = "您还未选择加入的地区，请先选择加入修仙界或异世界。"
			break
		}
		responseStr = fmt.Sprintf("未选择闭关地点，已为您默认选择初始地区%s", cultivationLocation)
	case action == "进入迷宫":
		responseStr = s.IDungeon.EnterDungeon()

	case action == "修炼":
		responseStr = s.ITraining.Training()
	}

	return operations.NewActionOK().WithPayload(responseStr)
}
