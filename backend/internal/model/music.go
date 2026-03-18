package model

import "time"

type Music struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255);index;not null" json:"name"`
	Url       string    `gorm:"type:varchar(255)" json:"url"`
	PicUrl    string    `gorm:"type:varchar(500)" json:"pic_url"`
	Artists   string    `gorm:"type:varchar(255)" json:"artists"`
	Duration  int       `gorm:"type:int" json:"duration"`
	Lyric     string    `gorm:"type:text" json:"lyric"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
