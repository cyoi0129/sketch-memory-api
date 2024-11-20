CREATE TABLE sketch_users (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  password text NOT NULL
);
CREATE TABLE sketch_authors (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  birth integer,
  user_id integer,
  CONSTRAINT author_user_fk FOREIGN KEY (user_id) REFERENCES sketch_users (id)
);
CREATE TABLE sketch_tags (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  user_id integer,
  CONSTRAINT tag_user_fk FOREIGN KEY (user_id) REFERENCES sketch_users (id)
);
CREATE TABLE sketch_items (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title text NOT NULL,
  image text NOT NULL,
  status text,
  author_id integer,
  tags integer[],
  date timestamp,
  user_id integer,
  CONSTRAINT item_user_fk FOREIGN KEY (user_id) REFERENCES sketch_users (id),
  CONSTRAINT item_author_fk FOREIGN KEY (author_id) REFERENCES sketch_authors (id)
);
CREATE TABLE sketch_reviews (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  reviewer text NOT NULL,
  score integer,
  comment text,
  item_id integer,
  CONSTRAINT review_item_fk FOREIGN KEY (item_id) REFERENCES sketch_items (id)
);