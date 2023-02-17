CREATE TABLE IF NOT EXISTS "notif_statuses" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(20) NOT NULL CHECK ("name"=upper("name"))
);

CREATE TABLE IF NOT EXISTS "notifications" (
    "id" UUID PRIMARY KEY,
    "shipper_id" UUID not null,
    "content" TEXT NOT NULL ,
    "sender_id" UUID NOT NULL,
    "sender_type" VARCHAR(20) NOT NULL,
    "order_id" UUID,
    "receiver_type" VARCHAR(20),
    "title" TEXT,
    "to_all" BOOLEAN DEFAULT false,
    "created_at" TIMESTAMP NOT NULL DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS "notif_receivers" (
    "id" UUID PRIMARY KEY,
    "notif_id" UUID NOT NULL REFERENCES "notifications"("id"),
    "receiver_id" UUID NOT NULL,
    "receiver_type" VARCHAR(20) NOT NULL,
    "notif_status_id" UUID NOT NULL REFERENCES "notif_statuses"("id"),
    "send_at" TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS "notif_receivers_i1"  ON "notif_receivers" (
     "notif_id", "receiver_id", "receiver_type"
);

CREATE UNIQUE INDEX IF NOT EXISTS "notif_receivers_i1"  ON "notif_receivers" (
    "notif_status_id"
);

INSERT INTO "notif_statuses" VALUES('2cb46139-7c5f-4e0d-88f6-50fe79986fa6', 'NEW') ON CONFLICT DO NOTHING;
INSERT INTO "notif_statuses" VALUES('f9c31d41-7e91-42a1-859b-e0991357566c', 'SENT') ON CONFLICT DO NOTHING;
INSERT INTO "notif_statuses" VALUES('303d5102-7595-4ba4-9c67-0c5d969d6a3e', 'READ') ON CONFLICT DO NOTHING;
