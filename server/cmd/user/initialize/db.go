/**
 * @Author: zzy
 * @Description:
 * @File: db.go
 * @Version: 1.0.0
 * @Date: 2023/5/6 14:07
 */

package initialize

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dove-one/dove/server/cmd/user/config"
	"github.com/dove-one/dove/server/common/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// InitDB to init database
func InitDB() *gorm.DB {
	c := config.GlobalServerConfig.MysqlInfo
	dsn := fmt.Sprintf(consts.MySqlDSN, c.User, c.Password, c.Host, c.Port, c.Name)
	//newLogger := logger.New(
	//	logrus.NewWriter(), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second,   // Slow SQL Threshold
	//		LogLevel:      logger.Silent, // Log level
	//		Colorful:      true,          // Disable color printing
	//	},
	//)

	// global mode
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//Logger: newLogger,
	})
	if err != nil {
		klog.Fatalf("init gorm failed: %s", err.Error())
	}

	//if err = db.Use(tracing.NewPlugin()); err != nil {
	//	klog.Fatalf("use tracing plugin failed: %s", err.Error())
	//}
	return db
}
