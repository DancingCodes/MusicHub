package model

import "time"

type MusicRecord struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255);index;not null" json:"name"`
	PicUrl    string    `gorm:"type:varchar(500)" json:"pic_url"`
	Artists   string    `gorm:"type:varchar(255)" json:"artists"`
	Duration  int       `gorm:"type:int" json:"duration"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
