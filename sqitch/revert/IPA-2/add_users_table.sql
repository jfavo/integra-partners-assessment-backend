-- Revert integra-partners-assessment-db:createusers from pg

BEGIN;

DROP TABLE integra_partners.users;
DROP TYPE integra_partners.user_status;

COMMIT;
