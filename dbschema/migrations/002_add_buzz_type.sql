-- +migrate Up
CREATE TYPE test_type AS ENUM ('testA', 'testB', 'testC', 'testD');
ALTER TABLE buzz_feeds ADD COLUMN buzz_type test_type;

-- +migrate Down
DROP TABLE IF EXISTS buzz_feeds;
DROP TYPE test_type;