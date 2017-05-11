package configs

import "git.iguiyu.com/park/misc/helpers"


var AppConf = new(appConf)

type appConf struct {
	helpers.BaseConfig
	DbDsn            string
	RedisAddr        string
	SmsGatewayServer string
}

func (this *appConf) GetName() string {
	return "api"
}
