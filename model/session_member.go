package model

import "time"

//可以理解为某个用户的通讯录列表
type SessionMember struct {
	Id            int64
	SessionId     string `json:"session_id" xorm:"varchar(32) session_id comment('session_id')"`
	UserId        string `json:"user_id" xorm:"varchar(32) user_id comment('user_id')"`
	Name          string `json:"name" xorm:"varchar(32) name comment('name')"`
	Avatar        string `json:"avatar" xorm:"varchar(32) avatar comment('avatar')"`
	SessionName   string `json:"session_name" xorm:"varchar(32) session_name comment('session_name')"`
	SessionAvatar string `json:"session_avatar" xorm:"varchar(256) session_avatar comment('session_avatar')"`

	//创建时间
	CreatedAt time.Time `xorm:"timestamp  created comment('created_at')"`
}
