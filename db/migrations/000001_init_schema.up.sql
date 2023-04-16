CREATE TABLE "client" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "token" varchar NOT NULL
);

CREATE TABLE "order" (
  "id" bigserial PRIMARY KEY,
  "client_id" bigint NOT NULL,
  "origin_address" varchar NOT NULL,
  "origin_postal_code" varchar NOT NULL,
  "origin_ext_num" varchar NOT NULL,
  "origin_int_num" varchar,
  "origin_city" varchar NOT NULL,
  "destination_address" varchar NOT NULL,
  "destination_postal_code" varchar NOT NULL,
  "destination_ext_num" varchar NOT NULL,
  "destination_int_num" varchar,
  "destination_city" varchar NOT NULL,
  "product_quantity" int NOT NULL,
  "total_weight" float NOT NULL,
  "package_size" varchar NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "was_refunded" boolean NOT NULL
);

CREATE TABLE "auth" (
  "id" bigserial PRIMARY KEY,
  "client_id" bigint NOT NULL,
  "token" varchar NOT NULL
);

CREATE INDEX ON "client" ("email");

CREATE INDEX ON "order" ("client_id");

CREATE INDEX ON "auth" ("client_id");

ALTER TABLE "order" ADD FOREIGN KEY ("client_id") REFERENCES "client" ("id");

ALTER TABLE "auth" ADD FOREIGN KEY ("client_id") REFERENCES "client" ("id");
