package model

import "time"

//可以理解为某个用户的通讯录列表
type Session struct {
	Id        int64
	Avatar    string `json:"avatar" xorm:"varchar(32) avatar comment('avatar')"`
	Name      string `json:"name" xorm:"varchar(256) name comment('name')"`
	SessionId string `json:"session_id" xorm:"varchar(32) session_id comment('session_id')"`
	// session 类型 0是单独聊天
	Type   int    `json:"type" xorm:"tinyint  'type' comment('type')"`
	Config string `json:"config" xorm:"json config comment('config')"`
	//
	CreatedAt time.Time `json:"created_at" xorm:"timestamp  created comment('created_at')"`
}
