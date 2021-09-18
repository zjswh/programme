package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zjswh/go-tool/utils"
	"programme/middleware"
	"programme/service/ProgrammeService"
	"programme/types"
	"programme/types/response"
)

func SaveConfig(c *gin.Context) {
	userInfo := middleware.GetBUserInfo(c)
	sType := utils.DefaultIntFormValue("type", 0, c)
	channelId := utils.DefaultIntFormValue("channelId", 0, c)
	week := utils.DefaultIntFormValue("week", -1, c)
	config := c.Request.FormValue("configs")

	if channelId == 0 || week == -1 {
		response.Result(2, "", "参数缺失", c)
		return
	}
	var programmeList []types.ProgrammeStruct
	json.Unmarshal([]byte(config), &programmeList)
	var programmeService ProgrammeService.TvService
	//列表为空 清除当天全部配置
	if programmeList == nil {
		err := programmeService.ClearProgramme(sType, userInfo.Uin, channelId, week)
		if err != nil {
			response.Result(0, "", "设置失败", c)
			return
		}
		response.Result(0, "设置成功", "", c)
		return
	}
	programmeService.SaveConfig(programmeList, week, sType, channelId, userInfo)
	response.Result(0, "success", "", c)
	return
}

func GetPlayInfo(c *gin.Context) {
	url := c.Query("mssUrl")
	if url == "" {
		response.Result(2, "", "参数缺失", c)
		return
	}

	mssInfo, err := ProgrammeService.GetMssInfo(url)
	if err != nil {
		response.Result(2, "", err.Error(), c)
		return
	}
	response.Result(0, mssInfo, "", c)
	return
}

func GetList(c *gin.Context) {
	sType := utils.DefaultIntFormValue("type", 0, c)
	channelId := utils.DefaultIntFormValue("channelId", 0, c)
	week := utils.DefaultIntFormValue("week", -1, c)

	if channelId == 0 || week == -1 {
		response.Result(2, "", "参数缺失", c)
		return
	}

	var programmeService ProgrammeService.TvService
	list, err := programmeService.GetProgramme(channelId, week, sType)
	if err != nil {
		response.Result(2, "", err.Error(), c)
		return
	}
	response.Result(0, list, "", c)
	return
}

func SyncConfig(c *gin.Context) {
	sType := utils.DefaultIntFormValue("type", 0, c)
	channelId := utils.DefaultIntFormValue("channelId", 0, c)
	week := utils.DefaultIntFormValue("week", -1, c)
	toWeek := c.Request.FormValue("toWeek")

	if channelId == 0 || week == -1 || toWeek == "" {
		response.Result(2, "", "参数缺失", c)
		return
	}
	//解析toweek
	var weekArr []int
	err := json.Unmarshal([]byte(toWeek), &weekArr)
	if err != nil {
		response.Result(2, "", "参数异常", c)
		return
	}

	//获取用户信息
	userInfo := middleware.GetBUserInfo(c)

	var programmeService ProgrammeService.TvService

	//获取当天的配置
	list, _ := programmeService.GetProgramme(channelId, week, sType)
	if len(list) == 0 {
		response.Result(2, "", "列表为空,同步失败", c)
		return
	}
	for _, w := range weekArr {
		var programmeList []types.ProgrammeStruct
		for _, v := range list {
			programmeStruct := types.ProgrammeStruct{
				Name:      v.Name,
				StartTime: v.StartTime,
				EndTime:   v.EndTime,
				ChannelId: channelId,
				Uin:       userInfo.Uin,
				Type:      sType,
				Week:      w,
			}
			programmeList = append(programmeList, programmeStruct)
		}
		programmeService.SaveConfig(programmeList, w, sType, channelId, userInfo)
	}
	response.Result(0, "success", "", c)
	return
}

func SaveBillType(c *gin.Context) {
	channelId := utils.DefaultIntFormValue("channelId", 0, c)
	week := utils.DefaultIntFormValue("week", -1, c)
	if channelId == 0 || week == -1 {
		response.Result(2, "", "参数缺失", c)
		return
	}
	var programmeService ProgrammeService.TvService
	list, err := programmeService.GetListByChannel(week, channelId)
	if err != nil {
		response.Result(2, "", err.Error(), c)
		return
	}
	response.Result(0, list, "", c)
	return
}

