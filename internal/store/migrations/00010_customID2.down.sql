BEGIN TRANSACTION;
ALTER TABLE data_items
    ADD COLUMN meta_name VARCHAR(200);
COMMIT;