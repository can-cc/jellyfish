CREATE TABLE IF NOT EXISTS taco_box (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar(30) NOT NULL,
    creator_id text NOT NULL,
    icon varchar(20),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

ALTER TABLE taco ADD COLUMN box_id uuid;