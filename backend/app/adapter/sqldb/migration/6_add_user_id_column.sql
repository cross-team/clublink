-- +migrate Up
ALTER TABLE "user" ADD id CHARACTER VARYING(10);
ALTER TABLE user_url_relation ADD user_id CHARACTER VARYING(10);

-- +migrate Down
ALTER TABLE user_url_relation DROP user_id;
ALTER TABLE "user" DROP id;