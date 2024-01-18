-- +goose Up
CREATE TABLE "feed_follows" (
	"id" UUID NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	"user_id" UUID NOT NULL,
	"feed_id" UUID NOT NULL,
	PRIMARY KEY("id"),
	FOREIGN KEY("user_id") REFERENCES users("id"),
	FOREIGN KEY("feed_id") REFERENCES feeds("id")
);

-- +goose Down
DROP TABLE "feed_follows";
