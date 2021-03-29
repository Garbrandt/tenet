package model

import "time"

var ContentTableName = "contents"

type Content struct {
	ID        int64      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`        // 内容id
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // 内容创建时间
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` // 内容修改时间
	DeletedAt *time.Time `json:"deletedAt"`                                   // 内容删除时间

	Key           string `json:"key"`                        // 内容标记id
	Title         string `json:"title"`                      // 内容标题
	State         string `json:"state"`                      // 文章隐藏，草稿
	Type          string `json:"type"`                       // 内容类型
	Abstract      string `gorm:"type:text;" json:"abstract"` // 内容简介
	ContentUrl    string `json:"content_url"`                // 内容url
	Body          string `gorm:"type:text;" json:"body"`     // 内容详情
	SourceUrl     string `json:"source_url"`                 // 原始url
	Size          int64  `json:"size"`                       // 文件大小
	CopyrightStat string `json:"copyright_stat"`             // 来源状态
}

func (content *Content) TableName() string {
	return ContentTableName
}
