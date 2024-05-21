-- Drops indexes for user_name and email columns from the users table

BEGIN;

DROP INDEX integra_partners.users_user_name_idx;
DROP INDEX integra_partners.users_email_idx;

COMMIT;
