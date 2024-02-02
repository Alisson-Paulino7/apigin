CREATE SCHEMA "user";

CREATE TABLE IF NOT EXISTS "user".t_user(
  id UUID NOT NULL DEFAULT uuid_generate_v4(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NULL,
  deleted_at TIMESTAMPTZ NULL,
  "name" TEXT NOT NULL,
  age INT NOT NULL,
  document TEXT NOT NULL,
  CONSTRAINT t_user_pkey PRIMARY KEY (id)
);