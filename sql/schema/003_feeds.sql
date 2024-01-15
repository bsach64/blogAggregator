-- +goose Up
CREATE TABLE "feeds" (
	"id" UUID NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	"name" TEXT NOT NULL,
	"url" TEXT NOT NULL,
	"user_id" UUID,
	PRIMARY KEY("id"),
	FOREIGN KEY("user_id") REFERENCES users("id") ON DELETE CASCADE
);

-- +goose Down
DROP TABLE "feeds";
