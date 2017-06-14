package main

import (
	"flag"
	//"git.iguiyu.com/park/api/configs"
	"github.com/curry/travel_api/configs"
	//"git.iguiyu.com/park/api/handle"
	"github.com/curry/travel_api/handle"
	"git.iguiyu.com/park/api/misc"
	"git.iguiyu.com/park/misc/helpers"
	"github.com/labstack/echo"
)

var (
	env        = flag.String("env", "dev", "Running Environment")
	configFile = flag.String("config", "", "The path of configuration file")
)

func init() {
	parseFlag()
	loadConfig()
	initOrm()
	initRedis()
	initRpc()
}

func main() {
	helpers.RunEchoServer(configs.AppConf, routes)
}

func routes(e *echo.Echo) {

	e.POST("/login", handle.Login)
	e.POST("/signup", handle.SignUp)
	e.GET("/sendcode/:phonenumber/:deviceid", handle.SendCodeToNoneExistNumber)
	e.GET("/sendcode2/:phonenumber/:deviceid", handle.SendCodeToExistNumber)
	e.PUT("/resetpassword", handle.ResetPassword)
	e.POST("/push", handle.Push, misc.KeyAuth())

	me := e.Group("/me", misc.KeyAuth())
	{
		me.GET("", handle.Me)
		me.GET("/stats/summary", handle.SummaryStats)
		me.GET("/stats/brief", handle.BriefStats)
		//me.GET("/account", handle.ShowAccount)
		//me.POST("/account", handle.UpdateAccount)
		me.POST("/bankaccount", handle.AddBankAccount)
		me.GET("/bankaccount", handle.ListBankAccount)
		me.DELETE("/bankaccount/:id", handle.DeleteBankAccount)
		me.GET("/banklist", handle.GetBankList)
		me.POST("/withdraw", handle.Withdraw)
		me.GET("/withdraw/:start/:limit", handle.ListWithdraw)
		me.GET("/sendcode/:deviceid", handle.SendWithdrawCode)
		me.GET("/device/:deviceid", handle.RegDevice)
	}

	plan := e.Group("/plans", misc.KeyAuth())
	{
		plan.GET("", handle.ListPlan)
		plan.POST("", handle.CreatePlan)
		plan.PUT("/:id", handle.UpdatePlan)
		plan.PUT("/:id/rename", handle.RenamePlan)
		plan.DELETE("/:id", handle.DeletePlan)
		plan.GET("/:id/locks", handle.ListLocksUnderPlan)
	}

	lock := e.Group("/locks", misc.KeyAuth())
	{
		lock.GET("", handle.ListLocks)
		lock.GET("/:id", handle.ShowLock)
		lock.PUT("/qrcode", handle.GetLock)
		lock.PUT("/:id/bind", handle.BindLock)
		lock.GET("/:id/:onoff", handle.TurnOnOff)
		lock.GET("/:id/plan/:planid", handle.ChangePlan)
		lock.GET("/:id/op/:downup", handle.DownOrUp)
	}

	msg := e.Group("/msg", misc.KeyAuth())
	{
		msg.GET("/list/:start/:limit", handle.ListMsg)
		msg.GET("/:id/read", handle.MarkRead)
		msg.DELETE("/:id", handle.MarkDelete)
		msg.GET("/readall", handle.MarkAllToRead)
		msg.DELETE("/delall", handle.MarkAllToDeleted)
		msg.GET("/unread", handle.CountUnRead)
	}

	tourist := e.Group("/tourist")
	{
		tourist.GET("", handle.TouristList)
		tourist.GET("/:start/:limit", handle.TouristPage)
		tourist.DELETE("/:id", handle.DeleteTourist)
	}

	collect := e.Group("/collect")
	{
		collect.PUT("/:id/:addDel", handle.CollectOrCancel)
		collect.GET("/list/:start/:limit", handle.CollectList)
		collect.GET("/count", handle.CountCollect)
	}

	learn := e.Group("/learn")
	{
		learn.GET("/learn", handle.Learn)
	}



}
