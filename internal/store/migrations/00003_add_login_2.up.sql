BEGIN TRANSACTION;
ALTER TABLE data_items
    ADD FOREIGN KEY (owner_id)
REFERENCES users(id);
COMMIT;