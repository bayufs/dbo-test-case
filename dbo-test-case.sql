-- This script only contains the table creation statements and does not fully represent the table in database. It's still missing: indices, triggers. Do not use it as backup.

-- Squences
CREATE SEQUENCE IF NOT EXISTS authentications_id_seq

-- Table Definition
CREATE TABLE "public"."authentications" (
    "id" int4 NOT NULL DEFAULT nextval('authentications_id_seq'::regclass),
    "customer_id" int4,
    "username" varchar NOT NULL,
    "password" varchar NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    CONSTRAINT "authentications_customer_id_fkey" FOREIGN KEY ("customer_id") REFERENCES "public"."customers"("id"),
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in database. It's still missing: indices, triggers. Do not use it as backup.

-- Squences
CREATE SEQUENCE IF NOT EXISTS customers_id_seq

-- Table Definition
CREATE TABLE "public"."customers" (
    "id" int4 NOT NULL DEFAULT nextval('customers_id_seq'::regclass),
    "first_name" varchar NOT NULL,
    "last_name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "phone" varchar,
    "address" varchar,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in database. It's still missing: indices, triggers. Do not use it as backup.

-- Squences
CREATE SEQUENCE IF NOT EXISTS order_items_id_seq

-- Table Definition
CREATE TABLE "public"."order_items" (
    "id" int4 NOT NULL DEFAULT nextval('order_items_id_seq'::regclass),
    "order_id" int4,
    "product_id" int4,
    "quantity" int4 NOT NULL,
    "unit_price" numeric NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    CONSTRAINT "order_items_order_id_fkey" FOREIGN KEY ("order_id") REFERENCES "public"."orders"("id"),
    CONSTRAINT "order_items_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products"("id"),
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in database. It's still missing: indices, triggers. Do not use it as backup.

-- Squences
CREATE SEQUENCE IF NOT EXISTS orders_id_seq

-- Table Definition
CREATE TABLE "public"."orders" (
    "id" int4 NOT NULL DEFAULT nextval('orders_id_seq'::regclass),
    "order_date" timestamptz NOT NULL,
    "customer_id" int4,
    "total_amount" numeric NOT NULL,
    "status" varchar,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    CONSTRAINT "orders_customer_id_fkey" FOREIGN KEY ("customer_id") REFERENCES "public"."customers"("id"),
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in database. It's still missing: indices, triggers. Do not use it as backup.

-- Squences
CREATE SEQUENCE IF NOT EXISTS products_id_seq

-- Table Definition
CREATE TABLE "public"."products" (
    "id" int4 NOT NULL DEFAULT nextval('products_id_seq'::regclass),
    "product_name" varchar NOT NULL,
    "description" text,
    "price" numeric NOT NULL,
    "stock" int4 NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);









INSERT INTO "public"."products" ("id", "product_name", "description", "price", "stock", "created_at", "updated_at", "deleted_at") VALUES
(1, 'Wireless Mouse', 'A high-precision wireless mouse with ergonomic design', 259000.00, 150, '2024-05-23 22:05:46.737459+07', '2024-05-23 22:05:46.737459+07', NULL);
INSERT INTO "public"."products" ("id", "product_name", "description", "price", "stock", "created_at", "updated_at", "deleted_at") VALUES
(2, 'Mechanical Keyboard', 'A durable mechanical keyboard with RGB backlighting', 899000.00, 75, '2024-05-23 22:05:46.737459+07', '2024-05-23 22:05:46.737459+07', NULL);
INSERT INTO "public"."products" ("id", "product_name", "description", "price", "stock", "created_at", "updated_at", "deleted_at") VALUES
(3, 'Noise Cancelling Headphones', 'Over-ear headphones with active noise cancellation', 1299000.00, 50, '2024-05-23 22:05:46.737459+07', '2024-05-23 22:05:46.737459+07', NULL);
INSERT INTO "public"."products" ("id", "product_name", "description", "price", "stock", "created_at", "updated_at", "deleted_at") VALUES
(4, '4K Monitor', 'A 27-inch 4K UHD monitor with HDR support', 3499000.00, 30, '2024-05-23 22:05:46.737459+07', '2024-05-23 22:05:46.737459+07', NULL),
(5, 'Portable SSD', 'A 1TB portable SSD with high-speed data transfer', 1099000.00, 200, '2024-05-23 22:05:46.737459+07', '2024-05-23 22:05:46.737459+07', NULL);
