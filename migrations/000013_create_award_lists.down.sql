alter table if exists award_lists
drop foreign key fk_award_lists_user_id,
drop foreign key fk_award_lists_resume_id;

drop table if exists award_lists;
