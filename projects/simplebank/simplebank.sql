CREATE TABLE
    "Accounts" (
        "id" BIGSERIAL PRIMARY KEY,
        "owner" varchar NOT NULL,
        "balance" decimal NOT NULL,
        "currency" varchar NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now())
    );

CREATE TABLE
    "Entries" (
        "id" BIGSERIAL PRIMARY KEY,
        "account_id" bigint NOT NULL,
        "amount" decimal NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now())
    );

CREATE TABLE
    "Transfers" (
        "id" BIGSERIAL PRIMARY KEY,
        "from_account_id" bigint NOT NULL,
        "to_account_id" bigint NOT NULL,
        "amount" decimal NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now())
    );

CREATE INDEX ON "Accounts" ("owner");

CREATE INDEX ON "Entries" ("account_id");

CREATE INDEX ON "Transfers" ("from_account_id");

CREATE INDEX ON "Transfers" ("to_account_id");

CREATE INDEX ON "Transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "Entries"."amount" IS 'can be positive or negative';

COMMENT ON COLUMN "Transfers"."amount" IS 'must be positive';

ALTER TABLE "Entries"
ADD
    FOREIGN KEY ("account_id") REFERENCES "Accounts" ("id");

ALTER TABLE "Transfers"
ADD
    FOREIGN KEY ("from_account_id") REFERENCES "Accounts" ("id");

ALTER TABLE "Transfers"
ADD
    FOREIGN KEY ("to_account_id") REFERENCES "Accounts" ("id");