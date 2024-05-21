-- Deploy integra-partners-assessment-db:appschema to pg

BEGIN;

-- Add the schema to store our tables
CREATE SCHEMA integra_partners;

COMMIT;
