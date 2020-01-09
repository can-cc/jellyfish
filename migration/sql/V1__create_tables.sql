CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS app_user (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    username character(30) NOT NULL UNIQUE,
    avatar TEXT,
    hash character(150) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

comment on column app_user.avatar is 'base64 encoded data';

CREATE TABLE IF NOT EXISTS todo (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    creator_id uuid NOT NULL,
    content character(120) NOT NULL,
    detail character(500),
    deadline TIMESTAMP,
    status character(20),
    type character(20),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (creator_id) REFERENCES app_user (id)
);

