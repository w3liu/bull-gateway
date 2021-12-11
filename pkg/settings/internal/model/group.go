package model

import "time"

// 分组
type Group struct {
	Id          int64     `xorm:"BIGINT(11) not null pk autoincr"`
	Name        string    `xorm:"VARCHAR(100) not null comment('名称')"`
	Description string    `xorm:"VARCHAR(255) comment('描述')"`
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
}

// 分组绑定域名
type Domain struct {
	Id        int64     `xorm:"BIGINT(11) not null pk autoincr"`
	GroupId   int64     `xorm:"BIGINT(11) not null index comment('分组id')"`
	Url       string    `xorm:"VARCHAR(255) not null comment('域名url')"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
