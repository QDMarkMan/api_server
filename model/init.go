package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

// DataBase for different env
type DataBase struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *DataBase

// Init method for db
func (db *DataBase) Init() {
	fmt.Print(DB)
	DB = &DataBase{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
}

// Close database connect
func (db *DataBase) Close() {
	DB.Self.Close()
	DB.Docker.Close()
}

func setupDB(db *gorm.DB) {
	db.LogMode(true)
	db.DB().SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}
	log.Info("Connect Database success!")
	setupDB(db)
	return db
}

//
func InitSelfDB() *gorm.DB {
	return openDB(
		viper.GetString("DB.username"),
		viper.GetString("DB.password"),
		viper.GetString("DB.addr"),
		viper.GetString("DB.name"))
}
func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("DOCKER_DB.username"),
		viper.GetString("DOCKER_DB.password"),
		viper.GetString("DOCKER_DB.addr"),
		viper.GetString("DOCKER_DB.name"))
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}
