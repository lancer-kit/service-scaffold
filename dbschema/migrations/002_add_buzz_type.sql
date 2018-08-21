-- +migrate Up
CREATE TYPE type_enum AS ENUM ('ExampleTypeA', 'ExampleTypeB', 'ExampleTypeA');
ALTER TABLE buzz_feed ADD COLUMN IF NOT EXISTS buzz_type type_enum;

-- +migrate Down
ALTER TABLE buzz_feed DROP COLUMN IF EXISTS buzz_type;