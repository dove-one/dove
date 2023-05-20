/**
 * @Author: zzy
 * @Description:
 * @File: init.go
 * @Version: 1.0.0
 * @Date: 2023/5/6 14:00
 */

package model

import (
	"github.com/dove-one/dove/server/cmd/user/model/user"
	"gorm.io/gorm"
)

func InitTable(db *gorm.DB) {
	user.InitTable(db)
}
