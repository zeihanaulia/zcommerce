CREATE SEQUENCE IF NOT EXISTS order_id_seq;

CREATE TABLE IF NOT EXISTS orders(
   id integer NOT NULL DEFAULT nextval('order_id_seq'),
   trx_id VARCHAR (50) UNIQUE NOT NULL,
   payment_trx_id VARCHAR (50) NOT NULL,
   lock_items JSON NOT NULL,
   status VARCHAR (30)  NOT NULL,
   customer_name VARCHAR (50)  NOT NULL,
   customer_address VARCHAR (50)  NOT NULL
);

ALTER SEQUENCE order_id_seq
OWNED BY orders.id;