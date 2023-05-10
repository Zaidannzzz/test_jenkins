DROP TABLE IF EXISTS test;
CREATE SEQUENCE User_id_seq START 1;
CREATE TABLE test (
    id bigint NOT NULL DEFAULT PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    full_name text,
    email text,
    password text,
    role text
);