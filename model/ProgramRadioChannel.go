package model

import "gorm.io/gorm"

// 电台频道
type ProgramRadioChannel struct {
	Id                   int   `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name                 string `gorm:"column:name;type:varchar(180);NOT NULL" json:"name"`                            // 名称
	Introduce            string `gorm:"column:introduce;type:varchar(50);NOT NULL" json:"introduce"`                   // 简介
	Logo                 string `gorm:"column:logo;type:varchar(50);NOT NULL" json:"logo"`                             // 聚合页logo
	FlowType             int   `gorm:"column:flowType;type:tinyint(1) unsigned;default:1;NOT NULL" json:"flowType"`   // 推流类型1视频2音频
	AudioCover           string `gorm:"column:audioCover;type:varchar(50);NOT NULL" json:"audioCover"`                 // 纯音频封底图
	MpsLssAppId          string `gorm:"column:mpsLssAppId;type:varchar(20);NOT NULL" json:"mpsLssAppId"`               // 播放器实例id
	MpsDynamicVodAppId   string `gorm:"column:mpsDynamicVodAppId;type:varchar(20);NOT NULL" json:"mpsDynamicVodAppId"` // mps 动态点播实例id
	MpsBackground        string `gorm:"column:mpsBackground;type:varchar(64);NOT NULL" json:"mpsBackground"`
	Status               int   `gorm:"column:status;type:tinyint(1) unsigned;default:1;NOT NULL" json:"status"`                   // 状态1正常2回收站
	Uin                  int   `gorm:"column:uin;type:int(10) unsigned;default:0;NOT NULL" json:"uin"`                            // 操作人
	ChildId              int   `gorm:"column:childId;type:int(11) unsigned;default:0;NOT NULL" json:"childId"`                    // 子账号id-0主账号
	Aid                  int   `gorm:"column:aid;type:int(10) unsigned;default:0;NOT NULL" json:"aid"`                            // 频道id
	PlaylistsStatus      int   `gorm:"column:playlistsStatus;type:tinyint(1) unsigned;default:2;NOT NULL" json:"playlistsStatus"` // µã²¥Ä£°å×´Ì¬1ÆôÓÃ2¹Ø±Õ
	PlaylistsId          int   `gorm:"column:playlistsId;type:int(10) unsigned;default:0;NOT NULL" json:"playlistsId"`            // µã²¥Ä£°åid
	CreateTime           int   `gorm:"column:createTime;type:int(10) unsigned;default:0;NOT NULL" json:"createTime"`              // 创建时间
	UpdateTime           int   `gorm:"column:updateTime;type:int(10) unsigned;default:0;NOT NULL" json:"updateTime"`              // 修改时间
	PublishTaskId        int   `gorm:"column:publishTaskId;type:int(10) unsigned;default:0;NOT NULL" json:"publishTaskId"`        // ·¢²¼ÈÎÎñid
	PublishTaskMakeAppId string `gorm:"column:publishTaskMakeAppId;type:varchar(20);NOT NULL" json:"publishTaskMakeAppId"`         // ·¢²¼ÈÎÎñ°ó¶¨µÄÖÆ×÷µ¼²¥Ì¨ÊµÀýid
	PullUrl              string `gorm:"column:pullUrl;type:varchar(1000);NOT NULL" json:"pullUrl"`                                 // À­Á÷µØÖ·
	PullStatus           int    `gorm:"column:pullStatus;type:tinyint(1);default:0;NOT NULL" json:"pullStatus"`                    // À­Á÷µØÖ·ÊÇ·ñÆôÓÃ£¨0¹Ø±Õ£¬1ÆôÓÃ£©
	ChatAuthModel        int   `gorm:"column:chatAuthModel;type:tinyint(1) unsigned;default:2;NOT NULL" json:"chatAuthModel"`     // ÁÄÌìÉóºËÀàÐÍ1£ºÖðÌõÉóºË2Ãô¸ÐÉóºË3×Ô¶¯¹ýÂË
	ChatSee              int    `gorm:"column:chatSee;type:tinyint(1);default:1;NOT NULL" json:"chatSee"`
	ChatLogin            int   `gorm:"column:chatLogin;type:tinyint(1) unsigned;default:0;NOT NULL" json:"chatLogin"` // ÁÄÌìÊÇ·ñÐèÒªµÇÂ¼£¬1ÐèÒª0²»ÐèÒª
	ChatAdminName        string `gorm:"column:chatAdminName;type:varchar(20);NOT NULL" json:"chatAdminName"`           // ÁÄÌì¹ÜÀíÔ±Ãû
	ChatAdminAva         string `gorm:"column:chatAdminAva;type:varchar(64);NOT NULL" json:"chatAdminAva"`             // ÁÄÌì¹ÜÀíÔ±Í·Ïñ
	DefaultUserAva       string `gorm:"column:defaultUserAva;type:varchar(64);NOT NULL" json:"defaultUserAva"`
	PullStream           int   `gorm:"column:pullStream;type:int(1) unsigned;default:0;NOT NULL" json:"pullStream"` // ÊÇ·ñ½øÐÐÀ­Á÷ 1ÊÇ 0·ñ
	LssApp               string `gorm:"column:lssApp;type:varchar(64);NOT NULL" json:"lssApp"`                       // lss APP
	Stream               string `gorm:"column:stream;type:varchar(64);NOT NULL" json:"stream"`                       // 流stream
	StreamType           int   `gorm:"column:streamType;type:tinyint(1) unsigned;default:1;NOT NULL" json:"streamType"`
	WatchNum             int    `gorm:"column:watchNum;type:int(10);default:0;NOT NULL" json:"watchNum"`                      // 观看数量
	IsDisplayWatchNum    int    `gorm:"column:isDisplayWatchNum;type:tinyint(1);default:1;NOT NULL" json:"isDisplayWatchNum"` // 是否展示观看量1是2否
	CustomWatchStatus    int    `gorm:"column:customWatchStatus;type:tinyint(1);default:2;NOT NULL" json:"customWatchStatus"` // 是否开启自定义观看数据，2否1是
	ShareImg             string `gorm:"column:shareImg;type:varchar(64)" json:"shareImg"`
	ShareSubTitle        string `gorm:"column:shareSubTitle;type:varchar(200);NOT NULL" json:"shareSubTitle"`
	ShareTitle           string `gorm:"column:shareTitle;type:varchar(200);NOT NULL" json:"shareTitle"`
	LayoutColor          string `gorm:"column:layoutColor;type:varchar(10);NOT NULL" json:"layoutColor"`
	PlayStatus           int    `gorm:"column:playStatus;type:tinyint(1);default:1;NOT NULL" json:"playStatus"` // 是否开启直播
	LssPublishUrl        string `gorm:"column:lssPublishUrl;type:varchar(255)" json:"lssPublishUrl"`
	PlaybillType         int    `gorm:"column:playbillType;type:tinyint(1);default:0;NOT NULL" json:"playbillType"` // 节目单 0为默认1为自定义
}

func (m *ProgramRadioChannel) TableName() string {
	return "program_radio_channel"
}

func (m *ProgramRadioChannel) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *ProgramRadioChannel) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *ProgramRadioChannel) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}

func GetRadioInfo(Db *gorm.DB, id int) (ProgramRadioChannel, error) {
	var info ProgramRadioChannel
	err := Db.Model(&ProgramRadioChannel{}).Where("id = ? AND status = 1", id).Find(&info).Error
	return info, err
}
