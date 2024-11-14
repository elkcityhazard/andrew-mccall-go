create table if not exists education_items(
    id bigint not null auto_increment primary key,
    education_lists_id bigint not null,
    name varchar(140) not null default "",
    degree_year year not null default year(current_date),
    degree varchar(140) not null default "",
    address_1 varchar(140) not null default "",
    address_2 varchar(140) not null default "",
    city varchar(140) not null default "",
    state varchar(140) not null default "",
    zipcode varchar(140) not null default "",
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    version int not null default 1,
    constraint fk_education_lists_id foreign key (education_lists_id) references education_lists(id)
);

create index idx_education_lists_id ON education_items(education_lists_id);
