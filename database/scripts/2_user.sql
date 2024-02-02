\c manager

CREATE SCHEMA "user";

CREATE TABLE IF NOT EXISTS "user".t_user(
  id UUID NOT NULL DEFAULT uuid_generate_v4(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NULL,
  deleted_at TIMESTAMPTZ NULL,
  "name" TEXT NOT NULL,
  age INT,
  document TEXT NOT NULL,
  CONSTRAINT t_user_pkey PRIMARY KEY (id)
);
CREATE TRIGGER set_updated_at BEFORE
UPDATE ON "user".t_user FOR EACH ROW EXECUTE PROCEDURE public.tf_set_updated_at();

CREATE TYPE "user".e_contact_type AS ENUM('telephone', 'cellphone', 'email', 'whatsapp', 'instagram');
CREATE TABLE IF NOT EXISTS "user".t_user_contact(
  id UUID NOT NULL DEFAULT uuid_generate_v4(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NULL,
  deleted_at TIMESTAMPTZ NULL,
  user_id UUID NOT NULL,
  "target" TEXT NOT NULL,
  "type" "user".e_contact_type NOT NULL,

  CONSTRAINT t_user_contact_pkey PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES "user".t_user (id)
);
CREATE TRIGGER set_updated_at BEFORE
UPDATE ON "user".t_user_contact FOR EACH ROW EXECUTE PROCEDURE public.tf_set_updated_at();