-- +migrate Up
CREATE TYPE test_type AS ENUM ('testA', 'testB', 'testC', 'testD');
ALTER TABLE buzz_feeds ADD COLUMN buzz_type test_type;
ALTER TABLE buzz_feeds ADD COLUMN created_at integer;
ALTER TABLE buzz_feeds ADD COLUMN updated_at integer;

-- +migrate Down
ALTER TABLE buzz_feeds DROP COLUMN IF EXISTS buzz_type;
ALTER TABLE buzz_feeds DROP COLUMN IF EXISTS created_at;
ALTER TABLE buzz_feeds DROP COLUMN IF EXISTS updated_at;
DROP TYPE test_type;