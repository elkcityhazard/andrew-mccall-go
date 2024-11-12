create table if not exists reference_items(
    id bigint not null auto_increment primary key,
    ref_list_id bigint not null,
    first_name varchar(140) not null default "",
    last_name varchar(140) not null default "",
    email varchar(140) not null default "",
    phone_number varchar(140) not null default "",
    job_title varchar(140) not null default "",
    organization varchar(140) not null default "",
    type enum("personal", "professional", "academic", "industry") not null default "personal",
    address_1 varchar(140) not null default "",
    address_2 varchar(140) not null default "",
    city varchar(140) not null default "",
    state varchar(140) not null default "",
    zipcode varchar(140) not null default "",
    content text not null default "",
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    version int not null default 1,
    constraint fk_ref_items_ref_list_id foreign key (ref_list_id) references reference_lists(id)
);

create index idx_ref_items_ref_list_id on reference_items(ref_list_id);
