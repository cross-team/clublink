-- +migrate Up
CREATE TABLE user_url_relation
(
    -- TODO(issue#824): Add Primary Key Constraint to user_email and url_alias on migrate up
    -- TODO(issue#824): Set user_email to NOT NULL on migrate up
    user_email CHARACTER VARYING(254),
    url_id  CHARACTER VARYING  NOT NULL,
    FOREIGN KEY (user_email) REFERENCES "user" (email) ON UPDATE CASCADE,
    FOREIGN KEY (url_id) REFERENCES url (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +migrate Down
DROP TABLE user_url_relation;