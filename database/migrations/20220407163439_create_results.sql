-- +goose Up
-- +goose StatementBegin
CREATE TABLE "results"
(
  id SERIAL,
  user_id integer REFERENCES "users" ON DELETE CASCADE,
  keyword text NOT NULL,
  status text NOT NULL,
  page_cache text,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  CONSTRAINT result_pkey PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "results";
-- +goose StatementEnd
