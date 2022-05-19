-- +goose Up
-- +goose StatementBegin
CREATE TABLE "links"
(
  id SERIAL,
  result_id integer REFERENCES "results" ON DELETE CASCADE,
  link text NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  CONSTRAINT link_pkey PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "links";
-- +goose StatementEnd
