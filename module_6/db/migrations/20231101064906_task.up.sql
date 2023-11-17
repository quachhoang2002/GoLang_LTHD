-- db/migrations/000001_create_tasks_table.up.sql

CREATE TABLE tasks (
    id serial PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL
);
