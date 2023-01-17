-- +goose Up
-- +goose StatementBegin
 
 CREATE TABLE IF NOT EXISTS config (
    id uuid NOT NULL ,
    username text not NULL,
    email text not NULL,
    password text NOT NULL,
    mobile integer NOT NULL,
    address text NOT NULL,
    account_type text NOT NULL,
    created_at timestamp NOT NULL,
    created_by uuid NOT NULL,
    updated_at timestamp NOT NULL,
    updated_by uuid NOT NULL,
    CONSTRAINT config_pkey PRIMARY KEY (id)

 );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS config;
-- +goose StatementEnd
