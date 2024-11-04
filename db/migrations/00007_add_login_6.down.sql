BEGIN TRANSACTION;
alter table login_password_items
    drop column nonce_login;
COMMIT;