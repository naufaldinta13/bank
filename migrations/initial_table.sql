Customer
name, nik, no_handphone, account_number, saldo
Transaction
customer_id, nominal, type (in, out), created_at


CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE customer (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    nik name NOT NULL,
    name name NOT NULL,
    phone_number name NOT NULL,
    account_number name NOT NULL,
    saldo double precision,
    is_deleted boolean DEFAULT false,
    CONSTRAINT pk_customer PRIMARY KEY (id)
);


CREATE TABLE transaction (
    id uuid NOT NULL DEFAULT uuid_generate_v4 (),
    customer_id uuid NOT NULL,
    type name NOT NULL,
    nominal double precision,
    created_at timestamp
        with
            time zone DEFAULT now(),
    is_deleted boolean DEFAULT false,
    CONSTRAINT pk_transaction PRIMARY KEY (id),
    CONSTRAINT fk_transaction_1 FOREIGN KEY (customer_id) REFERENCES customer (id) MATCH FULL ON UPDATE NO ACTION ON DELETE NO ACTION
);
