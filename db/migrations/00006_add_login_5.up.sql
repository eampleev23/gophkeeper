BEGIN TRANSACTION;
alter table login_password_items
    add constraint login_password_items_data_items_id_fk
        foreign key (item_id) references data_items;
COMMIT;