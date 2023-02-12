package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
) //viper用于配置

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")    //配置文件名
	viper.AddConfigPath("config") //配置文件所在的路径
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))     //获取yaml中的属性"app"
	fmt.Println("config mysql:", viper.Get("mysql")) //获取yaml中的属性"mysql"
}
func InitMySQL() (err error) {
	println(viper.GetString("mysql.dns"))
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return nil
}
