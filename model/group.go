package model

import "time"

//群组
type Group struct {
	Id int64 `json:"-"`
	//群组名称
	Name string `json:"name" xorm:"varchar(32) name comment('name')"`
	//群组id
	GroupId string `json:"group_id" xorm:"varchar(32) group_id comment('group_id')"`
	//群组所有者
	OwnerUserId string `json:"owner_user_id" xorm:"varchar(32) owner_user_id comment('owner_user_id')"`
	//sessionId
	SessionId string `json:"session_id" xorm:"varchar(32) session_id comment('session_id')"`
	//创建时间
	CreatedAt time.Time `xorm:"timestamp  created comment('created_at')"`
	//描述
	Description string `json:"description" xorm:"varchar(256)   comment('description')"`
	//群组头像
	Avatar string `json:"avatar" xorm:"varchar(128)   comment('avatar')"`
	//群组类型  0 默认群
	Type int `json:"type" xorm:"tinyint  'type' comment('type ')"`
}
