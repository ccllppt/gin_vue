package Model

type Category struct {
	//gorm.Model
	Id        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreatedAt Time   `json:"created_at" gorm:"type:timestamp;"`
	UpdatedAt Time   `json:"updated_at" gorm:"type:timestamp;"`
}
