create table if not exists messages (
id bigint not null AUTO_INCREMENT PRIMARY KEY,
email varchar(140) not null default "",
message text not null default "",
created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
version int not null default 1

)
