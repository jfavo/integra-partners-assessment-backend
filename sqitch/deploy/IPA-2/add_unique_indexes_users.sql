-- Creates unique indexes for the user_name and email columns of the users table

BEGIN;

CREATE UNIQUE INDEX users_user_name_idx ON integra_partners.users (LOWER(user_name));
CREATE UNIQUE INDEX users_email_idx ON integra_partners.users (LOWER(email));

COMMIT;
