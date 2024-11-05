BEGIN TRANSACTION;
ALTER TABLE bank_card_items DROP CONSTRAINT bank_card_items_data_items_id_fk;
COMMIT;