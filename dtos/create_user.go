package dtos

type CreateUserRequest struct {
	LanguageCode string `db:"language_code" json:"language_code"`
	Firstname    string `db:"firstname" json:"firstname"`
	Lastname     string `db:"lastname" json:"lastname"`
} //@name CreateUserRequest
