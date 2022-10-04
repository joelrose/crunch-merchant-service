package dtos

type CreateUserRequest struct {
	LanguageCode string `db:"language_code" json:"language_code" validate:"required,len=2"`
	Firstname    string `db:"firstname" json:"firstname" validate:"required"`
	Lastname     string `db:"lastname" json:"lastname" validate:"required"`
} //@name CreateUserRequest
