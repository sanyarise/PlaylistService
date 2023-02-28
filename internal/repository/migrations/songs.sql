CREATE TABLE IF NOT EXISTS songs
(
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    title varchar(250),
    duration INTEGER
)