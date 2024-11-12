create table if not exists award_items(
    id bigint not null auto_increment primary key,
    award_list_id bigint not null,
    title varchar(140) not null default "",
    org_name varchar(140) not null default "",
    received_year year not null default year(current_date),
    content text not null default "",
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    version int not null default 1,
    constraint fk_award_items_award_list_id foreign key (award_list_id) references award_lists(id)
);

create index idx_award_items_award_list_id on award_items(award_list_id);
