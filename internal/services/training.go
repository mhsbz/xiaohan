package services

import (
	"fmt"
	"math/rand"
	"time"
)

type TrainingService struct{}

type TrainingOption struct {
	Description string  `json:"description"` // 修炼选项描述
	Duration    int     `json:"duration"`    // 修炼选项持续时间秒
	Probability float64 `json:"probability"` // 修炼选项概率
}

var trainingOptions = []*TrainingOption{
	{Description: "踏上云游修炼之旅（30分钟）", Duration: 30, Probability: 0.8},
	{Description: "被一个神秘的镜子吸入镜之迷宫（60分钟）", Duration: 60, Probability: 0.05},
	{Description: "被殿元山的神秘结界吸入反转世界（120分钟）", Duration: 120, Probability: 0.05},
	{Description: "在千年冰山芸冰山脚下，遇到一位垂钓老者（60分钟）", Duration: 60, Probability: 0.05},
	{Description: "被迫杀至悬崖，意外跌入一座古老的墓穴（120分钟）", Duration: 120, Probability: 0.05},
}

func randomTraining() *TrainingOption {
	rand.Seed(time.Now().UnixNano())
	// 计算总概率
	totalProbability := 0.0
	for _, item := range trainingOptions {
		totalProbability += item.Probability
	}

	// 生成一个0到总概率之间的随机数
	randomNum := rand.Float64() * totalProbability

	// 根据随机数选择结构体
	for _, item := range trainingOptions {
		randomNum -= item.Probability
		if randomNum < 0 {
			return item
		}
	}
	return trainingOptions[0]
}

func (s *TrainingService) Training() string {
	option := randomTraining()
	responseStr := fmt.Sprintf(" %s，预计时长：%d 分钟。", option.Description, option.Duration)
	return responseStr
}
