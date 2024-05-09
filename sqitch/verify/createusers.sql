-- Verify integra-partners-assessment-db:createusers on pg

BEGIN;

-- Verify that the user_status enum type exists
DO $$
BEGIN
    ASSERT (SELECT 1 FROM pg_type WHERE typname = 'user_status');
END $$;

-- Verify the values of user_status are correct
-- Currently, we expect 'I', 'A', and 'T' to be values
DO $$
DECLARE
    expected_values text[] := ARRAY['I','A','T'];
BEGIN
    -- Leveraging array operator to compare the values of the existing enum
    -- to our expected values array
    ASSERT (SELECT enum_range(NULL::integra_partners.user_status)::text[] <@ expected_values);
END $$;


ROLLBACK;
