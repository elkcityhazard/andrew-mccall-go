create table if not exists reference_lists(
    id bigint not null auto_increment primary key,
    resume_id bigint not null,
    user_id bigint not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    version int not null default 1,
    constraint fk_resume_id foreign key (resume_id) references resumes(id) on delete cascade,
    constraint fk_user_id foreign key (user_id) references users(id) on delete cascade
);

create index idx_reference_lists_resume_id on reference_lists(resume_id);
create index idx_reference_lists_user_id on reference_lists(user_id);
