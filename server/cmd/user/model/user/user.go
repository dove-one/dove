/**
 * @Author: zzy
 * @Description:
 * @File: user.go
 * @Version: 1.0.0
 * @Date: 2023/5/6 14:01
 */

package user

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	Id        uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id" form:"id"`                                                //用户唯d
	Uuid      string `gorm:"column:uuid;type:varchar(255) default '' comment '用户唯一uuid'" json:"uuid" form:"uuid"`                   //用户唯一uuid
	Role      uint64 `gorm:"column:role;type:bigint default '0' comment '角色id'" json:"role" form:"role"`                            //角色id
	Account   string `gorm:"column:account;type:varchar(255) default '' comment '用户登录账号'" json:"account" form:"account"`            //用户登录账号
	Password  string `gorm:"column:password;type:varchar(255) default '' comment '用户密码（bcrypt加密）'" json:"password" form:"password"` //用户密码（bcrypt加密）
	Contacts  string `gorm:"column:contacts;type:varchar(255) default '' comment '联系人姓名'" json:"contacts" form:"contacts"`          //联系人姓名
	Mobile    string `gorm:"column:mobile;type:varchar(255) default '' comment '联系电话'" json:"mobile" form:"mobile"`                 //联系电话
	Email     string `gorm:"column:email;type:varchar(255) default '' comment '电子邮箱'" json:"email" form:"email"`                    //电子邮箱
	Enable    uint8  `gorm:"column:enable;type:tinyint default '1' comment '是否启用 ：1/0 启用、禁用'" json:"enable" form:"enable"`          //是否启用 ：1/0 启用、禁用
	IsDel     uint8  `gorm:"column:isDel;type:tinyint default '0' comment '是否删除：1/0 已删除、未删除'" json:"is_del" form:"isDel"`           //是否删除：1/0 已删除、未删除
	CreatedAt int64  `gorm:"column:createdAt;type:bigint default '0' comment '创建时间'" json:"created_at" form:"createdAt"`            //创建时间
	UpdatedAt int64  `gorm:"column:updatedAt;type:bigint default '0' comment '修改时间'" json:"updated_at" form:"updatedAt"`            //修改时间
}

func (User) TableName() string {
	return "user"
}

func InitTable(db *gorm.DB) {
	table := &User{}
	err := db.AutoMigrate(table)
	if err != nil {
		//fmt.Println("user 初始化建表，失败：", err.Error())
		return
	}
	db.Exec("ALTER TABLE " + table.TableName() + " COMMENT '用户表';")
	//设置数据库最大连接数
	db.Exec("set GLOBAL MAX_CONNECTIONS=5000;")
}

type ManagerUser struct {
	salt string
	db   *gorm.DB
}

// NewUserManager creates a mysql manager.
func NewUserManager(db *gorm.DB, salt string) *ManagerUser {
	m := db.Migrator()
	if !m.HasTable(&User{}) {
		if err := m.CreateTable(&User{}); err != nil {
			panic(err)
		}
	}
	return &ManagerUser{
		db:   db,
		salt: salt,
	}
}

type MysqlManagerUser interface {
	GetUsersByCondition(ctx context.Context, opts ...Option) ([]*User, error)
}

type Option func(db *gorm.DB)

func Uuid(userID int64) Option {
	return func(db *gorm.DB) {
		db.Where("`uuid` = ?", userID)
	}
}

func (m *ManagerUser) GetUsersByCondition(ctx context.Context, opts ...Option) ([]*User, error) {
	for i := range opts {
		opts[i](m.db)
	}
	var users []*User
	if err := m.db.Model(&User{}).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
