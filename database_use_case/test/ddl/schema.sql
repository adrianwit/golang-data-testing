DROP TABLE IF EXISTS books;

CREATE TABLE books (
  id        INT PRIMARY KEY,
  isbn      VARCHAR(255) NOT NULL,
  title     VARCHAR(255) NOT NULL,
  summary   VARCHAR(255),
  author_id INT          NOT NULL
);


DROP TABLE IF EXISTS authors;

CREATE TABLE authors (
  id        INT PRIMARY KEY,
  name      VARCHAR(255) NOT NULL
);

