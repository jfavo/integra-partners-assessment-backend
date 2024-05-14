-- Verify integra-partners-assessment-db:appschema on pg

BEGIN;

-- Using has_schema_privilege to determine if our schema exists
-- Will throw an exception if it does not.
DO $$
BEGIN
    ASSERT (SELECT has_schema_privilege('integra_partners', 'usage'));
END $$;

ROLLBACK;
