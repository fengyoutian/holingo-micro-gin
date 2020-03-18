package model

type Article struct {
	ID         uint64 `gorm:"primary_key;auto_increment"`                                                   // 主键
	Author     string `gorm:"type:varchar(20)"`                                                             // 作者
	Title      string `gorm:"type:varchar(64);not null"`                                                    // 标题
	Content    string `gorm:"type:varchar(255);not nul"`                                                    // 内容
	ModifyTime string `gorm:"type:datetime;not null;default:current_timestamp on update current_timestamp"` // 修改时间
	CreateTime string `gorm:"type:datetime;not null;default:current_timestamp"`                             // 创建时间
}

func (Article) TableName() string {
	return "article"
}
