package domain

type Category struct {
	ID   int32  `gorm:"primaryKey;column:id;autoIncrement"`
	Name string `gorm:"column:name"`
}

func (category *Category) TableName() string {
	return "categories"
}
