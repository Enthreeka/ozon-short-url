CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS "user"
(
    id uuid default uuid_generate_v4(),
    primary key (id)
);

CREATE TABLE IF NOT EXISTS url
(
    id serial,
    original_url text unique,
    short_url varchar(30) unique,
    primary key (id)
);

SELECT original_url FROM url WHERE short_url = 'E1RJ8_FIda';