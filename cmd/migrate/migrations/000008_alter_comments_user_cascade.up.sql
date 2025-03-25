ALTER TABLE comments
    DROP CONSTRAINT comments_user_id_fkey,
    ADD CONSTRAINT comments_user_id_fkey
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;