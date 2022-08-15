package model

import "time"

type UserToken struct {
	Id int64 `json:"-"`
	//用户id
	UserId string `xorm:"varchar(32) not null 'user_id' comment('userId')"`
	//用户token
	Token string `xorm:"varchar(32) not null unique 'token' comment('token')"`
	//最后登陆时间
	CreatedAt time.Time `xorm:"timestamp  created comment('created_at')"`
	//最后登陆ip地址
	LastLoginIp string `json:"last_login_ip" xorm:"varchar(32) not null  comment('last_login_ip')"`
	//windows 允许的平台和版本信息
	Platform string `json:"platform" xorm:"varchar(32) not null  comment('platform')"`
	//客户端版本
	ClientVersion string `json:"client_version" xorm:"varchar(32) not null  comment('client_version')"`
}
