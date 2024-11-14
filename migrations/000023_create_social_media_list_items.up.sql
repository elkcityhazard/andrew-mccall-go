create table if not exists social_media_list_items(
    id bigint not null auto_increment primary key,
    social_media_lists_id bigint not null,
    company_name varchar(140) not null default "",
    username varchar(140) not null default "",
    web_address varchar(140) not null default "",
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    version int not null default 1,
    constraint fk_social_media_list_items_social_media_lists_id foreign key (social_media_lists_id) references social_media_lists(id)
);

create index idx_social_media_list_items_social_media_lists_id ON social_media_list_items(social_media_lists_id);
