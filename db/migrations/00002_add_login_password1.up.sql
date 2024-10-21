BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS data_items
(
    id       INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    data_type    INT,
    owner_id INT,
    meta_name VARCHAR(200),
    meta_value VARCHAR(200),
    );
ADD FOREIGN KEY (owner_id)
        REFERENCES users (id);
COMMIT;