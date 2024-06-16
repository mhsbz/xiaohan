package schemas

//本文档为属性定义值，用户数据设定
import "github.com/mhsbz/xiaohan/pkg/utils"

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
	MM       string `json:"mM"`
}

func NewUser(MemberID string) *User {
	return &User{
		MemberID: MemberID,
		Uid:      utils.GenerateUID(),
		Nickname: utils.GenerateRandomChinese(),
		MM:       utils.GenerateRandomChinese(),
		Level:    1,
		Rank:     1,
		HP:       100,
		MP:       100,
		Power:    0,
	}
}
