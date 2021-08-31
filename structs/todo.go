package structs

type Todo struct {
	Id          string `gorm:"primary_key" json:"id,omitempty"`
	Title       string `gorm:"size:255;not null" json:"title,omitempty"`
	Description string `gorm:"size:255" json:"description,omitempty"`
	CreatedAt   int64  `gorm:"size:20" json:"created_at,omitempty"`
}