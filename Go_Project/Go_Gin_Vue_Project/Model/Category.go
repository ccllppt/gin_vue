package Model

// Category 是分类模型，对应数据库中的分类表
type Category struct {
	// gorm.Model // 如果使用 GORM 的默认模型，可以取消注释
	Id        uint   `json:"id" gorm:"primary_key"`                        // 分类 ID，主键
	Name      string `json:"name" gorm:"type:varchar(50);not null;unique"` // 分类名称，唯一且不能为空
	CreatedAt Time   `json:"created_at" gorm:"type:timestamp;"`            // 创建时间
	UpdatedAt Time   `json:"updated_at" gorm:"type:timestamp;"`            // 更新时间
}
