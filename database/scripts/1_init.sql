CREATE DATABASE manager;
\c manager

CREATE EXTENSION unaccent
    SCHEMA "public"
    VERSION "1.1";

CREATE EXTENSION "uuid-ossp"
    SCHEMA "public"
    VERSION "1.1";

CREATE OR REPLACE FUNCTION public.tf_set_updated_at()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
DECLARE 
  tabela_raw TEXT;
  tabela TEXT; 
  temp_mensagem TEXT;
BEGIN
  IF (TG_OP = 'UPDATE') THEN
    NEW.updated_at = now();
    RETURN NEW;
  ELSE 
  
    tabela_raw := replace(TG_TABLE_NAME, 't_', '');
    tabela := replace(tabela_raw, '_', ' ');
    temp_mensagem := 'Atributo updated_at, atualização NãO autorizada para gatilhos diferentes de UPDATE para ('||tabela||') !!'; 
    RAISE EXCEPTION feature_not_supported USING HINT = temp_mensagem;

  END IF;
END;

$function$
;