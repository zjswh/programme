
package router

import (
	"github.com/gin-gonic/gin"
	v1 "programme/api/v1"
	"programme/middleware"
)

func InitRouter(Router *gin.RouterGroup) {
	ProgrammeCheckLoginRouter := Router.Group("").
		Use(middleware.CheckLogin())
	{
		ProgrammeCheckLoginRouter.POST("v1/programme/Index/saveConfig", v1.SaveConfig)
		ProgrammeCheckLoginRouter.GET("v1/programme/Index/getPlayInfo", v1.GetPlayInfo)
		ProgrammeCheckLoginRouter.GET("v1/programme/Index/getList", v1.GetList)
		ProgrammeCheckLoginRouter.POST("v1/programme/Index/syncConfig", v1.SyncConfig)
		ProgrammeCheckLoginRouter.POST("v1/programme/Index/saveBillType", v1.SaveBillType)
		ProgrammeCheckLoginRouter.GET("v1/programme/Index/getListByChannel", v1.GetListByChannel)
		ProgrammeCheckLoginRouter.POST("v1/programme/Radio/saveConfig", v1.RadioSaveConfig)
		ProgrammeCheckLoginRouter.GET("v1/programme/Radio/getPlayInfo", v1.RadioGetPlayInfo)
		ProgrammeCheckLoginRouter.GET("v1/programme/Radio/getList", v1.RadioGetList)
		ProgrammeCheckLoginRouter.POST("v1/programme/Radio/syncConfig", v1.RadioRadioSyncConfig)
		ProgrammeCheckLoginRouter.POST("v1/programme/Radio/saveBillType", v1.RadioSaveBillType)
		ProgrammeCheckLoginRouter.GET("v1/programme/Radio/getListByChannel", v1.RadioGetListByChannel)
	}

	ProgrammeRouter := Router.Group("")
	{
		ProgrammeRouter.GET("v1/programme/Index/getBillPlayInfo", v1.GetBillPlayInfo)
		ProgrammeRouter.POST("v1/programme/Index/replayOpRadio", v1.ReplayOpRadio)
		ProgrammeRouter.POST("v1/programme/Index/setRadioRebroadcast", v1.SetRadioRebroadcast)
		ProgrammeRouter.GET("v1/programme/Radio/getBillPlayInfo", v1.RadioGetBillPlayInfo)
		ProgrammeRouter.POST("v1/programme/Radio/replayOpRadio", v1.RadioReplayOpRadio)
		ProgrammeRouter.POST("v1/programme/Radio/setRadioRebroadcast", v1.RadioSetRadioRebroadcast)
	}


}
