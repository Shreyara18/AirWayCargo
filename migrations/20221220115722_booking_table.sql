-- +goose Up
-- +goose StatementBegin
 CREATE TABLE IF NOT EXISTS booking (
    id	uuid NOT NULL ,
customer_id	uuid NOT NULL, 
customer_name text NOT NULL,
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
remarks	Text NULL,
created_at timestamp NOT NULL ,
created_by uuid NOT NULL ,
updated_at	timestamp NOT NULL, 
updated_by	uuid NOT NULL ,
 CONSTRAINT booking_pkey PRIMARY KEY (id),
 CONSTRAINT booking_config_id_fkey FOREIGN KEY (config_id) REFERENCES config(id)
 );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS booking;
-- +goose StatementEnd
