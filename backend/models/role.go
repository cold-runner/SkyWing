package models

type RoleCharacter struct {
	Uuid uint64 `json:"uuid" db:"uuid"`
	Role string `json:"role" db:"role"`
}
