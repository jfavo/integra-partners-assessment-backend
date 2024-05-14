-- Revert integra-partners-assessment-db:appschema from pg

BEGIN;

-- Drops the schema from the DB
DROP SCHEMA integra_partners;

COMMIT;
