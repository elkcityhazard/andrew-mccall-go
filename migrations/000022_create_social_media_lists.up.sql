create table if not exists social_media_lists(
id bigint not null auto_increment primary key,
resume_id bigint not null,
created_at datetime not null default current_timestamp,
updated_at datetime not null default current_timestamp on update current_timestamp,
version int not null default 1,
constraint fk_social_media_lists_resume_id foreign key (resume_id) references resumes(id) on delete cascade
);

create index idx_social_media_lists_resume_id ON social_media_lists(resume_id);
