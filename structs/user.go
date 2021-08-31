package structs

type User struct {
	Id        string `gorm:"primary_key" json:"id,omitempty"`
	Fullname  string `gorm:"size:255" json:"fullname,omitempty"`
	Username  string `gorm:"size:255;not null" json:"username,omitempty"`
	Email     string `gorm:"size:255" json:"email,omitempty"`
	Password  string `gorm:"size:255;not null" json:"-"`
	CreatedAt int64  `gorm:"size:20" json:"created_at,omitempty"`
}
