package models

import (
	"github.com/google/uuid"
	"time"
)

type Config struct {
	ID uuid.UUID `json:"id;default:uuid_generate_v4()"`
    UserName string `json:"username"`
    Email string `json:"email"`
    Password string `json:"password"`
    Mobile int `json:"mobile"`
    Address string `json:"address"`
    AccountType string `json:"account_type"`
	CreatedAt time.Time    `json:"created_at"`
	CreatedBy string       `json:"created_by"`
	UpdatedAt time.Time    `json:"updated_at"`
	UpdatedBy string       `json:"updated_by"`
}
