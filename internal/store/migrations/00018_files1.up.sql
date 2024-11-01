BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS file_items (
                                          item_id INT,
                                          server_path VARCHAR (200)
    );
COMMIT;