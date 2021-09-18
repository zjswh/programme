package aodianyun

import (
	"encoding/json"
	"fmt"
	"github.com/zjswh/go-tool/utils"
	"programme/config"
	"programme/model"
	"time"
)

const (
	BASE_OPENAPI_URL = "https://openapi.aodianyun.com/v2/"
)

type ApiResult struct {
	Flag       int    `json:"flag"`
	FlagString string `json:"flagString"`
}

type AodianAccess struct {
	AccessId string `json:"access_id"`
	AccessKey string `json:"access_key"`
}

func OpenApi(url string, uin int, param map[string]interface{}) ([]byte, error) {
	url = BASE_OPENAPI_URL + url
	fmt.Println(uin)
	userInfo := GetAodianAccess(uin)
	param["access_id"] = userInfo.AccessId
	param["access_key"] = userInfo.AccessKey
	res, err := utils.Request(url, param, map[string]interface{}{}, "POST", "json")
	if err != nil {
		panic(err)
	}
	paramEncode, _ := json.Marshal(param)
	log := fmt.Sprintf("params:%s, rst:%s", paramEncode, string(res))
	fmt.Println(log)
	return res, nil
}

func GetAodianAccess(uin int) AodianAccess {
	userAodianAccessCacheKey := fmt.Sprintf("user_aodian_access:%d", uin)
	info, _ := config.GVA_REDIS.Get(userAodianAccessCacheKey).Result()
	var aodianAccess AodianAccess
	if info != "" {
		json.Unmarshal([]byte(info), &aodianAccess)
		return aodianAccess
	}
	userInfo, _ := model.GetUserInfo(config.GVA_DB, uin)
	aodianAccess.AccessId = userInfo.AodianAccessId
	aodianAccess.AccessKey = userInfo.AodianAccessKey
	cacheInfo, _ := json.Marshal(aodianAccess)
	config.GVA_REDIS.Set(userAodianAccessCacheKey, cacheInfo, time.Hour)
	return aodianAccess
}
