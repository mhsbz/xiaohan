package services

// 本文档为指令系统
import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/mhsbz/xiaohan/api/gen/xiaohan/server/operations"
	"github.com/mhsbz/xiaohan/internal/repository"
	"github.com/mhsbz/xiaohan/internal/schemas"
	"github.com/mhsbz/xiaohan/pkg/utils"
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

// 修炼选项
type CultivationOption struct {
	Description string
	Probability int // 这里用int类型简化概率表示，实际应用中应使用float类型以精确表示概率
	Duration    int
}

// 修炼选项列表
var cultivationOptions = []CultivationOption{
	{Description: "踏上云游修炼之旅（30分钟）", Probability: 80, Duration: 30},
	{Description: "被一个神秘的镜子吸入镜之迷宫（60分钟）", Probability: 5, Duration: 60},
	{Description: "被殿元山的神秘结界吸入反转世界（120分钟）", Probability: 5, Duration: 120},
	{Description: "在千年冰山芸冰山脚下，遇到一位垂钓老者（60分钟）", Probability: 5, Duration: 60},
	{Description: "被迫杀至悬崖，意外跌入一座古老的墓穴（120分钟）", Probability: 5, Duration: 120},
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
func specialEvent() string {
	randNum := rand.Float64()
	switch {
	case randNum <= 0.3: // 30%概率空手而返
		return "走到迷宫的尽头，空手而返。"
	case randNum <= 0.5: // 20%概率进入神秘的机关道
		return "进入神秘的机关道，获得一个【神秘宝箱】。"
	case randNum <= 0.7: // 20%概率发现前人遗留的尸骨
		return "发现了前人遗留的尸骨，一个包我品三遍，每次都有新发现，获得了【前人的尸骨-绿】。"
	default: // 剩余30%概率发现传送门
		return "发现了进入深层迷宫的传送门入口，解锁下一层！"
	}
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

	switch {

	case action == "":
		responseStr = "当你召唤我的时候，你的路就只有一条，加入赛博修仙界,输入“加入异世界修仙”进入游戏"
	case action == "加入异世界修仙":
		responseStr = "那是一个风雨交加的晚上，你在公司加完班迎着雨滴走在马路上，一辆刹车打滑的大卡车创向了你，一串串回马灯般的画面闪入你的脑海中，但是你并没有死，你降临到了一个名为AU的大陆，这片大陆本身是一个异世界大陆，几百年前，世界上最顶尖的巫术师聚集在了一起，不知是何原因，创造了一个超大的魔法阵，将另一个位面的大陆拖拽合并了进来，然而这片大陆上生存的，居然是修仙界的人类......\n请选择你的阵营：\n1. 加入修仙界\n2. 加入异世界"
		responseStr += "请注意，地区是不可更换的"
	case action == "加入修仙界":
		s.a = "修仙界" // 用户选择了加入修仙界
		responseStr = "请选择术修方向或者剑修方向————法系和物理系的攻击和战斗方式有很大区别，如需具体了解请发送战斗相关"
	case action == "术修":
		s.b = "术修" // 用户在修仙界中选择了术修方向
		responseStr = "您已选择术修方向，输入生成角色获取您的初始信息"
	case action == "剑修":
		s.b = "剑修" // 用户在修仙界中选择了剑修方向
		responseStr = "您已选择剑修方向，输入生成角色获取您的初始信息"
	case action == "加入异世界":
		s.a = "异世界" // 用户选择了加入异世界
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
		initialMeridian := generateMeridian()
		responseStr = fmt.Sprintf("创建角色成功，您是第%d位进入AU界的玩家，您的角色名称为%s，是%s的%s，诞生于公元%s年，你从母亲怀中降生之日，AU大陆的光芒赐福于您，获得了初始命脉：%s，输入菜单进入游戏主界面",
			newUser.Uid, newUser.Nickname, s.a, combineBCValue(s.b, s.c), time.Now().Format("2006"), initialMeridian)
	case action == "战斗相关":
		responseStr = "本文档为法系和物理系的对战系统注释\n"
		responseStr += "术修，法师的攻击方式为技能攻击，优先默认释放1号主技能，后续按排序释放2，3号技能，绝技/奥义技每场战斗不限次数，消耗一定百分比蓝量之后下回合判定释放，需要消耗非常庞大的蓝量若蓝量不足则回到默认施法顺序，当蓝条不足以支撑任何技能释放的时候，将进行普通攻击\n"
		responseStr += "剑修，剑士的攻击方式为平a攻击，攻击成功的情况下生成一点能量，每受到10%最大生命值伤害也会生成一点能量，回合开始当能量足够释放技能的情况下1＞2＞3的顺序轮次进行释放，绝技/觉醒技每场战斗限一次，累积生成共计10点以上能量之后，回合开始必定释放，不消耗能量"
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
		var cultivationLocation string
		if s.a == "修仙界" {
			cultivationLocation = "丹塔"
		} else if s.a == "异世界" {
			cultivationLocation = "天冠山"
		} else {
			responseStr = "您还未选择加入的地区，请先选择加入修仙界或异世界。"
			break
		}
		responseStr = fmt.Sprintf("未选择闭关地点，已为您默认选择初始地区%s", cultivationLocation)
	case action == "进入迷宫":
		return s.enterDungeon(params)

	case action == "修炼":
		// 随机选择修炼选项
		rand.Seed(time.Now().UnixNano()) // 确保每次运行都有不同的随机结果
		totalProbability := 0
		for _, option := range cultivationOptions {
			totalProbability += option.Probability
		}
		randomNum := rand.Intn(totalProbability) // 生成一个0到totalProbability之间的随机数

		// 选择修炼
		for _, option := range cultivationOptions {
			randomNum -= option.Probability
			if randomNum < 0 {
				responseStr = fmt.Sprintf(" %s，预计时长：%d 分钟。",
					option.Description, option.Duration)
				break
			}
		}
		if responseStr == "" {
			//理论上不应该到达这里，但作为保险措施，提供一个默认响应
			responseStr = "修炼过程中发生未知错误，请稍后重试。"
		}
		//判定已经开始的，我先不动让墨寒来操作
		//return operations.NewActionOK().WithPayload(responseStr)
		//user, err := s.CreateOrGetUser(params.MemberID)
		//if err != nil {
		//	return operations.NewActionInternalServerError().WithPayload("Internal Server Error")
		//}
		//if user.XStatus {
		//	responseStr = "你已经在修炼中"
		//}
		//responseStr = "已经开始修炼"

	}

	return operations.NewActionOK().WithPayload(responseStr)
}

// 生成角色时的假设的辅助函数来组合b和c的值，根据实际需要实现
func combineBCValue(b string, c string) string {
	return fmt.Sprintf("%s%s", b, c)
}

// 命脉及其概率定义
var meridians = []struct {
	Name        string
	Probability float64
}{
	{"轮回觉醒者", 0.005},
	{"怪盗基德", 0.01},
	{"天道轮回", 0.0985 / 8},
	{"阴阳逆转", 0.0985 / 8},
	{"你是主角", 0.0985 / 8},
	{"武神", 0.0985 / 8},
	{"贵族", 0.0985 / 8},
	{"终极反派", 0.0985 / 8},
	{"肾虚子", 0.0985 / 8},
	{"Saber", 0.0985 / 8},
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

// 新增处理迷宫功能的函数
func (s *Service) enterDungeon(params operations.ActionParams) middleware.Responder {
	// 直接设定为第一层
	floor := 1
	// 计算本层怪物数量
	monsterCount := 5 + (floor - 1)

	// 击杀怪物统计
	killedMonsters := make(map[string]int)
	for i := 0; i < monsterCount; i++ {
		selectedMonster := selectMonsterBasedOnProbability(monsters)
		killedMonsters[selectedMonster.Name]++
	}

	// 怪物掉落处理
	drops := []string{}
	for monsterName, count := range killedMonsters {
		monster := findMonsterByName(monsters, monsterName)
		if rand.Float64() <= monster.DropRate {
			dropItem := "未知物品"
			switch monster.Name {
			case "哥布林战士":
				dropItem = "装备1"
			case "哥布林巫师":
				dropItem = "道具1"
			case "哥布林弓箭手":
				if rand.Intn(2) == 0 {
					dropItem = "道具1"
				} else {
					dropItem = "装备1"
				}
			case "哥布林头目":
				dropItem = "稀有物品"
			}
			drops = append(drops, fmt.Sprintf("击杀%d只%s，获得%s", count, monster.Name, dropItem))
		}
	}

	// 特殊事件判定
	eventOutcome := specialEvent()

	// 结算信息
	responseStr := fmt.Sprintf("在第%d层迷宫，您击杀了:", floor)
	for monster, count := range killedMonsters {
		responseStr += fmt.Sprintf("%d只%s，", count, monster)
	}
	if len(drops) > 0 {
		responseStr += "掉落物品：" + strings.Join(drops, "") + "，"
	}
	responseStr += eventOutcome

	// 如果到下一关终点额外发送询问
	if strings.Contains(eventOutcome, "解锁下一层") {
		responseStr += "是否直接探索下一层？(是/否)"
	}
	return operations.NewActionOK().WithPayload(responseStr)
}
