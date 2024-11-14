alter table if exists comments
drop foreign key if exists fk_comments_user_id,
drop foreign key if exists fk_comments_post_id;

drop index if exists idx_comments_user_id on comments;
drop index if exists idx_comments_post_id on comments;

DROP table if exists comments;
