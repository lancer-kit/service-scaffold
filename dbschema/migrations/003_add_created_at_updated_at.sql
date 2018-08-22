-- +migrate Up
ALTER TABLE buzz_feeds ADD COLUMN created_at integer;
ALTER TABLE buzz_feeds ADD COLUMN updated_at integer;

-- +migrate Down