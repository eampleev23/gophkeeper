BEGIN TRANSACTION;
alter table login_password_items
    add nonce_login VARCHAR;
COMMIT;