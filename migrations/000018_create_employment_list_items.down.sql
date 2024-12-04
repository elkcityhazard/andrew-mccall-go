alter table if exists employment_list_items
drop foreign key if exists fk_employment_list_items_employment_lists_id;
drop index if exists idx_employment_list_items_employment_lists_id ON employment_list_items;

drop table if exists employment_list_items;
