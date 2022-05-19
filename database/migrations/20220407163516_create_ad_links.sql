-- +goose Up
-- +goose StatementBegin
CREATE TABLE "ad_links"
(
  id SERIAL,
  result_id integer REFERENCES "results" ON DELETE CASCADE,
  type text NOT NULL,
  position text NOT NULL,
  link text NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  CONSTRAINT ad_link_pkey PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "ad_links";
-- +goose StatementEnd
