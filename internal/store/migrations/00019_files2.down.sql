BEGIN TRANSACTION;
ALTER TABLE file_items
DROP COLUMN IF EXISTS item_id;
COMMIT;