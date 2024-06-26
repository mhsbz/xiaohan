package services

import "fmt"

// 任务数据结构
type Quest struct {
	ID          int
	Title       string
	Description string
	Objective   string
	Reward      string
}

// 可接受任务列表
var availableQuests = []Quest{
	{ID: 1, Title: "解救被困的人族圣女", Description: "深入迷宫，救出被囚禁的人族圣女。", Objective: "解救圣女", Reward: "1000金币, 声望值：10"},
	{ID: 2, Title: "击杀10只哥布林", Description: "清理附近的哥布林威胁，保护村庄的安全。", Objective: "击杀10只哥布林", Reward: "50金币, 哥布林耳朵x10，声望值：1"},
	{ID: 3, Title: "击杀哥布林头目3只", Description: "消灭哥布林的领导层，削弱它们的组织力量。", Objective: "击杀3只哥布林头目", Reward: "250金币, 稀有装备箱x1，声望值：2"},
}

type Daskservice struct{}

func (s *Daskservice) Dasks() string {
	var taskListStr string
	for _, quest := range availableQuests {
		taskListStr += fmt.Sprintf("任务ID: %d\n任务名称: %s\n任务描述: %s\n目标: %s\n奖励: %s\n",
			quest.ID, quest.Title, quest.Description, quest.Objective, quest.Reward)
	}

	// 将 responseStr 初始化为字符串
	responseStr := "当前可接受任务列表:\n" + taskListStr
	return responseStr
}
