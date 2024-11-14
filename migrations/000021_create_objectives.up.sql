create table if not exists objectives(
    id bigint not null auto_increment primary key,
    resume_id bigint not null,
    content text not null default "",
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp, 
    version int not null default 1,
    constraint fk_objectives_resume_id foreign key (resume_id) references resumes(id)
);

create index idx_objectives_resume_id ON objectives(resume_id);
