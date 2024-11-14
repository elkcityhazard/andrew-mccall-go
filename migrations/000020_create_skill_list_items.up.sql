create table if not exists skill_list_items(
    id bigint not null auto_increment primary key,
    skill_lists_id bigint not null,
    title varchar(140) not null default "",
    content text not null default "",
    duration year not null default year(current_date),
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    version int not null default 1,
    constraint fk_skill_list_items_skill_lists_id foreign key (skill_lists_id) references skill_lists(id) on delete cascade
);

create index idx_skill_list_items_skill_lists_id ON skill_list_items(skill_lists_id);
