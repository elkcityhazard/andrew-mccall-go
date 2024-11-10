CREATE TABLE IF NOT EXISTS activation_tokens (
    id bigint not null AUTO_INCREMENT PRIMARY KEY,
    hash LONGBLOB NOT NULL,
    user_id bigint not null,
    expiry DATETIME NOT NULL,
    scope text NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
    
);
