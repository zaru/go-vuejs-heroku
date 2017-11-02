-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE post (
    id SERIAL,
    title TEXT,
    body TEXT
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE post;
