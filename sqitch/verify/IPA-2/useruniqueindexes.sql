-- Verify integra-partners-assessment-db:createusers on pg

BEGIN;

-- Verify that the user_status enum type exists
DO $$
BEGIN
    ASSERT (
        SELECT 1
        FROM pg_indexes
        WHERE schemaname = 'integra_partners'
        AND indexname = 'users_user_name_idx'
    );

    ASSERT (
        SELECT 1
        FROM pg_indexes
        WHERE schemaname = 'integra_partners'
        AND indexname = 'users_email_idx'
    );
END $$;

ROLLBACK;