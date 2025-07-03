package models

type User struct {
	ID       string    `json:"id" gorm:"primaryKey"`
	Email    string    `json:"email" gorm:"not null,unique"`
	Password string    `json:"password" gorm:"not null"`
	Sessions []Session `gorm:"foreignkey:UserId" json:"-"`
	Name     string    `gorm:"name"`
}

type Session struct {
	ID      string `json:"id" gorm:"primaryKey"`
	UserId  string `json:"userId"`
	User    User   `gorm:"foreignkey:UserId" json:"-"`
	Token   string `json:"token"`
	Expires int64  `json:"expires"`
}