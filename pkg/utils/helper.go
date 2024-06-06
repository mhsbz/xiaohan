package utils

import (
	"math/rand"
	"time"
)

// GenerateUID 生成一个唯一的整数UID
func GenerateUID() int {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机整数
	uid := rand.Intn(100000000)

	// 返回生成的UID
	return uid
}

// GenerateRandomChinese 根据ASCII码生成随机汉字组合
func GenerateRandomChinese() string {
	str := "汉皇重色思倾国御宇多年求不得杨家有女初长成养在深闺人未识天生丽质难自弃一朝选在君王侧回眸一笑百媚生六宫粉黛无颜色" +
		"春寒赐浴华清池温泉水滑洗凝脂侍儿扶起娇无力" +
		"始是新承恩泽时云鬓花颜金步摇芙蓉帐暖度春宵春宵苦短日高起从此君王不早朝" +
		"承欢侍宴无闲暇春从春游夜专夜后宫佳丽三千人三千宠爱在一身金屋妆" +
		"成娇侍夜玉楼宴罢醉和春姊妹弟兄皆列土可怜光彩生门户遂令" +
		"天下父母心不重生男重生女骊宫高处入青云仙乐风飘处处闻缓歌慢" +
		"舞凝丝竹尽日君王看不足渔阳鼙鼓动地来惊破霓裳羽衣曲九重城阙烟" +
		"尘生千乘万骑西南行翠华摇摇行复止西出都门百余里六军不发无奈何宛" +
		"转蛾眉马前死花钿委地无人收翠翘金雀玉搔头君王掩面救不得回看血泪相" +
		"和流黄埃散漫风萧索云栈萦纡登剑阁峨嵋山下少人行旌旗无光日色薄" +
		"蜀江水碧蜀山青圣主朝朝暮暮情行宫见月伤心色夜雨闻铃肠断声天旋地" +
		"转回龙驭到此踌躇不能去马嵬坡下泥土中不见玉颜空死处君臣相顾尽" +
		"沾衣东望都门信马归归来池苑皆依旧太液芙蓉未央柳芙蓉如面柳如眉" +
		"对此如何不泪垂春风桃李花开日秋雨梧桐叶落时西宫南内多秋草落叶满阶红不扫梨园弟子白发新椒房阿监青娥老夕殿萤飞思悄然" +
		"孤灯挑尽未成眠迟迟钟鼓初长夜耿耿星河欲曙天鸳鸯瓦冷霜华重翡翠衾寒谁与共悠悠生死别经年魂魄不曾来入梦临邛道士鸿都客能以精诚" +
		"致魂魄为感君王辗转思遂教方士殷勤觅排空驭气奔如电升天入地求之遍上穷碧落下黄泉两处茫茫皆不见忽闻海上有仙山山在虚无缥缈间楼阁" +
		"玲珑五云起其中绰约多仙子中有一人字太真雪肤花貌参差是金阙西厢叩玉扃转教小玉报双成闻道汉家天子使九华帐里梦魂惊揽衣推枕起徘徊珠箔银屏迤逦开云鬓半偏新睡觉花" +
		"冠不整下堂来风吹仙袂飘飖举犹似霓裳羽衣舞玉容寂寞泪阑干梨花一枝春带雨含情凝睇谢君王一别音容两渺茫昭" +
		"阳殿里恩爱绝蓬莱宫中日月长回头下望人寰处不见长安见尘雾惟将旧物表深情钿合金钗寄将去钗留一股合一扇钗擘黄金合" +
		"分钿但教心似金钿坚天上人间会相见临别殷勤重寄词词中有誓两心知七月七日长生殿夜半无人私语时在天愿作比翼鸟在地愿为" +
		"连理枝天长地久有时尽此恨绵绵无绝期"

	rand.Seed(time.Now().UnixNano())

	runes := []rune(str)

	rint := rand.Intn(4)
	if rint == 0 {
		rint = 2
	}
	name := ""
	for i := 0; i < rint; i++ {
		name += string(runes[rand.Intn(len(runes)-1)])
	}
	return name
}
