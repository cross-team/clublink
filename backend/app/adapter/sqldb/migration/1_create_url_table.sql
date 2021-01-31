-- +migrate Up
CREATE TABLE "Url"
(
    "id"          CHARACTER VARYING PRIMARY KEY,
    "alias"       CHARACTER VARYING,
    "originalUrl" TEXT,
    "expireAt"    TIMESTAMP WITH TIME ZONE,
    "createdAt"   TIMESTAMP WITH TIME ZONE,
    "updatedAt"   TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE "Url";