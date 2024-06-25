package services

type DescriptionService struct{}

// 新增处理描述功能的函数

func (s *DescriptionService) FightDescription() string {
	responseStr := "本文档为法系和物理系的对战系统注释\n\n"
	responseStr += "术修，法师的攻击方式为技能攻击，优先默认释放1号主技能，后续按排序释放2，3号技能，绝技/奥义技每场战斗不限次数，消耗一定百分比蓝量之后下回合判定释放，需要消耗非常庞大的蓝量若蓝量不足则回到默认施法顺序，当蓝条不足以支撑任何技能释放的时候，将进行普通攻击\n\n"
	responseStr += "剑修，剑士的攻击方式为平a攻击，攻击成功的情况下生成一点能量，每受到10%最大生命值伤害也会生成一点能量，回合开始当能量足够释放技能的情况下1＞2＞3的顺序轮次进行释放，绝技/觉醒技每场战斗限一次，累积生成共计10点以上能量之后，回合开始必定释放，不消耗能量"
	return responseStr
}
