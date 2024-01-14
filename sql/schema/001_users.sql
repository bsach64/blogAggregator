-- +goose Up
CREATE TABLE "users" (
	"id" UUID, 
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	"name" VARCHAR(32) NOT NULL,
	PRIMARY KEY("id")
);

-- +goose Down
DROP TABLE "users";
