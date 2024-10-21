BEGIN TRANSACTION;
alter table login_password_items
    add nonce_password VARCHAR;
COMMIT;