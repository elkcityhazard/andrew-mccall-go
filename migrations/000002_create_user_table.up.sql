create table passwords (
    id bigint primary key AUTO_INCREMENT,
    hash varchar(255) not null,
    created_at datetime not null default CURRENT_TIMESTAMP,
    updated_at datetime not null default CURRENT_TIMESTAMP,
    is_active boolean not null default false,
    version bigint not null default 1
);

insert into 
passwords
(id, hash, is_active,version)
values(1,"lalalala",true,1);


create table users (
    id bigint primary key AUTO_INCREMENT,
    email varchar(140) unique not null,
    username varchar(140) unique not null,
    password bigint  not null,
    created_at datetime not null default CURRENT_TIMESTAMP,
    updated_at datetime not null default CURRENT_TIMESTAMP,
    is_active boolean not null default false,
    role ENUM('super_admin', 'admin','user','commenter') not null default 'user',
    version bigint not null default 1,
    CONSTRAINT fk_password foreign key (password) references passwords(id) on delete cascade
);


create index users_pw_idx on users (email,password);

insert into
users
(id, email, username, password, is_active, role, version)
VALUES(1, "fake@fake.com","fakeuser",1, true, "admin",1);
