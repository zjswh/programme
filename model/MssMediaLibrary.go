package model

import "gorm.io/gorm"

// 媒资库
type MssMediaLibrary struct {
	MssId     uint   `gorm:"column:mssId;type:int(10) unsigned;AUTO_INCREMENT;NOT NULL" json:"mssId"`
	Title     string `gorm:"column:title;type:varchar(32);NOT NULL" json:"title"`
	Url       string `gorm:"column:url;type:varchar(2000);NOT NULL" json:"url"`                        // 播放地址
	Duration  uint   `gorm:"column:duration;type:int(10) unsigned;default:0;NOT NULL" json:"duration"` // 视频的时长
	Thumbnail string `gorm:"column:thumbnail;type:varchar(2000);NOT NULL" json:"thumbnail"`            // 视频的封面图
}

func (m *MssMediaLibrary) TableName() string {
	return "mss_media_library"
}

func (m *MssMediaLibrary) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *MssMediaLibrary) GetInfo(Db *gorm.DB) error  {
	err := Db.Model(m).Where("url", m.Url).Find(&m).Error
	return err
}