func GetListByChannel(c *gin.Context) {
	channelId := utils.DefaultIntFormValue("channelId", 0, c)
	week := utils.DefaultIntFormValue("week", -1, c)
	if channelId == 0 || week == -1 {
		response.Result(2, "", "参数缺失", c)
		return
	}
	var programmeService ProgrammeService.TvService
	list, err := programmeService.GetListByChannel(week, channelId)
	if err != nil {
		response.Result(2, "", err.Error(), c)
		return
	}
	response.Result(0, list, "", c)
	return
}


func RadioSaveConfig(c *gin.Context) {
	var saveConfigRequest types.Request
	err := c.ShouldBind(&saveConfigRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = ProgrammeService.RadioSaveConfig(saveConfigRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func RadioGetPlayInfo(c *gin.Context) {
	var getPlayInfoRequest types.Request
	err := c.ShouldBind(&getPlayInfoRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = ProgrammeService.RadioGetPlayInfo(getPlayInfoRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func RadioGetList(c *gin.Context) {
	var getListRequest types.Request
	err := c.ShouldBind(&getListRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = ProgrammeService.RadioGetList(getListRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func RadioRadioSyncConfig(c *gin.Context) {
	var syncConfigRequest types.Request
	err := c.ShouldBind(&syncConfigRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = ProgrammeService.RadioRadioSyncConfig(syncConfigRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func RadioSaveBillType(c *gin.Context) {
	var saveBillTypeRequest types.Request
	err := c.ShouldBind(&saveBillTypeRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = ProgrammeService.RadioSaveBillType(saveBillTypeRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func RadioGetListByChannel(c *gin.Context) {
	var getListByChannelRequest types.Request
	err := c.ShouldBind(&getListByChannelRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = ProgrammeService.RadioGetListByChannel(getListByChannelRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func GetBillPlayInfo(c *gin.Context) {
	channelId := utils.DefaultIntFormValue("channelId", 0, c)
	id := utils.DefaultIntFormValue("id", 0, c)
	if channelId == 0 || id == 0 {
		response.Result(2, "", "参数缺失", c)
		return
	}
	var programmeService ProgrammeService.TvService
	info, err := programmeService.GetProgrammeInfo(channelId, id)
	if err != nil {
		response.Result(2, "", err.Error(), c)
		return
	}
	status := 0
	if info.MTaskId > 0 {
		status = 1
	}
	programmeInfo := response.ProgrammeInfo{
		ID:     id,
		MssUrl: info.MssUrl,
		Status: status,
	}

	response.Result(0, programmeInfo, "", c)
	return
}

func ReplayOpRadio(c *gin.Context) {
	id := utils.DefaultIntParam("id", 0, c)
	isEnd := utils.DefaultIntParam("isEnd", 0, c)
	channelId := utils.DefaultIntParam("channelId", 0, c)
	if id == 0 || channelId == 0 {
		response.Result(2, "", "参数缺失", c)
		return
	}
	var programmeService ProgrammeService.TvService
	err := programmeService.ReplayOp(channelId, id, isEnd)
	if err != nil {
		response.Result(2, "", err.Error(), c)
		return
	}
	response.Result(0, "success", "", c)
	return
}

func SetRadioRebroadcast(c *gin.Context) {
	var setRadioRebroadcastRequest types.Request
	err := c.ShouldBind(&setRadioRebroadcastRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = ProgrammeService.SetRadioRebroadcast(setRadioRebroadcastRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func RadioGetBillPlayInfo(c *gin.Context) {
	var getBillPlayInfoRequest types.Request
	err := c.ShouldBind(&getBillPlayInfoRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = ProgrammeService.RadioGetBillPlayInfo(getBillPlayInfoRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func RadioReplayOpRadio(c *gin.Context) {
	var replayOpRadioRequest types.Request
	err := c.ShouldBind(&replayOpRadioRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = ProgrammeService.RadioReplayOpRadio(replayOpRadioRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func RadioSetRadioRebroadcast(c *gin.Context) {
	var setRadioRebroadcastRequest types.Request
	err := c.ShouldBind(&setRadioRebroadcastRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = ProgrammeService.RadioSetRadioRebroadcast(setRadioRebroadcastRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

