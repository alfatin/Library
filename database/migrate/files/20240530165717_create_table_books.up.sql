CREATE TABLE books (
  code varchar NOT NULL,
  title varchar NOT NULL,
  author varchar NOT NULL,
  stock int NOT NULL,
  borrowed_by varchar,
  date_borrowed timestamp, 
  PRIMARY KEY (code)
);