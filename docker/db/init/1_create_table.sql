CREATE TABLE sketch_authors (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  birth integer
);
CREATE TABLE sketch_tags (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL
);
CREATE TABLE sketch_items (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title text NOT NULL,
  image text NOT NULL,
  status text,
  author_id integer,
  tag_ids integer [],
  date timestamp,
  user_id text
);
CREATE TABLE sketch_reviews (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  reviewer text NOT NULL,
  score integer,
  comment text,
  item text
);