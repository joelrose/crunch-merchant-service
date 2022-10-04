package dtos

type WhitelistRequest struct {
	Id string `json:"identifier" validate:"required"`
} // @Name WhitelistRequest
