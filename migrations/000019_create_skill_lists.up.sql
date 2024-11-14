create table if not exists skill_lists(
id bigint not null auto_increment primary key,
resume_id bigint not null,
title varchar(140) not null default "",
created_at datetime not null default current_timestamp,
updated_at datetime not null default current_timestamp on update current_timestamp,
version int not null default 1,
constraint fk_skill_lists_resume_id foreign key (resume_id) references resumes(id)
);

create index idx_skill_lists_resume_id on skill_lists(resume_id);

