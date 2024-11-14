alter table if exists skill_lists
drop foreign key if exists fk_skill_lists_resume_id;

drop index if exists idx_skill_lists_resume_id on skill_lists;

drop table if exists skill_lists;
