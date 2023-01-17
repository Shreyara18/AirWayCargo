-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS quotes(
id uuid NOT NULL ,
booking_id uuid NOT NULL ,
config_id uuid NOT NULL ,
terms_of_shipment text NOT NULL, 
HS_code	integer NOT NULL ,
cargo_ready_date integer NOT NULL, 
origin_port	text NOT NULL ,
destination_port text NOT NULL, 
dimension_unit text NOT NULL ,
weight Float NOT NULL ,
length Float NOT NULL ,
width	Float NOT NULL ,
height	Float NOT NULL ,
count	Integer NOT NULL, 
total_transit_days	Integer NOT NULL ,
free_days	Integer NOT NULL ,
validity	Integer NOT NULL ,
origin_date	Integer NOT NULL ,
destination_date	Integer NOT NULL ,
partner_id uuid NOT NULL ,
charge text NOT NULL ,
freight_type text NOT NULL ,
currency text NOT NULL ,
partner_rate  float NOT NULL ,
partner_tax float NOT NULL ,
units float NOT NULL ,
customer_rate float NOT NULL ,
customer_tax float NOT NULL ,   
remarks	Text NOT NULL ,
created_at timestamp NOT NULL ,
created_by uuid NOT NULL ,
updated_at timestamp NOT NULL ,
updated_by uuid NOT NULL ,
CONSTRAINT quotes_pkey PRIMARY KEY (id),
CONSTRAINT quotes_booking_id_fkey FOREIGN KEY (booking_id) REFERENCES booking(id)
CONSTRAINT quotes_config_id_fkey FOREIGN KEY (config_id) REFERENCES config(id)

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quotes;
-- +goose StatementEnd
