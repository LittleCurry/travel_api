package main

import (
	"flag"
	"fmt"
	//"git.iguiyu.com/park/api/configs"
	"github.com/curry/travel_api/configs"
	//"git.iguiyu.com/park/api/db"
	"github.com/curry/travel_api/db"
	//"git.iguiyu.com/park/api/rpc_client"
	"git.iguiyu.com/park/misc/helpers"
)

/**
 * 解析参数
 */
func parseFlag() {
	flag.Parse()
	if *configFile == "" {
		fmt.Println("place provide the configuration file path!")
		fmt.Println("Usage:\n --config configpath [--env prod]")
	}
}

/**
 * 载入配置文件
 */
func loadConfig() {
	helpers.LoadConfigAndSetupEnv(configs.AppConf, *env, *configFile)
}

/**
 * 初始化orm引擎
 */
func initOrm() {
	db.InitMySQL(configs.AppConf.DbDsn)
}

/**
 * 初始化redis
 */
func initRedis() {
	db.InitRedis(configs.AppConf.RedisAddr)
}

/**
 * 初始化rpc
 */
func initRpc() {

}
