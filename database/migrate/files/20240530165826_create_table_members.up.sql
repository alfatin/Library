CREATE TABLE members (
  code varchar NOT NULL,
  name varchar NOT NULL,
  date_penalty_expired timestamp,
  PRIMARY KEY (code)
);