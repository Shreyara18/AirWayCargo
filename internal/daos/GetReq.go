package daos

import (
	"Air_Way_Cargo/dtos"
	"Air_Way_Cargo/internal/daos/models"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/wiz-freight-org/adapters/utils/db"
)

type DB struct {
	dbConn *db.DBConn
}

func NewDB(dbConn *db.DBConn) *DB {
	return &DB{
		dbConn: dbConn,
	}
}

func (n *DB) NewCreateUser(con *models.Config) error {
	err := n.dbConn.GetQueryer().QueryRow(context.Background(),
		`INSERT INTO config(id, username, email, mobile, address, password, account_type, created_at, created_by, updated_at, updated_by)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT DO NOTHING)`,
		con.ID, con.UserName, con.Email, con.Mobile, con.Address, con.Password, con.AccountType, con.CreatedAt, con.CreatedBy, con.UpdatedAt, con.UpdatedBy)

	if err != nil {
		return errors.New("User not created")
	}

	return nil
}

func (n *DB) NewLogin(email string, password string) (string, error) {
	Login := new(models.Config)
	err := n.dbConn.GetQueryer().QueryRow(context.Background(),
		`SELECT password FROM  config where email= $1 ON CONFLICT DO NOTHING`, email).Scan(&Login.Password)

	if err != nil {
		return "", errors.New("Login Unsuccessfull")

	}

	if password == Login.Password {
		return "Login Successfull", nil
	} else {
		return "Login Unsuccessfull", nil
	}

}

func (n *DB) NewGetQuote(config_id uuid.UUID) (*dtos.GetResp, error) {
	Quote := new(models.Quotes)
	err := n.dbConn.GetQueryer().QueryRow(context.Background(),
		`SELECT * FROM quotes where config_id= $1 ON CONFLICT DO NOTHING`, config_id).Scan(&Quote.ConfigId)
	if err != nil {
		return nil, errors.New("Cannot fetch details")

	}
	return nil, err

}

func (n *DB) NewCreateQuote(quo *models.Quotes) error {
	err := n.dbConn.GetQueryer().QueryRow(context.Background(),
		`INSERT INTO quotes(id, customer_id, terms_of_shipment , hs_code, cargo_ready_date, origin_port , destination_port, dimension_unit, weight, length , width, height , count, remarks, created_at , created_by, updated_at, updated_by)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11,$12, $13, $14, $15, $16,$17, $18) ON CONFLICT DO NOTHING`,
		quo.ID, quo.CustomerId, quo.TermsOfShipment, quo.HSCode, quo.CargoReadyDate, quo.OriginPort, quo.DestinationPort, quo.DimensionUnit, quo.Weight, quo.Length, quo.Width, quo.Height, quo.Count, quo.Remarks, quo.CreatedAt, quo.CreatedBy, quo.UpdatedAt, quo.UpdatedBy)

	if err != nil {
		return errors.New("Quote not created")
	}

	return nil
}

func (n *DB) NewListQuote(id uuid.UUID) (*dtos.GetResp, error) {
	Quote := new(models.Quotes)
	err := n.dbConn.GetQueryer().QueryRow(context.Background(),
		`SELECT * FROM quotes where id= $1 ON CONFLICT DO NOTHING`, id).Scan(&Quote.ID)
	if err != nil {
		return nil, errors.New("Cannot fetch details")

	}
	return nil, err

}

func (n *DB) NewGetBooking(config_id uuid.UUID) (*dtos.GetResp, error) {
	Booking := new(models.Booking)
	err := n.dbConn.GetQueryer().QueryRow(context.Background(),
		`SELECT * FROM booking where config_id= $1 ON CONFLICT DO NOTHING`, config_id).Scan(&Booking.ConfigId)
	if err != nil {
		return nil, errors.New("Cannot fetch details")

	}
	return nil, err
}

func (n *DB) NewCreateBooking(book *models.Booking) error {
	err := n.dbConn.GetQueryer().QueryRow(context.Background(),
		`INSERT INTO booking(id, customer_id, customer_name, config_id, terms_of_shipment , hs_code, cargo_ready_date, origin_port , destination_port, dimension_unit, weight, length , width, height , count, remarks, created_at , created_by, updated_at, updated_by)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11,$12, $13, $14, $15, $16,$17, $18, $19, $20) ON CONFLICT DO NOTHING`,
		book.ID, book.CustomerId, book.CustomerName, book.ConfigId, book.TermsOfShipment, book.HSCode, book.CargoReadyDate, book.OriginPort, book.DestinationPort, book.DimensionUnit, book.Weight, book.Length, book.Width, book.Height, book.Count, book.Remarks, book.CreatedAt, book.CreatedBy, book.UpdatedAt, book.UpdatedBy)

	if err != nil {
		return errors.New("Booking not created")
	}

	return nil
}

func (n *DB) NewListBooking(id uuid.UUID) (*dtos.GetResp, error) {
	Booking := new(models.Task)
	err := n.dbConn.GetQueryer().QueryRow(context.Background(),
		`SELECT * FROM task where id= $1 ON CONFLICT DO NOTHING`, id).Scan(&Booking.ID)
	if err != nil {
		return nil, errors.New("Cannot fetch details")

	}
	return nil, err
}
