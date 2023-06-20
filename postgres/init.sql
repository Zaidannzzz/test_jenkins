CREATE DATABASE testpostgre;
DROP TABLE IF EXISTS testpostgre;
CREATE SEQUENCE testpostgre_id_seq START 1;
CREATE TABLE testpostgre (
    id bigint DEFAULT nextval('testpostgre_id_seq'),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    full_name text,
    email text,
    password text,
    role text
);
