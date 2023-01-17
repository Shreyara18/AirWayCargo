-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS task (
id	uuid  NOT NULL ,
booking_id	uuid  NOT NULL , 
BookingCreated	bool NULL DEFAULT true ,
FlightDeparture	bool  NOT NULL ,
FlightArrival	bool  NOT NULL ,
BookingCompleted bool  NOT NULL ,
created_at timestamp NOT NULL ,
created_by uuid NOT NULL ,
updated_at timestamp NOT NULL ,
updated_by uuid NOT NULL ,
CONSTRAINT task_pkey PRIMARY KEY (id),
CONSTRAINT task_booking_id_fkey FOREIGN KEY (booking_id) REFERENCES booking(id)
 );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXIST task;
-- +goose StatementEnd
