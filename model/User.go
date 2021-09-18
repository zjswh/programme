package model

import "gorm.io/gorm"

type User struct {
	Uin                      int   `gorm:"column:uin;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"uin"`
	Password                 string `gorm:"column:password;type:char(32);NOT NULL" json:"password"`                                    // 密码
	RealName                 string `gorm:"column:realName;type:varchar(50);NOT NULL" json:"realName"`                                 // 真实姓名
	Company                  string `gorm:"column:company;type:varchar(50);NOT NULL" json:"company"`                                   // µ¥Î»
	SpecialRoomType          int   `gorm:"column:specialRoomType;type:tinyint(1) unsigned;default:0;NOT NULL" json:"specialRoomType"` // 1特殊域名对接2一般特殊对接
	Ava                      string `gorm:"column:ava;type:varchar(200);NOT NULL" json:"ava"`                                          // 头像
	AodianUin                int   `gorm:"column:aodianUin;type:int(10) unsigned;default:0;NOT NULL" json:"aodianUin"`                // 奥点云用户id
	AodianAccessId           string `gorm:"column:aodianAccessId;type:varchar(64);NOT NULL" json:"aodianAccessId"`                     // 奥点云access id
	AodianAccessKey          string `gorm:"column:aodianAccessKey;type:varchar(64);NOT NULL" json:"aodianAccessKey"`                   // 奥点云access key
	AodianUinPStream         int   `gorm:"column:aodianUinPStream;type:int(10) unsigned;default:0;NOT NULL" json:"aodianUinPStream"`  // °ÂµãÔÆÓÃ»§À­Á÷id
	AodianAccessIdPStream    string `gorm:"column:aodianAccessIdPStream;type:varchar(64);NOT NULL" json:"aodianAccessIdPStream"`       // °ÂµãÔÆÀ­Á÷access id
	AodianAccessKeyPStream   string `gorm:"column:aodianAccessKeyPStream;type:varchar(64);NOT NULL" json:"aodianAccessKeyPStream"`     // °ÂµãÔÆÀ­Á÷access key
	LssApp                   string `gorm:"column:lssApp;type:varchar(64);NOT NULL" json:"lssApp"`
	LssAppPStream            string `gorm:"column:lssAppPStream;type:varchar(64);NOT NULL" json:"lssAppPStream"`
	LpsApp                   string `gorm:"column:lpsApp;type:varchar(64);NOT NULL" json:"lpsApp"`
	BroadcastMakeHsmsGroupId string `gorm:"column:broadcastMakeHsmsGroupId;type:varchar(64);NOT NULL" json:"broadcastMakeHsmsGroupId"`     // 制作场景组ID
	AuthStatus               int   `gorm:"column:authStatus;type:tinyint(1) unsigned;default:0;NOT NULL" json:"authStatus"`               // 实名认证状态，0未认证，1认证
	RegisterTime             int   `gorm:"column:registerTime;type:int(10) unsigned;default:0;NOT NULL" json:"registerTime"`              // 注册时间
	LastLoginTime            int   `gorm:"column:lastLoginTime;type:int(10) unsigned;default:0;NOT NULL" json:"lastLoginTime"`            // 最后登录时间
	LcpsUsedTime             int   `gorm:"column:lcpsUsedTime;type:int(10) unsigned;default:0;NOT NULL" json:"lcpsUsedTime"`              // 累积使用时间
	AccountStatus            int   `gorm:"column:accountStatus;type:tinyint(1) unsigned;default:1;NOT NULL" json:"accountStatus"`         // 账户状态，1正常
	DmsKeyInfo               string `gorm:"column:dmsKeyInfo;type:varchar(500);NOT NULL" json:"dmsKeyInfo"`                                // dms key json信息
	MsgDmsKey                string `gorm:"column:msgDmsKey;type:varchar(500);NOT NULL" json:"msgDmsKey"`                                  // 消息dms
	UserType                 int   `gorm:"column:userType;type:tinyint(3) unsigned;default:0;NOT NULL" json:"userType"`                   // 用户分类
	UserFrom                 int   `gorm:"column:userFrom;type:tinyint(1) unsigned;default:1;NOT NULL" json:"userFrom"`                   // 用户来源1商务客户2代理商客户
	BroadcastMakeType        int   `gorm:"column:broadcastMakeType;type:tinyint(1) unsigned;default:2;NOT NULL" json:"broadcastMakeType"` // 制作导播台信源数类型，1：5个，2：4个
	LiveProgramMpsId         int   `gorm:"column:liveProgramMpsId;type:int(10) unsigned;default:0;NOT NULL" json:"liveProgramMpsId"`      // 活动直播播放器ID
	TvProgramMpsId           int   `gorm:"column:tvProgramMpsId;type:int(10) unsigned;default:0;NOT NULL" json:"tvProgramMpsId"`          // 电视直播播放器id
	RadioProgramMpsId        int   `gorm:"column:radioProgramMpsId;type:int(10) unsigned;default:0;NOT NULL" json:"radioProgramMpsId"`    // 电台直播播放器
	BusinessManagerId        int   `gorm:"column:businessManagerId;type:int(10) unsigned;default:0;NOT NULL" json:"businessManagerId"`    // 对接商务经理id
	Status                   int   `gorm:"column:status;type:tinyint(1) unsigned;default:1;NOT NULL" json:"status"`                       // ×´Ì¬1ÆôÓÃ2½ûÓÃ
	Maxrecordduration        int   `gorm:"column:maxrecordduration;type:int(2) unsigned;default:1;NOT NULL" json:"maxrecordduration"`     // Ã½×Ê×î´ó´æ´¢Ê±¼ä
	NewsTitle                string `gorm:"column:newsTitle;type:varchar(255);NOT NULL" json:"newsTitle"`                                  // ÐÂÎÅÒ³Ãæ×Ô¶¨Òåtitle
	RbacStatus               int   `gorm:"column:rbacStatus;type:tinyint(1) unsigned;default:0;NOT NULL" json:"rbacStatus"`
	IsExamine                int    `gorm:"column:isExamine;type:tinyint(1);default:0;NOT NULL" json:"isExamine"`
	LeagueStatus             int    `gorm:"column:leagueStatus;type:tinyint(1);default:0;NOT NULL" json:"leagueStatus"`
	Cname                    string `gorm:"column:cname;type:varchar(255);NOT NULL" json:"cname"`
	CnameStatus              int   `gorm:"column:cnameStatus;type:tinyint(1) unsigned;default:0;NOT NULL" json:"cnameStatus"`   // cname https 0¹Ø±Õ1ÆôÓÃ
	AgencyLoginBackground    string `gorm:"column:agencyLoginBackground;type:varchar(50);NOT NULL" json:"agencyLoginBackground"` // 登录页背景图
	TvLssApp                 string `gorm:"column:tvLssApp;type:varchar(64);NOT NULL" json:"tvLssApp"`                           // 电视电台lss app
	CopyLssApp               string `gorm:"column:copyLssApp;type:varchar(64);NOT NULL" json:"copyLssApp"`                       // 直播间复制拉流lss app
	CmsExamine               int   `gorm:"column:cmsExamine;type:tinyint(1) unsigned;default:0;NOT NULL" json:"cmsExamine"`     // cms直播审核，开启1，关闭0
	TvStorage                int   `gorm:"column:tvStorage;type:tinyint(1) unsigned;default:1;NOT NULL" json:"tvStorage"`       // 电视直播是否存储1是2否
	RadioStorage             int   `gorm:"column:radioStorage;type:tinyint(1) unsigned;default:1;NOT NULL" json:"radioStorage"` // 电台直播是否存储1是2否
	ProjectVersion           int   `gorm:"column:projectVersion;type:tinyint(1) unsigned;default:1;NOT NULL" json:"projectVersion"`
	LiveVersion              int   `gorm:"column:liveVersion;type:tinyint(1) unsigned;default:1;NOT NULL" json:"liveVersion"`
	AuthKey                  string `gorm:"column:authKey;type:varchar(30);NOT NULL" json:"authKey"`                                       // 认证口令
	RoleInit                 int    `gorm:"column:roleInit;type:tinyint(1);default:0;NOT NULL" json:"roleInit"`                            // 角色列表初始化标志
	TwAuditing               int    `gorm:"column:twAuditing;type:tinyint(1);default:1;NOT NULL" json:"twAuditing"`                        // 图文直播审核开关  1开0关
	ClientStore              int    `gorm:"column:clientStore;type:tinyint(1);default:1;NOT NULL" json:"clientStore"`                      // 是否开启连麦导播存储
	Version                  string `gorm:"column:version;type:varchar(10);default:produce;NOT NULL" json:"version"`                       // produce为正式pre为灰度
	Platform                 string `gorm:"column:platform;type:varchar(30);default:gdy;NOT NULL" json:"platform"`                         // 平台
	UserId                   string `gorm:"column:userId;type:varchar(40);default:0" json:"userId"`                                        // 平台唯一Id
	DomainConfig             string `gorm:"column:domainConfig;type:varchar(255);NOT NULL" json:"domainConfig"`                            // 域名防盗链配置
	DemandReferStatus        int   `gorm:"column:demandReferStatus;type:tinyint(3) unsigned;default:2;NOT NULL" json:"demandReferStatus"` // 点播防盗链状态，1开启，2关闭
}

func (m *User) TableName() string {
	return "user"
}

func (m *User) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func GetUserInfo(Db *gorm.DB, uin int) (User, error) {
	var info User
	err := Db.Model(User{}).Where("uin = ? AND status = 1", uin).Find(&info).Error
	return info, err
}
