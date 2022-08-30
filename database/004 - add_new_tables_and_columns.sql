CREATE TABLE "AP_Provinces" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "name" character varying(50) NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_Cantons" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "name" character varying(50) NOT NULL,
  "province_id" int NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_Districts" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "name" character varying(50) NOT NULL,
  "canton_id" int NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_Directions" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "exact_direction" text NOT NULL,
  "district_id" int NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

ALTER TABLE "AP_Cantons" ADD FOREIGN KEY ("province_id") REFERENCES "AP_Provinces" ("id");

ALTER TABLE "AP_Districts" ADD FOREIGN KEY ("canton_id") REFERENCES "AP_Cantons" ("id");

ALTER TABLE "AP_Directions" ADD FOREIGN KEY ("district_id") REFERENCES "AP_Districts" ("id");

CREATE TABLE "AP_Management_Category" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "name" character varying(50) NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

ALTER TABLE "AP_Users"
  ADD "identification" character varying(50) NOT NULL,
  ADD "collegiate_number" character varying(50) NULL,
  ADD "sinac_registry" character varying(50) NOT NULL;

ALTER TABLE "AP_Shelters"
  ADD "trade_name" character varying(100) NULL,
  ADD "management_category_id" int NULL,
  ADD "sinac_resolution_number" character varying(20) NULL,
  ADD "direction_id" int NULL,
  ADD "phone" character varying(20) NULL,
  ADD "owner" int NULL,
  ADD "regent_biologist" int NULL,
  ADD "regent_vet" int NULL


ALTER TABLE "AP_Shelters" ADD FOREIGN KEY ("management_category_id") REFERENCES "AP_Management_Category" ("id");

ALTER TABLE "AP_Shelters" ADD FOREIGN KEY ("direction_id") REFERENCES "AP_Directions" ("id");

ALTER TABLE "AP_Shelters" ADD FOREIGN KEY ("owner") REFERENCES "AP_Users" ("id");

ALTER TABLE "AP_Shelters" ADD FOREIGN KEY ("regent_biologist") REFERENCES "AP_Users" ("id");

ALTER TABLE "AP_Shelters" ADD FOREIGN KEY ("regent_vet") REFERENCES "AP_Users" ("id");

ALTER TABLE "AP_Shelters" 
  ALTER COLUMN "trade_name" SET NOT NULL,
  ALTER COLUMN "management_category_id" SET NOT NULL,
  ALTER COLUMN "sinac_resolution_number" SET NOT NULL,
  ALTER COLUMN "direction_id" SET NOT NULL,
  ALTER COLUMN "phone" SET NOT NULL,
  ALTER COLUMN "owner" SET NOT NULL,
  ALTER COLUMN "regent_biologist" SET NOT NULL,
  ALTER COLUMN "regent_vet" SET NOT NULL