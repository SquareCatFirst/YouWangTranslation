package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 配置结构体
type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	Schema   string
}

type JWTConfig struct {
	Secret string
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

// 全局可访问的配置和数据库连接
var (
	Cfg *Config
	DB  *gorm.DB
)

// init 在包导入时自动执行一次
func init() {
	// 1. 加载配置
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("解析配置失败: %v", err)
	}
	Cfg = &cfg

	// 2. 连接数据库
	dbConf := Cfg.Database
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.DBName, dbConf.Schema,
	)

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	DB = db

	fmt.Println("配置加载完成，数据库连接成功，使用 schema:", dbConf.Schema)
}
