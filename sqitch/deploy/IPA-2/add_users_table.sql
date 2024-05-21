-- Deploy the users table to the integra_partners schema

BEGIN;

CREATE TYPE integra_partners.user_status AS ENUM (
    'I', -- Inactive
    'A', -- Active
    'T'  -- Terminated
);

CREATE TABLE IF NOT EXISTS integra_partners.users (
    user_id     BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    user_name   VARCHAR(50) NOT NULL,
    first_name  VARCHAR(255) NOT NULL,
    last_name   VARCHAR(255) NOT NULL,
    email       VARCHAR(255) NOT NULL,
    user_status integra_partners.user_status,
    department  VARCHAR(255)
);

COMMIT;
