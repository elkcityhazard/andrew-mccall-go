alter table if exists award_lists
drop foreign key if exists fk_award_lists_user_id,
drop foreign key if exists fk_award_lists_resume_id;

drop index if exists idx_award_lists_resume_id on award_lists;

drop table if exists award_lists;
