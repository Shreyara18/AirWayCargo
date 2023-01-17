package models

import (
	"github.com/google/uuid"
	"time"
)

type Quotes struct{
ID uuid.UUID `json:"id;default:uuid_generate_v4()"`
BookingId uuid.UUID `json:"booking_id"`
ConfigId uuid.UUID `json:"config_id"`
CustomerId	uuid.UUID `json:"customer_id"`
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
TotalTransitDays int `json:"total_transit_days"`
FreeDays int `json:"free_days"`
Validity int64 `json:"validity"`
OriginDate	int64 `json:"origin_date"`
DestinationDate	int64 `json:"destination_date"`
PartnerId uuid.UUID `json:"partner_id"`
Charge string `json:"charge"`
FreightType string `json:"freight_type"`
Currency string `json:"currency"`
PartnerRate float64 `json:"partner_rate"`
PartnerTax float64 `json:"partner_tax"`
Units float64 `json:"units"`
CustomerRate float64 `json:"customer_rate"`
Customer_tax float64 `json:"customer_tax"`  
Remarks	string `json:"remarks"`
CreatedAt time.Time    `json:"created_at"`
CreatedBy string       `json:"created_by"`
UpdatedAt time.Time    `json:"updated_at"`
UpdatedBy string       `json:"updated_by"`
}