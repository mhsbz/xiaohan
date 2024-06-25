package schemas

//本文档为属性定义值，用户数据设定
import (
	"github.com/mhsbz/xiaohan/pkg/utils"
)

// 定义一个新用法sync/atomic搜集用户id

type User struct {
	MemberID string `json:"member_id"`
	Uid      int    `json:"uid"`
	Nickname string `json:"nickname"`
	Level    int    `json:"level"`
	Rank     int    `json:"rank"`
	HP       int    `json:"hp"`
	MP       int    `json:"mp"`
	Power    int    `json:"power"`
	XStatus  bool   `json:"x_status"`
	Meridian string `json:"meridian"`
	Location string `json:"location"`
	Vocation string `json:"Vocation"`
}

func NewUser(MemberID string) *User {
	uidCounter++
	return &User{
		MemberID: MemberID,
		Uid:      uidCounter,
		Nickname: utils.GenerateRandomChinese(),
		Meridian: randomMeridian(),
		Level:    1,
		Rank:     1,
		HP:       100,
		MP:       100,
		Power:    0,
		XStatus:  false,
		Location: "AU",
		Vocation: "AU",
	}
}

var uidCounter int = 1
