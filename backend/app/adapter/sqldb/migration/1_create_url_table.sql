-- +migrate Up
CREATE TABLE "Url"
(
    "id"          CHARACTER VARYING(10) PRIMARY KEY,
    "alias"       CHARACTER VARYING,
    "originalUrl" TEXT,
    "room"        CHARACTER VARYING(255),
    "expireAt"    TIMESTAMP WITH TIME ZONE,
    "createdAt"   TIMESTAMP WITH TIME ZONE,
    "updatedAt"   TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE "Url";