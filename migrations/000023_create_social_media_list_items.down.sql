alter table if exists social_media_list_items
drop foreign key if exists fk_social_media_list_items_social_media_lists_id;

drop index if exists idx_social_media_list_items_social_media_lists_id ON social_media_list_items;

drop table if exists social_media_list_items;
