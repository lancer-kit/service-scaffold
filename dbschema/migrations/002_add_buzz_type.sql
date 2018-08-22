-- +migrate Up
CREATE TYPE test_type AS ENUM ('testA', 'testB', 'testC', 'testD');
ALTER TABLE buzz_feeds ADD COLUMN buzz_type test_type;

-- +migrate Down
ALTER TABLE buzz_feeds DROP COLUMN IF EXISTS buzz_type;
DROP TYPE test_type;