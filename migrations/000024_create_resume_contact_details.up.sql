create table if not exists resume_contact_details(
    id bigint not null auto_increment primary key,
    resume_id bigint not null,
    first_name varchar(140) not null default "",
    last_name varchar(140) not null default "",
    address_1 varchar(140) not null default"",
    address_2 varchar(140) not null default"",
    city varchar(140) not null default"",
    state varchar(140) not null default"",
    zipcode varchar(140) not null default"",
    email varchar(140) not null default"",
    phone_number varchar(140) not null default"",
    web_address varchar(140) not null default"",
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    version int not null default 1,
    constraint fk_resume_contact_details_resume_id foreign key (resume_id) references resumes(id)
);

create index idx_resume_contact_details_resume_id on resume_contact_details(resume_id);
