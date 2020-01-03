CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS todo (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
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

CREATE TABLE IF NOT EXISTS avatar {
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    user_id uuid NOT NULL,
    avatar_data TEXT NOT NULL COMMENT 'base64 encoded'
};

CREATE TABLE IF NOT EXISTS user (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    username character(30) NOT NULL UNIQUE,
    avatar_id uuid,
    hash character(150) NOT NULL,
    created_at DATE,
    updated_at DATE
);
