package dtos

import (
	"time"

	"github.com/google/uuid"
)

type GetResp struct {
	ID               uuid.UUID `json:"id"`
	OriginPort       string    `json:"origin_port"`
	DestinationPort  string    `json:"destination_port"`
	TermsOfShipment  string    `json:"terms_of_shipment"`
	AccountType      string    `json:"account_type"`
	Validity         int64     `json:"validity"`
	PartnerId        uuid.UUID `json:"partner_id"`
	TotalTransitDays int       `json:"total_transit_days"`
	FreeDays         int       `json:"free_days"`
	OriginDate       int64     `json:"origin_date"`
	DestinationDate  int64     `json:"destination_date"`
	PartnerRate      float64   `json:"partner_rate"`
	PartnerTax       float64   `json:"partner_tax"`
	Units            float64   `json:"units"`
	CustomerRate     float64   `json:"customer_rate"`
	Customer_tax     float64   `json:"customer_tax"`
	BookingCreated   bool      `json:"booking_created"`
	FlightDeparture  bool      `json:"flight_departure"`
	FlightArrival    bool      `json:"flight_arrival"`
	BookingCompleted bool      `json:"booking_completed"`
	CustomerName     string    `json:"customer_name"`
	CreatedAt        time.Time `json:"created_at"`
	CreatedBy        string    `json:"created_by"`
	UpdatedAt        time.Time `json:"updated_at"`
	UpdatedBy        string    `json:"updated_by"`
}
