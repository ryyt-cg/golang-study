CREATE TABLE author (
    id   integer PRIMARY KEY AUTOINCREMENT,
    name text    NOT NULL,
    bio  text
);

CREATE TABLE book (
    id         integer PRIMARY KEY AUTOINCREMENT,
    title      text    NOT NULL,
    author_id  integer NOT NULL,
    published  date    NOT NULL,
    FOREIGN KEY (author_id) REFERENCES author (id)
);