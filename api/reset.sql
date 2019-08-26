DROP TABLE tasks;

CREATE TABLE tasks (
  id serial NOT NULL,
  title varchar NOT NULL,
  description varchar,
  priority integer NOT NULL DEFAULT 0,
  done boolean NOT NULL DEFAULT false
);
