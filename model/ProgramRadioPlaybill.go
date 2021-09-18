package model

import (
	"gorm.io/gorm"
	"time"
)

type ProgramRadioPlaybill struct {
	Id         int    `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Uin        int    `gorm:"column:uin;type:int(11) unsigned;default:0;NOT NULL" json:"uin"`
	ChannelId  int    `gorm:"column:channelId;type:int(11) unsigned;default:0;NOT NULL" json:"channelId"`   // 频道id
	Name       string `gorm:"column:name;type:varchar(15);NOT NULL" json:"name"`                            // 节目名称
	Type       int    `gorm:"column:type;type:tinyint(1) unsigned;default:0;NOT NULL" json:"type"`          // 0自动1手动设置
	Week       int    `gorm:"column:week;type:tinyint(1);NOT NULL" json:"week"`                             // 周一-周日0-6
	MssUrl     string `gorm:"column:mssUrl;type:varchar(255)" json:"mssUrl"`                                // 媒资url
	StartTime  string `gorm:"column:startTime;type:char(8);default:00:00:00;NOT NULL" json:"startTime"`     // 开始时间
	EndTime    string `gorm:"column:endTime;type:char(8);default:00:00:00;NOT NULL" json:"endTime"`         // 结束时间
	STaskId    int    `gorm:"column:sTaskId;type:int(11) unsigned;default:0;NOT NULL" json:"sTaskId"`       // 开始计划任务ID
	ETaskId    int    `gorm:"column:eTaskId;type:int(11) unsigned;default:0;NOT NULL" json:"eTaskId"`       // 结束计划任务ID
	MTaskId    int    `gorm:"column:mTaskId;type:int(11) unsigned" json:"mTaskId"`                          // 合并id
	Deleted    int    `gorm:"column:deleted;type:tinyint(1) unsigned;default:0;NOT NULL" json:"deleted"`    // 1删除0正常
	DeleteTime int64  `gorm:"column:deleteTime;type:int(11) unsigned;default:0;NOT NULL" json:"deleteTime"` // 删除时间
	CreateTime int64  `gorm:"column:createTime;type:int(11);NOT NULL" json:"createTime"`                    // 创建时间
	UpdateTime int64  `gorm:"column:updateTime;type:int(11) unsigned;default:0;NOT NULL" json:"updateTime"` // 更新时间
}

func (m *ProgramRadioPlaybill) TableName() string {
	return "program_radio_playbill"
}

func (m *ProgramRadioPlaybill) Create(Db *gorm.DB) error {
	err := Db.Model(&m).Create(&m).Error
	return err
}

func (m *ProgramRadioPlaybill) Update(Db *gorm.DB, field ...string) error {
	sql := Db.Model(&m)
	if len(field) > 0 {
		sql = sql.Select(field)
	}
	err := sql.Where("id", m.Id).Updates(m).Error
	return err
}

func (m *ProgramRadioPlaybill) GetInfo(Db *gorm.DB) error {
	sql := Db.Model(m).Where("id = ? ", m.Id)
	err := sql.First(&m).Error
	return err
}

func (m *ProgramRadioPlaybill) GetList(Db *gorm.DB) ([]ProgramRadioPlaybill, error) {
	var list []ProgramRadioPlaybill
	err := Db.Model(m).Where("channelId = ? AND week = ? AND deleted = 0 AND type = ?", m.ChannelId, m.Week, m.Type).Find(&list).Error
	return list, err
}

func GetRadioProgrammeInfo(Db *gorm.DB, channelId int, id int) (ProgramRadioPlaybill, error) {
	var info ProgramRadioPlaybill
	err := Db.Model(&ProgramRadioPlaybill{}).Where("channelId = ? AND id = ? AND deleted = 0", channelId, id).Find(&info).Error
	return info, err
}

func GetRadioProgrammeInfoById(Db *gorm.DB, id int) (ProgramRadioPlaybill, error) {
	var info ProgramRadioPlaybill
	err := Db.Model(&ProgramRadioPlaybill{}).Where("id = ? AND deleted = 0", id).Find(&info).Error
	return info, err
}

func (m *ProgramRadioPlaybill) Clear(Db *gorm.DB) error {
	err := Db.Model(m).
		Where("type = ? AND uin = ? AND channelId = ? AND week = ? ", m.Type, m.Uin, m.ChannelId, m.Week).
		Updates(ProgramRadioPlaybill{Deleted: 1, DeleteTime: time.Now().Unix()}).Error
	return err
}

func DeleteRadioPlayBill(Db *gorm.DB, idArr []int) error {
	err := Db.Model(ProgramRadioPlaybill{}).Where("id IN (?)", idArr).
		Updates(ProgramRadioPlaybill{Deleted: 1, DeleteTime: time.Now().Unix()}).Error
	return err
}
