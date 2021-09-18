package ProgrammeService

import (
	"encoding/json"
	"fmt"
	utils "github.com/zjswh/go-tool/cron"
	"github.com/zjswh/go-tool/timex"
	"programme/config"
	"programme/lib/aodianyun"
	"programme/lib/tools"
	"programme/model"
	"programme/types"
	"time"
)

type TvService struct {}

func(t *TvService) GetProgramme(channelId int, week int, sType int) ([]model.ProgramTvPlaybill, error) {
	programme := model.ProgramTvPlaybill{
		ChannelId: channelId,
		Week:      week,
		Type:      sType,
	}
	list, err := programme.GetList(config.GVA_DB)
	return list, err
}

func(t *TvService) CreateCron(programmeStruct types.ProgrammeStruct) {
	var url string
	url = config.GVA_CONFIG.Param.ProjectHost + "/v1/programme/Index/replayOpRadio"
	url = fmt.Sprintf("%s?id=%d&channelId=%d", url, programmeStruct.Id, programmeStruct.ChannelId)
	//$name =  $channelInfo['uin'] . '-'.$taskName.'-' . $id;
	name := fmt.Sprintf("%d-Live-Scron-%d", programmeStruct.Uin, programmeStruct.Id)
	//创建开始定时任务
	res, err := utils.CreateCron(name, tools.GetCronTime(programmeStruct.StartTime, programmeStruct.Week), url, programmeStruct.STaskId)
	if err != nil || res.Code != 0 {
		return
	}
	sTaskId := res.Data
	name = fmt.Sprintf("%d-Live-Ecron-%d", programmeStruct.Uin, programmeStruct.Id)
	url = fmt.Sprintf("%s&isEnd=%d", url, 1)
	//创建结束定时任务
	res, err = utils.CreateCron(name, tools.GetCronTime(programmeStruct.EndTime, programmeStruct.Week), url, programmeStruct.ETaskId)
	if err != nil || res.Code != 0 {
		return
	}
	eTaskId := res.Data

	//更新数据
	programme := model.ProgramTvPlaybill{
		Id:      programmeStruct.Id,
		ETaskId: eTaskId,
		STaskId: sTaskId,
	}
	err = programme.Update(config.GVA_DB)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func(t *TvService) AddProgramme(programmeStruct types.ProgrammeStruct) (int, error) {
	programme := model.ProgramTvPlaybill{
		Uin:        programmeStruct.Uin,
		ChannelId:  programmeStruct.ChannelId,
		Name:       programmeStruct.Name,
		Type:       programmeStruct.Type,
		Week:       programmeStruct.Week,
		MssUrl:     programmeStruct.MssUrl,
		StartTime:  programmeStruct.StartTime,
		EndTime:    programmeStruct.EndTime,
		CreateTime: time.Now().Unix(),
	}
	err := programme.Create(config.GVA_DB)
	return programme.Id, err
}

//设置为默认节目单 调用底层接口
func(t *TvService) setLiveReview(channelId int, playbillType int) {
	info, _ := model.GetTvInfo(config.GVA_DB, channelId)
	fmt.Println(info)
	if playbillType == 1 {
		aodianyun.RemoveLiveReview(info.Uin, info.LssApp, info.Stream)
	} else {
		aodianyun.EnableLiveReview(info.Uin, info.LssApp, info.Stream)
	}
}

func(t *TvService) UpdateProgramme(programmeStruct types.ProgrammeStruct) error {
	programme := model.ProgramTvPlaybill{
		Id:         programmeStruct.Id,
		Uin:        programmeStruct.Uin,
		ChannelId:  programmeStruct.ChannelId,
		Name:       programmeStruct.Name,
		Type:       programmeStruct.Type,
		Week:       programmeStruct.Week,
		MssUrl:     programmeStruct.MssUrl,
		StartTime:  programmeStruct.StartTime,
		EndTime:    programmeStruct.EndTime,
		UpdateTime: time.Now().Unix(),
	}
	err := programme.Update(config.GVA_DB)
	return err
}

func DeleteCron(sTaskId int, eTaskId int) {
	utils.DeleteCron(sTaskId)
	utils.DeleteCron(eTaskId)
}

func(t *TvService)  ClearProgramme(stype int, uin int, channelId int, week int) error {
	programme := model.ProgramTvPlaybill{
		Uin:       uin,
		ChannelId: channelId,
		Type:      stype,
		Week:      week,
	}
	err := programme.Clear(config.GVA_DB)
	return err
}

func GetMssInfo(url string)  (model.MssMediaLibrary, error) {
	mss := model.MssMediaLibrary{
		Url: url,
	}
	err := mss.GetInfo(config.GVA_DB)
	return mss , err
}

func (t *TvService) SaveConfig(programmeList []types.ProgrammeStruct, week int, sType int, channelId int, userInfo config.UserInfo){
	//获取旧配置
	list, _ := t.GetProgramme(channelId, week, sType)
	mapList := map[int]model.ProgramTvPlaybill{}
	if len(list) > 0 {
		for _, v := range list {
			mapList[v.Id] = v
		}
	}

	for _, v := range programmeList {
		v.Week = week
		v.Type = sType
		v.ChannelId = channelId
		v.Uin = userInfo.Uin
		if v.Id == 0 {
			id, err := t.AddProgramme(v)
			if err == nil {
				v.Id = id
				//添加定时任务
				go t.CreateCron(v)
			}
		} else {
			err := t.UpdateProgramme(v)
			if err == nil {
				//添加定时任务
				if _, ok := mapList[v.Id]; ok {
					v.STaskId = mapList[v.Id].STaskId
					v.ETaskId = mapList[v.Id].ETaskId
					//删除key
					delete(mapList, v.Id)
				}
				go t.CreateCron(v)
			}
		}
	}

	var deleteIdArr []int
	//删除不要的旧数据
	if len(mapList) > 0 {
		for k, v := range mapList {
			deleteIdArr = append(deleteIdArr, k)
			//删除定时任务
			go DeleteCron(v.STaskId, v.ETaskId)
		}
		model.DeleteTvPlayBill(config.GVA_DB, deleteIdArr)
	}

	t.setLiveReview(channelId, sType)
	key := fmt.Sprintf("tv_bill_config_new_%d_%d", channelId, week)

	//清除缓存
	config.GVA_REDIS.Del(key).Result()
	return
}

func(t *TvService) UpdateBillType(channelId int, playbillType int) error {
	channel := model.ProgramTvChannel{
		Id:           channelId,
		PlaybillType: playbillType,
	}

	//设置为默认节目单 调用底层接口
	t.setLiveReview(channelId, playbillType)

	err := channel.Update(config.GVA_DB)
	return err
}

func(t *TvService) GetListByChannel(week int, channelId int) ([]model.ProgramTvPlaybill, error) {
	key := fmt.Sprintf("tv_bill_config_new_%d_%d", channelId, week)
	info, _ := config.GVA_REDIS.Get(key).Result()

	if info != "" {
		var list []model.ProgramTvPlaybill
		json.Unmarshal([]byte(info), &list)
		return list, nil
	}

	channelInfo, _ := model.GetTvInfo(config.GVA_DB, channelId)
	//获取对应的数据
	list, err := t.GetProgramme(channelId, week, channelInfo.PlaybillType)
	if err == nil {
		keyInfo, _ := json.Marshal(list)
		err := config.GVA_REDIS.Set(key, keyInfo, time.Hour).Err()
		if err != nil {
			panic(err)
		}
	}
	return list, err
}

func(t *TvService) GetProgrammeInfo(channelId int, id int) (model.ProgramTvPlaybill, error) {
	programme, err := model.GetTvProgrammeInfo(config.GVA_DB, channelId, id)
	return programme, err
}

func(t *TvService) ReplayOp(channelId int, id int, isEnd int) error {
	channelInfo, _ := model.GetTvInfo(config.GVA_DB, channelId)
	programmeInfo, _ := model.GetTvProgrammeInfo(config.GVA_DB, channelId, id)
	if channelInfo.Id == 0 {
		// 直播间不存在 删除对应节目单
		DeleteCron(programmeInfo.STaskId, programmeInfo.ETaskId)
		return fmt.Errorf("直播间不存在")
	}

	uin := channelInfo.Uin
	result, err := aodianyun.ReplayOp(uin, channelInfo.LssApp, channelInfo.Stream, "video")
	if err != nil {
		return err
	}
	if result.Flag != 100 {
		return fmt.Errorf(result.FlagString)
	}

	if isEnd == 1 {
		nowTime := time.Now()
		tim := nowTime.Format("2006-01-02")
		// 底层异步需要加30秒延迟
		startTime := timex.DateToTimeStamp(tim+" "+programmeInfo.StartTime, "") + 60
		endTime := timex.DateToTimeStamp(tim+" "+programmeInfo.EndTime, "") + 60
		url := config.GVA_CONFIG.Param.ProjectHost + "/v1/programme/Index/setRadioRebroadcast"
		query := fmt.Sprintf("startTime=%d&endTime=%d&uin=%d&app=%s&stream=%s&title=%s&id=%d",
			startTime, endTime, uin, channelInfo.LssApp, channelInfo.Stream, programmeInfo.Name, id)
		url = url + "?" + query
		name := fmt.Sprintf("%d-tvMerge-cron-%d", uin, id)
		date := "@every 5m1s"
		result, _ := utils.CreateCron(name, date, url, 0)
		if result.Code != 0 {
			errMsg := fmt.Sprintf("合并任务创建失败+%s", err.Error())
			return fmt.Errorf(errMsg)
		}
		//更新节目单 将合并id保存
		programme := model.ProgramTvPlaybill{
			Id:      id,
			MTaskId: result.Data,
		}
		programme.Update(config.GVA_DB)
		url = fmt.Sprintf( "%s&taskId=%d", url, result.Data)
		//更新定时器
		result, _ = utils.CreateCron(name, date, url, result.Data)
		if result.Code != 0 {
			errMsg := fmt.Sprintf("修改合并任务失败+%s", err.Error())
			return fmt.Errorf(errMsg)
		}
	}

	return nil
}

func RadioSaveConfig(req types.Request) error {
	//add your code ...
	return nil
}

func RadioGetPlayInfo(req types.Request) error {
	//add your code ...
	return nil
}

func RadioGetList(req types.Request) error {
	//add your code ...
	return nil
}

func RadioRadioSyncConfig(req types.Request) error {
	//add your code ...
	return nil
}

func RadioSaveBillType(req types.Request) error {
	//add your code ...
	return nil
}

func RadioGetListByChannel(req types.Request) error {
	//add your code ...
	return nil
}

func SetRadioRebroadcast(req types.Request) error {
	//add your code ...
	return nil
}

func RadioGetBillPlayInfo(req types.Request) error {
	//add your code ...
	return nil
}

func RadioReplayOpRadio(req types.Request) error {
	//add your code ...
	return nil
}

func RadioSetRadioRebroadcast(req types.Request) error {
	//add your code ...
	return nil
}

