CREATE SEQUENCE IF NOT EXISTS accounts_id_seq;

CREATE TABLE IF NOT EXISTS accounts(
   id integer NOT NULL DEFAULT nextval('accounts_id_seq'),
   name VARCHAR (50)  NOT NULL,
   email VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (1000) NOT NULL
);

ALTER SEQUENCE accounts_id_seq
OWNED BY accounts.id;