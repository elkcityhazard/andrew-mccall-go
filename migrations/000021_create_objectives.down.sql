alter table if exists objectives
drop foreign key if exists fk_objectives_resume_id;

drop index if exists idx_objectives_resume_id on objectives;
