BEGIN TRANSACTION;
ALTER TABLE text_items
    ADD FOREIGN KEY (item_id)
        REFERENCES data_items(id);
COMMIT;