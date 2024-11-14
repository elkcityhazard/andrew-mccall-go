create table if not exists employment_list_items(
    id bigint not null auto_increment primary key,
    employment_lists_id bigint not null,
    title varchar(140) not null default "",
    date_from year not null default year(current_date),
    date_to year not null default year(current_date),
    check (date_to >= date_from),
    job_title varchar(140) not null default "",
    summary text not null default "",
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    version int not null default 1,
    constraint fk_employment_list_items_employment_lists_id foreign key (employment_lists_id) references employment_lists(id)
);

create index idx_employment_list_items_employment_lists_id ON employment_list_items(employment_lists_id);
