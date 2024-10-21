BEGIN TRANSACTION;
alter table bank_card_items
    add constraint bank_card_items_data_items_id_fk
        foreign key (item_id) references data_items;
COMMIT;