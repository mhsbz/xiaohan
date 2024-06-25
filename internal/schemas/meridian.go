package schemas

import "math/rand"

var MeridianMap = map[string]float64{
	"轮回觉醒者": 0.005,
	"怪盗基德":  0.01,
	"天道轮回":  0.0985 / 8,
	"阴阳逆转":  0.0985 / 8,
	"你是主角":  0.0985 / 8,
	"武神":    0.0985 / 8,
	"贵族":    0.0985 / 8,
	"终极反派":  0.0985 / 8,
	"肾虚子":   0.0985 / 8,
	"Saber": 0.0985 / 8,
}

func randomMeridian() string {
	// 计算所有命脉的总概率
	totalProbability := 0.0
	for _, v := range MeridianMap {
		totalProbability += v
	}

	// 生成一个介于0到总概率之间的随机数
	randomNum := rand.Float64() * totalProbability

	// 遍历命脉列表，累加概率直到找到对应的命脉
	accumulatedProb := 0.0
	for k, v := range MeridianMap {
		accumulatedProb += v
		if randomNum <= accumulatedProb {
			return k
		}
	}
	// 理论上不会走到这里，但作为一个安全措施，如果没有匹配到任何命脉，则返回一个默认值
	return "贵族"
}
