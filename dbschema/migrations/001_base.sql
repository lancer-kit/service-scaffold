-- +migrate Up
CREATE TABLE IF NOT EXISTS buzz_feeds (
  id          SERIAL PRIMARY KEY,
  name        VARCHAR(255),
  description TEXT  NOT NULL DEFAULT '',
  details     JSONB NOT NULL DEFAULT '{}'
);

-- +migrate Down
DROP TABLE IF EXISTS buzz_feeds;
