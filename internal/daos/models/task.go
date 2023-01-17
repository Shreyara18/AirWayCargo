package models

import (
	"github.com/google/uuid"
	"time"
)

type Task struct{
ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
BookingId uuid.UUID `json:"booking_id"`
BookingCreated	bool `json:"booking_created"`
FlightDeparture	bool `json:"flight_departure"`
FlightArrival	bool  `json:"flight_arrival"`
BookingCompleted bool  `json:"booking_completed"`
CreatedAt time.Time    `json:"created_at"`
CreatedBy string       `json:"created_by"`
UpdatedAt time.Time    `json:"updated_at"`
UpdatedBy string       `json:"updated_by"`
}