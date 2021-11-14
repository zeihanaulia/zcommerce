CREATE SEQUENCE IF NOT EXISTS payments_id_seq;

CREATE TABLE IF NOT EXISTS payments(
   id integer NOT NULL DEFAULT nextval('payments_id_seq'),
   trx_id VARCHAR (50) UNIQUE NOT NULL,
   reference_trx_id VARCHAR (50) NOT NULL,
   types VARCHAR (50) NOT NULL,
   final_amount decimal (20,2) NOT NULL,
   payloads JSON NOT NULL,
   status VARCHAR (30)  NOT NULL
);

ALTER SEQUENCE payments_id_seq
OWNED BY payments.id;

