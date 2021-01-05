CREATE TABLE "user"
(
  id SERIAL,
  email text UNIQUE,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  CONSTRAINT user_pkey PRIMARY KEY (id)
);
