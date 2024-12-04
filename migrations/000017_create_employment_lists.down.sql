alter table if exists employment_lists
drop foreign key if exists fk_employment_lists_resume_id;
drop index if exists idx_employment_lists_resume_id ON employment_lists;
drop table if exists employment_lists;
