
alter table if exists skill_list_items
drop foreign key fk_skill_list_items_skill_lists_id;

drop index if exists idx_skill_list_items_skill_lists_id on skill_list_items;

drop table if exists skill_list_items;
