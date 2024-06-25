package services

import (
	"fmt"
	"math/rand"
)

type DateService struct{}

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

// Date 结构体的用途在此上下文中不清晰，若无特殊用途，可以考虑移除或修改其用途

func (s *DateService) EnterDate() string {
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
	var responseStr string
	if rewardGold > 0 {
		responseStr = fmt.Sprintf("签到成功！获得 %d 金币。", rewardGold)
	} else {
		responseStr = "签到失败，请稍后再试。"
	}
	return responseStr
}
