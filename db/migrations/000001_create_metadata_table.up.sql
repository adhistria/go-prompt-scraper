CREATE TABLE IF NOT EXISTS metadatas(
  id BIGSERIAL PRIMARY KEY,
  num_of_links int,
  site VARCHAR(255),
  num_of_images int,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);