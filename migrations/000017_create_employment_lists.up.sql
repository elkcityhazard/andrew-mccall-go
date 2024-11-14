create table if not exists employment_lists(
id bigint not null auto_increment primary key,
resume_id bigint not null,
created_at datetime not null default current_timestamp,
updated_at datetime not null default current_timestamp on update current_timestamp,
version int not null default 1,
constraint fk_employment_lists_resume_id foreign key (resume_id) references resumes(id)
);


create index idx_employment_lists_resume_id ON employment_lists(resume_id);