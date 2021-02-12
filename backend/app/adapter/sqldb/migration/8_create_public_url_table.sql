-- +migrate Up
CREATE TABLE public_url (
    id CHARACTER VARYING NOT NULL,
    FOREIGN KEY (id) REFERENCES url (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +migrate Down
DROP TABLE public_url;