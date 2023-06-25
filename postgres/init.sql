SELECT 'CREATE DATABASE testpostgre' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'testpostgre');
\c testpostgre;
CREATE SEQUENCE User_id_seq START 1;
CREATE TABLE "User" (
    id bigint DEFAULT nextval('User_id_seq'),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    full_name text,
    email text,
    password text
);
