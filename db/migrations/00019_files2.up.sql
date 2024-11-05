BEGIN TRANSACTION;
ALTER TABLE file_items
    ADD FOREIGN KEY (item_id)
        REFERENCES data_items(id);
COMMIT;