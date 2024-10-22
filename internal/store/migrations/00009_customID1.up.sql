BEGIN TRANSACTION;
ALTER TABLE data_items
    ADD COLUMN created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP;
COMMIT;