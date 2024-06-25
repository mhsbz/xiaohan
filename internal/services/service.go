package services

// 本文档为指令系统
import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/mhsbz/xiaohan/api/gen/xiaohan/server/operations"
	"github.com/mhsbz/xiaohan/internal/schemas"
	"github.com/mhsbz/xiaohan/pkg/utils"
	"math/rand"
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
	b            string
	c            string
}

// 签到奖励选项
type SigninRewardOption struct {
	Gold   int
	Chance float64
}

// 签到奖励选项列表
var signinRewards = []SigninRewardOption{
	{Gold: 6, Chance: 0.02},
	{Gold: 30, Chance: 0.08},
	{Gold: 98, Chance: 0.3},
	{Gold: 198, Chance: 0.4},
	{Gold: 328, Chance: 0.15},
	{Gold: 648, Chance: 0.04},
	{Gold: 8888, Chance: 0.01},
}

// 怪物相关
type Monster struct {
	Name     string
	Level    int
	MeetProb float64 // 遭遇概率
	DropRate float64 // 掉落概率
}

var monsters = []Monster{
	{Name: "哥布林战士", Level: 1, MeetProb: 0.4, DropRate: 0.1},  // 40%
	{Name: "哥布林巫师", Level: 2, MeetProb: 0.3, DropRate: 0.1},  // 30%
	{Name: "哥布林弓箭手", Level: 3, MeetProb: 0.2, DropRate: 0.1}, // 20%
	{Name: "哥布林头目", Level: 4, MeetProb: 0.1, DropRate: 0.05}, // 10%
}

func findMonsterByName(monsters []Monster, name string) Monster {
	for _, monster := range monsters {
		if monster.Name == name {
			return monster
		}
	}

	// 如果找不到对应的名字，返回一个默认的 Monster 结构体实例
	return Monster{}
}

func selectMonsterBasedOnProbability(monsters []Monster) Monster {
	totalProb := 0.0
	for _, m := range monsters {
		totalProb += m.MeetProb
	}

	randomNum := rand.Float64() * totalProb
	accumulatedProb := 0.0
	for _, m := range monsters {
		accumulatedProb += m.MeetProb
		if randomNum <= accumulatedProb {
			return m
		}
	}
	// 理论上不会执行到这里，但作为安全措施返回第一个怪物
	return monsters[0]
}

// 特殊事件判定函数
func SpecialEvent() string {
	randNum := rand.Float64()
	switch {
	case randNum <= 0.3: // 30%概率空手而返
		return "\n- 迷宫结算：走到迷宫的尽头，空手而返。"
	case randNum <= 0.5: // 20%概率进入神秘的机关道
		return "\n- 特殊事件：进入神秘的机关道，获得一个【神秘宝箱】。"
	case randNum <= 0.7: // 20%概率发现前人遗留的尸骨
		return "\n- 特殊事件：发现了前人遗留的尸骨，一个包我品三遍，每次都有新发现，获得了【前人的尸骨-绿】。"
	default: // 剩余30%概率发现传送门
		return "\n- 迷宫结算：发现了进入深层迷宫的传送门入口，解锁下一层！"
	}
}

func NewService() *Service {
	return &Service{
		IUser:        &UserService{},
		ITraining:    &TrainingService{},
		IMenu:        &MenuService{},
		IDungeon:     &DungeonService{},
		IFight:       &FightService{},
		IDescription: &DescriptionService{},

		b: "",
		c: "",
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
		responseStr += "请注意，地区是不可更换的"
	case action == "加入修仙界":
		user := schemas.NewUser("username")
		user.Location = "修仙界" // 用户选择了加入修仙界
		responseStr = "请选择术修方向或者剑修方向————法系和物理系的攻击和战斗方式有很大区别，如需具体了解请发送战斗相关"
	case action == "术修":
		s.b = "术修" // 用户在修仙界中选择了术修方向
		responseStr = "您已选择术修方向，输入生成角色获取您的初始信息"
	case action == "剑修":
		s.b = "剑修" // 用户在修仙界中选择了剑修方向
		responseStr = "您已选择剑修方向，输入生成角色获取您的初始信息"
	case action == "加入异世界":
		user := schemas.NewUser("username")
		user.Location = "修仙界" // 用户选择了加入修仙界
		responseStr = "请选择魔法路线和剑士路线————法系和物理系的攻击和战斗方式有很大区别，如需具体了解请发送战斗相关"
	case action == "魔法":
		s.c = "魔法师" // 用户在异世界中选择了魔法路线
		responseStr = "您已选择魔法路线，输入生成角色获取您的初始信息"
	case action == "剑士":
		s.c = "剑士" // 用户在异世界中选择了剑士路线
		responseStr = "您已选择剑士路线，输入生成角色获取您的初始信息"
	case action == "生成角色":
		// 使用schemas.User来创建新角色
		newUser := schemas.NewUser("username")
		// 根据之前的选择生成角色信息
		responseStr = fmt.Sprintf("创建角色成功，您是第%d位进入AU界的玩家，您的角色名称为%s，是%s的%s，诞生于公元%s年，你从母亲怀中降生之日，AU大陆的光芒赐福于您，获得了初始命脉：%s，输入菜单进入游戏主界面",
			newUser.Uid, newUser.Nickname, newUser.Location, combineBCValue(s.b, s.c), time.Now().Format("2006"), newUser.Meridian)
	case action == "个人信息":
		responseStr = "\n地区：\n职业：\n名称：\n战力：\n等级： \npower/修为：\n力量/真气：\n敏捷/灵气：\n防御/元气：\n武器：\n防具：\n项链/护符： \n心法： \n技能列表：\n \n \n金币: \n店铺id："
	case action == "战斗相关":
		responseStr = s.IDescription.FightDescription()
	case action == "签到":
		// 生成一个0到1之间的随机数
		randomNum := rand.Float64()
		// 初始化累积概率
		accumulatedChance := 0.0
		// 遍历签到奖励列表，累加概率直到随机数小于等于当前累积概率
		var rewardGold int
		for _, reward := range signinRewards {
			accumulatedChance += reward.Chance
			if randomNum <= accumulatedChance {
				rewardGold = reward.Gold
				break
			}
		}

		// 将签到成功和获得金币的消息设置到responseStr中
		if rewardGold > 0 {
			responseStr = fmt.Sprintf("签到成功！获得 %d 金币。", rewardGold)
		} else {
			responseStr = "签到失败，请稍后再试。"
		}
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

// 生成角色时的假设的辅助函数来组合b和c的值，根据实际需要实现
func combineBCValue(b string, c string) string {
	return fmt.Sprintf("%s%s", b, c)
}
