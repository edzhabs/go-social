ALTER TABLE 
    POSTS
ADD 
    COLUMN TAGS VARCHAR(100) [];

ALTER TABLE 
    POSTS
ADD 
    COLUMN updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW();