CREATE TABLE IF NOT EXISTS cards(
  id SERIAL PRIMARY KEY NOT NULL,
  source_member_id VARCHAR(50) NOT NULL,
  target_member_id VARCHAR(50) NOT NULL,
  point SMALLINT NOT NULL,
  message VARCHAR(400) NOT NULL
);
