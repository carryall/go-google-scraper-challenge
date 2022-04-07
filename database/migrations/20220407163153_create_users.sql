-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION citext;
CREATE TABLE "users"
(
  id SERIAL,
  email citext UNIQUE,
  hashed_password text,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  CONSTRAINT user_pkey PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users";
-- +goose StatementEnd
