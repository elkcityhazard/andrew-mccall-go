alter table if exists resume_contact_details
drop foreign key if exists fk_resume_contact_details_resume_id;

drop index if exists idx_resume_contact_details_resume_id on resume_contact_details;

drop table if exists resume_contact_details;
