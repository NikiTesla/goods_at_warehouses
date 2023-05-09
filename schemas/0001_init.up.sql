CREATE TABLE IF NOT EXISTS warehouses(
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    availability boolean
);

CREATE TABLE IF NOT EXISTS goods(
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    size real CHECK (size > 0),
    code integer UNIQUE NOT NULL,
    amount integer NOT NULL CHECK (amount >= 0)
);

CREATE TABLE IF NOT EXISTS warehouse_goods(
    id serial PRIMARY KEY,
    warehouse_id INTEGER REFERENCES warehouses(id),
    goods_id INTEGER REFERENCES goods(code),
    available_amount INTEGER NOT NULL CHECK(available_amount >= 0),
    reserved_amount INTEGER NOT NULL DEFAULT 0
        CHECK(reserved_amount < available_amount AND reserved_amount >= 0)
);