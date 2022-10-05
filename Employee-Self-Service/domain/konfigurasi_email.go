package domain

type EmailValidation struct {
	IdEmailValidation   int64  `gorm:"PrimaryKey"`
	NamaEmailValidation string `db:"nama_email_validation"`
}

func (EmailValidation) TableName() string {
	return "email_validation"
}
