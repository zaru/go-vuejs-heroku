-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE post (
    id INTEGER PRIMARY KEY,
    title TEXT,
    body TEXT
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE post;
