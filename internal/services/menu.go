package services

type IMenuLogic interface {
	MenuList() string
}

type MenuService struct{}

func (s *MenuService) MenuList() string {
	var responseStr string
	responseStr = "蓝字快捷正在申请中，暂时使用文本菜单\n\n"
	responseStr += "修炼： 修炼  闭关  进入迷宫  双修/补魔\n"
	responseStr += "战斗： 副本  切磋  登仙台  生死对决 \n"
	responseStr += "战斗： 炼金  炼药  异火  技能配置  \n"
	responseStr += "信息： 命脉更换  名称更换  个人信息  背包  仓储  状态\n"
	responseStr += "任务： 当前主线  已接取任务  日常任务 \n"
	responseStr += "交易： 店铺  云游商人  拍卖行  赠送金币\n"
	responseStr += "休闲： 赌场  挖矿 \n"
	responseStr += "帮助： 战斗相关  属性相关  道具表  装备表  命脉表  技能表\n\n"
	responseStr += "目前可用的功能：修炼 闭关 进入迷宫"
	return responseStr
}
