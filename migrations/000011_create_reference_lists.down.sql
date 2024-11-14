alter table reference_lists
drop foreign key if exists fk_resume_id;

drop index if exists idx_reference_lists_resume_id on reference_lists;

drop table if exists reference_lists;
