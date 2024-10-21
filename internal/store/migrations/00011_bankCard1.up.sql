BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS bank_card_items
(
    item_id       INT,
    card_number    VARCHAR(200),
    valid_thru VARCHAR(200),
    owner_name VARCHAR(200),
    cvc VARCHAR(200),
    nonce_card_number VARCHAR,
    nonce_valid_thru VARCHAR,
    nonce_owner_name VARCHAR,
    nonce_cvc VARCHAR

    );
COMMIT;