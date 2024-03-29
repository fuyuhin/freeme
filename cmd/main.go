package main

import (
	_ "freeme/business/controllers"
	_ "freeme/business/repositorys"
	"freeme/components/config"
	"github.com/8treenet/freedom"
	"github.com/8treenet/freedom/middleware"
	"github.com/8treenet/gcache"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kataras/iris"
)

func main() {
	/*
		installRedis() //安装redis
		installLogrus() //安装第三方logger
		http2 h2c 服务
		h2caddrRunner := freedom.CreateH2CRunner(config.Get().App.Other["listen_addr"].(string))
	*/
	app := freedom.NewApplication()
	app.InstallGorm(gormInit)
	app.InstallMiddleware(middleware.NewTrace("TRACE-ID"))
	app.InstallMiddleware(middleware.NewLogger("TRACE-ID", true))
	app.InstallMiddleware(middleware.NewRuntimeLogger("TRACE-ID"))
	app.Run(iris.Addr(config.Get().App.Other["listen_addr"].(string)), *config.Get().App)
}

func gormInit() (db *gorm.DB, cache gcache.Plugin) {
	// conf := config.Get().DB
	var e error
	// db, e = gorm.Open("mysql", conf.Addr)
	db, e = gorm.Open("sqlite3", "./chinook.db")
	if e != nil {
		freedom.Logger().Fatal(e.Error())
	}
	freedom.Logger().Info("Database connected.")
	// db.DB().SetMaxIdleConns(conf.MaxIdleConns)
	// db.DB().SetMaxOpenConns(conf.MaxOpenConns)
	// db.DB().SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeTime) * time.Second)

	/*
		启用缓存中间件
		cfg := config.Get().DB.Cache
		ropt := gcache.RedisOption{
			Addr:               cfg.Addr,
			Password:           cfg.Password,
			DB:                 cfg.DB,
			MaxRetries:         cfg.MaxRetries,
			PoolSize:           cfg.PoolSize,
			ReadTimeout:        time.Duration(cfg.ReadTimeout) * time.Second,
			WriteTimeout:       time.Duration(cfg.WriteTimeout) * time.Second,
			IdleTimeout:        time.Duration(cfg.IdleTimeout) * time.Second,
			IdleCheckFrequency: time.Duration(cfg.IdleCheckFrequency) * time.Second,
			MaxConnAge:         time.Duration(cfg.MaxConnAge) * time.Second,
			PoolTimeout:        time.Duration(cfg.PoolTimeout) * time.Second,
		}
		opt := gcache.DefaultOption{}
		opt.Expires = cfg.Expires      //缓存时间，默认60秒。范围 30-900
		opt.Level = gcache.LevelSearch //缓存级别，默认LevelSearch。LevelDisable:关闭缓存，LevelModel:模型缓存， LevelSearch:查询缓存
		//缓存中间件 注入到Gorm
		cache = gcache.AttachDB(db, &opt, &ropt)
	*/
	return
}

// func installRedis() {
// 	freedom.InstallRedis(func() (client *redis.Client) {
// 		cfg := config.Get().Redis
// 		opt := &redis.Options{
// 			Addr:               cfg.Addr,
// 			Password:           cfg.Password,
// 			DB:                 cfg.DB,
// 			MaxRetries:         cfg.MaxRetries,
// 			PoolSize:           cfg.PoolSize,
// 			ReadTimeout:        time.Duration(cfg.ReadTimeout) * time.Second,
// 			WriteTimeout:       time.Duration(cfg.WriteTimeout) * time.Second,
// 			IdleTimeout:        time.Duration(cfg.IdleTimeout) * time.Second,
// 			IdleCheckFrequency: time.Duration(cfg.IdleCheckFrequency) * time.Second,
// 			MaxConnAge:         time.Duration(cfg.MaxConnAge) * time.Second,
// 			PoolTimeout:        time.Duration(cfg.PoolTimeout) * time.Second,
// 		}
// 		client = redis.NewClient(opt)
// 		if e := client.Ping().Err(); e != nil {
// 			freedom.Logger().Fatal(e.Error())
// 		}
// 		return
// 	})
// }

// func installLogrus() {
// 	logrus.SetLevel(logrus.InfoLevel)
// 	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"})
// 	freedom.Logger().Install(logrus.StandardLogger())
// }
