CREATE TABLE IF NOT EXISTS warehouses(
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    availability integer
);

CREATE TABLE IF NOT EXISTS goods(
    id serial PRIMARY KEY,
    size real CHECK (size > 0),
    code integer UNIQUE NOT NULL,
    amount integer NOT NULL CHECK (amount > 0)
);