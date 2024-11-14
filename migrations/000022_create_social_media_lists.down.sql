alter table if exists social_media_lists
drop foreign key if exists fk_social_media_lists_resume_id;

drop index if exists idx_social_media_lists_resume_id on social_media_lists;

drop table if exists social_media_lists;
