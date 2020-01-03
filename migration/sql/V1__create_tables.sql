CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS todos (
    id uuid DEFAULT uuid_generate_v4 () NOT NULL PRIMARY KEY,
    creater_id uuid NOT NULL,
    content character(120) NOT NULL,
    detail character(500),
    deadline DATE,
    status character(20),
    type character(20),
    done INTEGER DEFAULT 0,
    created_at DATE,
    updated_at DATE
);

CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4 () NOT NULL PRIMARY KEY,
    username character(30) NOT NULL UNIQUE,
    avatar character(150),
    hash character(150) NOT NULL,
    created_at DATE,
    updated_at DATE
);
