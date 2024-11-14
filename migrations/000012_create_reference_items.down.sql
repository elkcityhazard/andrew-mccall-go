alter table if exists reference_items
drop foreign key if exists fk_ref_items_ref_list_id;

drop table if exists reference_items;
