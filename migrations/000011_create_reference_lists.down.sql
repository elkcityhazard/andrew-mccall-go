alter table reference_lists
drop foreign key fk_resume_id,
drop foreign key fk_user_id;

drop table if exists reference_lists;
