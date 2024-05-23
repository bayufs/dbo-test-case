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

INSERT INTO "public"."authentications" ("id", "customer_id", "username", "password", "created_at", "updated_at", "deleted_at") VALUES
(1, 1, 'johndoe', '$2a$10$s/09hvCPPcICOBhEmVVGIeOFEt9i8c8oIykhDbH0y872j3jnwpzny', '2024-05-23 14:41:23.414533+07', '2024-05-23 14:56:59.875655+07', NULL);
INSERT INTO "public"."authentications" ("id", "customer_id", "username", "password", "created_at", "updated_at", "deleted_at") VALUES
(3, 3, 'bayli', '$2a$10$DY3dVHKiWWuRy9ZydNAkjelslY1XyaxYQDXczj2mBUICbzve/qkUq', '2024-05-23 15:20:33.463339+07', '2024-05-23 15:20:33.463339+07', NULL);
INSERT INTO "public"."authentications" ("id", "customer_id", "username", "password", "created_at", "updated_at", "deleted_at") VALUES
(4, 4, 'libay', '$2a$10$jjUETQZWdiyZj57dfhAHpuS8iCOFSgrQm8uwYfdzUk2r5UPeV2v/G', '2024-05-23 15:20:59.315137+07', '2024-05-23 15:20:59.315137+07', NULL);
INSERT INTO "public"."authentications" ("id", "customer_id", "username", "password", "created_at", "updated_at", "deleted_at") VALUES
(5, 5, 'kopet', '$2a$10$696d8rmASNR3mK/WqC7sYu0CxAupxwjjyEdlY7EnsDmOTt28biJOi', '2024-05-23 15:21:29.097428+07', '2024-05-23 15:21:29.097428+07', NULL),
(6, 6, 'kopet', '$2a$10$SVEQitJSrrqyyeTVoPIbpugUjAM/j2qhoXlq18BqlE4pIlByHVBAm', '2024-05-23 15:22:36.31727+07', '2024-05-23 15:22:36.31727+07', NULL),
(7, 7, 'kopet', '$2a$10$B5EysDQl.sVRhZMGLe0g3uXthAeCRaothpoEw.tD0WPClKSd02piW', '2024-05-23 15:22:37.316723+07', '2024-05-23 15:22:37.316723+07', NULL),
(8, 8, 'kopet', '$2a$10$WtDDgFp9hQQ9B6fbrfw3wu3z7ls7DQONsPRIQkIskedoXnWn/okcm', '2024-05-23 15:22:38.405485+07', '2024-05-23 15:22:38.405485+07', NULL),
(9, 9, 'kopet', '$2a$10$SRedf/BLUX/NqzdqQYbxn.NAcq6eIpX7KfJiXr3u3f7E9SSsicZfW', '2024-05-23 15:22:39.264746+07', '2024-05-23 15:22:39.264746+07', NULL),
(11, 11, 'kopet', '$2a$10$RnXLZyp2dTuuzmf79XYC9e3Aqlm9yC6BHuJSBWn/E8EdPcVpynCoW', '2024-05-23 15:22:41.03656+07', '2024-05-23 15:22:41.03656+07', NULL),
(12, 12, 'kopet', '$2a$10$GKzvEsHb7urHbqC1f5DYOu2hCh6hBWET1xskH1EMKZ5Z7DGyyeQFm', '2024-05-23 15:22:41.929147+07', '2024-05-23 15:22:41.929147+07', NULL),
(13, 13, 'kopet', '$2a$10$NYGdh.alL5mpxZt.9XUPouvWIeUzaZpBvrGpbvcFKMnXv7FA8uot.', '2024-05-23 15:22:45.311863+07', '2024-05-23 15:22:45.311863+07', NULL);

