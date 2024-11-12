create table if not exists resumes(
    id bigint auto_increment primary key,
    user_id bigint,
    job_title varchar(255) not null default "",
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    version int not null default 1,
    constraint fk_resume_user_id foreign key (user_id) references users(id) on delete cascade
);

create index idx_resume_user_id ON resumes(user_id);
