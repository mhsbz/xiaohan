package schemas

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
}

func NewUser(MemberID string) *User {
	return &User{
		MemberID: MemberID,
		Uid:      utils.GenerateUID(),
		Nickname: utils.GenerateRandomChinese(2),
		Level:    1,
		Rank:     1,
		HP:       100,
		MP:       100,
		Power:    0,
	}
}