INSERT INTO "public"."customers" ("id", "first_name", "last_name", "email", "phone", "address", "created_at", "updated_at", "deleted_at") VALUES
(1, 'john', 'Cok', 'john.doe@example.com', '1234567890', '123 Main St', '2024-05-23 14:41:23.395036+07', '2024-05-23 14:56:59.782053+07', NULL);
INSERT INTO "public"."customers" ("id", "first_name", "last_name", "email", "phone", "address", "created_at", "updated_at", "deleted_at") VALUES
(3, 'bayu', 'fajar', 'bayu@gmail.com', '123321123', '123 Main St', '2024-05-23 15:20:33.453799+07', '2024-05-23 15:20:33.453799+07', NULL);
INSERT INTO "public"."customers" ("id", "first_name", "last_name", "email", "phone", "address", "created_at", "updated_at", "deleted_at") VALUES
(4, 'lia', 'awaliyah', 'lia@gmail.com', '123321123', '123 Main St', '2024-05-23 15:20:59.309179+07', '2024-05-23 15:20:59.309179+07', NULL);
INSERT INTO "public"."customers" ("id", "first_name", "last_name", "email", "phone", "address", "created_at", "updated_at", "deleted_at") VALUES
(5, 'jajang', 'kopet', 'kopet@gmail.com', '123321123', '123 Main St', '2024-05-23 15:21:29.086882+07', '2024-05-23 15:21:29.086882+07', NULL),
(6, 'jajang', 'kopet', 'kopet@gmail.com', '123321123', '123 Main St', '2024-05-23 15:22:36.300009+07', '2024-05-23 15:22:36.300009+07', NULL),
(7, 'jajang', 'kopet', 'kopet@gmail.com', '123321123', '123 Main St', '2024-05-23 15:22:37.31385+07', '2024-05-23 15:22:37.31385+07', NULL),
(8, 'jajang', 'kopet', 'kopet@gmail.com', '123321123', '123 Main St', '2024-05-23 15:22:38.401319+07', '2024-05-23 15:22:38.401319+07', NULL),
(9, 'jajang', 'kopet', 'kopet@gmail.com', '123321123', '123 Main St', '2024-05-23 15:22:39.260732+07', '2024-05-23 15:22:39.260732+07', NULL),
(11, 'jajang', 'kopet', 'kopet@gmail.com', '123321123', '123 Main St', '2024-05-23 15:22:41.033044+07', '2024-05-23 15:22:41.033044+07', NULL),
(12, 'jajang', 'kopet', 'kopet@gmail.com', '123321123', '123 Main St', '2024-05-23 15:22:41.927142+07', '2024-05-23 15:22:41.927142+07', NULL),
(13, 'jajang', 'kopet', 'kopet@gmail.com', '123321123', '123 Main St', '2024-05-23 15:22:45.307323+07', '2024-05-23 15:22:45.307323+07', NULL);

INSERT INTO "public"."order_items" ("id", "order_id", "product_id", "quantity", "unit_price", "created_at", "updated_at", "deleted_at") VALUES
(2, 2, 1, 2, 10000.00, '2024-05-23 22:57:54.278977+07', '2024-05-23 22:57:54.278977+07', NULL);


INSERT INTO "public"."orders" ("id", "order_date", "customer_id", "total_amount", "status", "created_at", "updated_at", "deleted_at") VALUES
(2, '2024-05-23 22:57:54.252698+07', 1, 100.50, 'pending', '2024-05-23 22:57:54.252698+07', '2024-05-23 22:57:54.252698+07', NULL);


INSERT INTO "public"."products" ("id", "product_name", "description", "price", "stock", "created_at", "updated_at", "deleted_at") VALUES
(1, 'Wireless Mouse', 'A high-precision wireless mouse with ergonomic design', 259000.00, 150, '2024-05-23 22:05:46.737459+07', '2024-05-23 22:05:46.737459+07', NULL);
INSERT INTO "public"."products" ("id", "product_name", "description", "price", "stock", "created_at", "updated_at", "deleted_at") VALUES
(2, 'Mechanical Keyboard', 'A durable mechanical keyboard with RGB backlighting', 899000.00, 75, '2024-05-23 22:05:46.737459+07', '2024-05-23 22:05:46.737459+07', NULL);
INSERT INTO "public"."products" ("id", "product_name", "description", "price", "stock", "created_at", "updated_at", "deleted_at") VALUES
(3, 'Noise Cancelling Headphones', 'Over-ear headphones with active noise cancellation', 1299000.00, 50, '2024-05-23 22:05:46.737459+07', '2024-05-23 22:05:46.737459+07', NULL);
INSERT INTO "public"."products" ("id", "product_name", "description", "price", "stock", "created_at", "updated_at", "deleted_at") VALUES
(4, '4K Monitor', 'A 27-inch 4K UHD monitor with HDR support', 3499000.00, 30, '2024-05-23 22:05:46.737459+07', '2024-05-23 22:05:46.737459+07', NULL),
(5, 'Portable SSD', 'A 1TB portable SSD with high-speed data transfer', 1099000.00, 200, '2024-05-23 22:05:46.737459+07', '2024-05-23 22:05:46.737459+07', NULL);
