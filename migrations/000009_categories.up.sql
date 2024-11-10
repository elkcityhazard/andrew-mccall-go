create table if not exists categories(
    id bigint not null AUTO_INCREMENT PRIMARY KEY,
    name varchar(140) not null unique default "",
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version int not null default 1
);



create table if not exists category_joins(
    id bigint not null AUTO_INCREMENT PRIMARY KEY,
    cat_id bigint not null default 0,
    post_id bigint not null default 0,
    FOREIGN KEY (cat_id) REFERENCES categories(id) ON DELETE cascade,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE cascade 
);
