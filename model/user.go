package model

import "time"

type User struct {
	Id        int64     `json:"-"`
	Account   string    `json:"-" xorm:"varchar(25) not null unique 'account' comment('account')"`
	Password  string    `json:"-" xorm:"varchar(256)  not null comment('password')"`
	NickName  string    `json:"nick_name" xorm:"varchar(32)   comment('nickname')"`
	Status    string    `json:"status" xorm:"varchar(256)   comment('status')"`
	Cover     string    `json:"cover" xorm:"varchar(32)   comment('cover')"`
	Avatar    string    `json:"avatar" xorm:"varchar(128)   comment('avatar')"`
	Email     string    `json:"email" xorm:"varchar(32) not null  comment('email')"`
	UserId    string    `json:"user_id" xorm:"varchar(32)  not null user_id comment('userId')"`
	CreatedAt time.Time `json:"created_at" xorm:"timestamp  created comment('created_at')"`
	//是否可以被搜索到
	AllowSearch bool `json:"allow_search"  xorm:"tinyint(1)  comment('allow_search')"`
	//是否接收匿名消息
	AllowAnonSession bool `json:"allow_anon_session"  xorm:"tinyint(1)   comment('allow_anon_session')"`
}
