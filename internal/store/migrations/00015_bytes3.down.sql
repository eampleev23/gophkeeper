BEGIN TRANSACTION;
ALTER TABLE login_password_items
DROP COLUMN IF EXISTS login;
COMMIT;