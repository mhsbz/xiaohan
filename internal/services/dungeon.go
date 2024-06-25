package services

import (
	"fmt"
	"math/rand"
)

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

type DungeonService struct{}

// 新增处理迷宫功能的函数
func (s *DungeonService) EnterDungeon() string {
	var responseStr string
	currentFloor := 1 // 当前迷宫层数，初始为第一层
	// 循环处理每一层，直到玩家选择不继续探索
	for {
		// 计算本层怪物数量
		monsterCount := 5 + (currentFloor - 1)

		// 击杀怪物统计
		killedMonsters := make(map[string]int)
		for i := 0; i < monsterCount; i++ {
			selectedMonster := selectMonsterBasedOnProbability(monsters)
			killedMonsters[selectedMonster.Name]++
		}

		// 怪物掉落处理
		var uniqueDrops []string
		for monsterName := range killedMonsters {
			monster := findMonsterByName(monsters, monsterName)
			if rand.Float64() <= monster.DropRate {
				var dropItem string
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
					dropItem = "失落的人族圣女"
				}
				uniqueDrops = append(uniqueDrops, fmt.Sprintf("获得%s", dropItem))
			}
		}
	}
	return responseStr
}
