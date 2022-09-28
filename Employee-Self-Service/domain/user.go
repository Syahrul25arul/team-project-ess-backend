package domain

type User struct {
	IdUser         int64  `gorm:"primaryKey"`
	Email          string `db:"email"`
	Password       string `db:"password"`
	UserRole       string `db:"user_role"`
	StatusVerified string `db:"status_verified"`
}
