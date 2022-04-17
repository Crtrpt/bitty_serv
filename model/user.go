package model

import "time"

type User struct {
	Id       int64
	Account  string `xorm:"varchar(25) not null unique 'account' comment('account')"`
	Password string `xorm:"varchar(25)  not null comment('password')"`
	NickName string `xorm:"varchar(25)   comment('nickname')"`
	UserId   string `xorm:"varchar(32) not null unique comment('userId')"`

	CreatedAt time.Time `xorm:"timestamp  created comment('created_at')"`
}
