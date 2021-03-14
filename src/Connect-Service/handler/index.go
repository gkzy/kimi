package handler

import "github.com/gkzy/gow"

func IndexHandler(c *gow.Context){

	userId,_:=c.GetInt64("user_id",0)
	deviceId,_:=c.GetInt64("device_id",0)
	token:=c.GetString("token")

	c.Data["userId"] = userId
	c.Data["deviceId"] = deviceId
	c.Data["token"] = token
	c.HTML("index.html")
}