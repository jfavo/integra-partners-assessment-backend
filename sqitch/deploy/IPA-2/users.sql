-- Deploy the users table to the integra_partners schema

BEGIN;

CREATE TYPE integra_partners.user_status AS ENUM (
    'I', -- Inactive
    'A', -- Active
    'T'  -- Terminated
);

CREATE TABLE IF NOT EXISTS integra_partners.users (
    user_id     BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_name   VARCHAR(50),
    first_name  VARCHAR(255),
    last_name   VARCHAR(255),
    email       VARCHAR(255),
    user_status integra_partners.user_status,
    department  VARCHAR(255)
);

COMMIT;
