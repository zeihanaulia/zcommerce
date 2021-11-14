CREATE SEQUENCE IF NOT EXISTS order_detail_id_seq;

CREATE TABLE IF NOT EXISTS order_detail(
   id integer NOT NULL DEFAULT nextval('order_detail_id_seq'),
   order_id integer NOT NULL,
   name VARCHAR (50) NOT NULL,
   quantity integer NOT NULL,
   price  DECIMAL (20,2) NOT NULL
);

ALTER SEQUENCE order_detail_id_seq
OWNED BY orders.id;