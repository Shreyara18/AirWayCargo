package models

import (
	"github.com/google/uuid"
	"time"
)
 
type Booking struct {
ID uuid.UUID `json:"id;default:uuid_generate_v4()"`
CustomerId	uuid.UUID `json:"customer_id"`
CustomerName string `json:"customer_name"`
ConfigId uuid.UUID `json:"config_id"`
TermsOfShipment string `json:"terms_of_shipment"` 
HSCode int `json:"hs_code"`
CargoReadyDate int64 `json:"cargo_ready_date"`
OriginPort string `json:"origin_port"`
DestinationPort string `json:"destination_port"`
DimensionUnit string `json:"dimension_unit"`
Weight float64 `json:"weight"`
Length float64 `json:"length"`
Width float64 `json:"width"`
Height float64 `json:"height"`
Count int `json:"count"`
Remarks	string `json:"remarks"`
CreatedAt time.Time    `json:"created_at"`
CreatedBy string       `json:"created_by"`
UpdatedAt time.Time    `json:"updated_at"`
UpdatedBy string       `json:"updated_by"`
}