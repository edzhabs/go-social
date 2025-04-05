CREATE TABLE IF NOT EXISTS user_invitations (
    token bytea NOT NULL,
    user_id bigint NOT NULL,
    expiry timestamp(0) with time zone NOT NULL,

    PRIMARY KEY (token, user_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);