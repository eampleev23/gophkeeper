BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS login_password_items
(
    id       INT,
    hash_login    VARCHAR(200),
    hash_password VARCHAR(200)

    );
COMMIT;