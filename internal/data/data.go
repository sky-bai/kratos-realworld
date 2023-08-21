package data

import (
	"gorm.io/gorm"
	"kratos-realworld/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewUserRepo, NewProfileRepo)

// Data .
type Data struct {
	// TODO wrapped database client 1.填充所要用到的数据库连接对象 2.实现对象的New方法 3.注册到wire里面
	db *gorm.DB
}

func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// todo 这里还需要做具体的mysql配置
	return db
}

// 初始化配置文件的时候的处理 如果没有初始化成功就全部应该报错
// 定义好每一个依赖模块
// 数据层也是一个依赖模块 gorm redis 等等 所依赖的基础设施都是模块 然后通过wire去做依赖注入
// 所以说依赖的模块 需要配置两个模块 一个是依赖的配置对象 nil --> 配置对象 一个是依赖的实例化对象 函数签名 配置对象 --> 实例化对象
