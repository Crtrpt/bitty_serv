package model

import "time"

type UserToken struct {
	Id int64
	//用户id
	UserId string `xorm:"varchar(32) not null 'user_id' comment('userId')"`
	//用户token
	Token string `xorm:"varchar(32) not null unique 'token' comment('token')"`

	CreatedAt time.Time `xorm:"timestamp  created comment('created_at')"`
}
