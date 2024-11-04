BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS text_items (
                                 item_id INT,
                                 nonce_text_content VARCHAR (200),
                                 text_content TEXT
);
COMMIT;