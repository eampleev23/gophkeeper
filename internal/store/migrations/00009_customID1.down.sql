BEGIN TRANSACTION;
ALTER TABLE data_items
DROP COLUMN IF EXISTS created_at,
    DROP COLUMN IF EXISTS updated_at;
COMMIT;