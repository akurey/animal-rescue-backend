CREATE TABLE "AP_Users" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "first_name" character varying(50) NOT NULL,
  "last_name" character varying(50) NOT NULL,
  "username" character varying(100) NOT NULL,
  "email" character varying(200) NOT NULL,
  "password" character varying(500) NOT NULL,
  "is_enabled" bit DEFAULT '1' NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_Forms" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "shelter_id" bigint NOT NULL,
  "type_id" int NOT NULL,
  "name" character varying(100) NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_Fields" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "field_type_id" int NOT NULL,
  "name" character varying(100) NOT NULL,
  "is_required" bit NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_Field_Types" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(50) NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_Report_Field_Values" (
  "report_id" bigint NOT NULL,
  "field_id" bigint NOT NULL,
  "value" json NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL
);

CREATE TABLE "AP_Form_Fields" (
  "form_id" bigint NOT NULL,
  "field_id" bigint NOT NULL,
  "is_public" bit NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_Animal_Reports" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "form_id" bigint NOT NULL,
  "reporter_id" bigint NOT NULL,
  "animal_id" bigint NOT NULL,
  "is_public" bit NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_Form_Types" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "type" varchar(50) NOT NULL
);

CREATE TABLE "AP_Animals" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "name" character varying(100) NOT NULL,
  "scientific_name" character varying(200) NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_Field_Options" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "field_id" bigint NOT NULL,
  "option" _varchar NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_Shelters" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "name" character varying(100) NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_User_Shelters" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "user_id" bigint,
  "shelter_id" bigint,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_Permissions" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(100) NOT NULL,
  "description" varchar(500) NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_Role_Shelters" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(50) NOT NULL,
  "description" varchar(500) NOT NULL,
  "shelter_id" bigint NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_Permissions_Role" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "permission_id" bigint NOT NULL,
  "role_id" bigint NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

CREATE TABLE "AP_User_Role" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "user_id" bigint NOT NULL,
  "role_id" bigint NOT NULL,
  "created_at" timestamp DEFAULT NOW() NOT NULL,
  "updated_at" timestamp DEFAULT NOW() NOT NULL,
  "is_deleted" bit DEFAULT '0' NOT NULL
);

ALTER TABLE "AP_Forms" ADD FOREIGN KEY ("type_id") REFERENCES "AP_Form_Types" ("id");

ALTER TABLE "AP_Fields" ADD FOREIGN KEY ("field_type_id") REFERENCES "AP_Field_Types" ("id");

ALTER TABLE "AP_Report_Field_Values" ADD FOREIGN KEY ("report_id") REFERENCES "AP_Animal_Reports" ("id");

ALTER TABLE "AP_Report_Field_Values" ADD FOREIGN KEY ("field_id") REFERENCES "AP_Fields" ("id");

ALTER TABLE "AP_Form_Fields" ADD FOREIGN KEY ("form_id") REFERENCES "AP_Forms" ("id");

ALTER TABLE "AP_Form_Fields" ADD FOREIGN KEY ("field_id") REFERENCES "AP_Fields" ("id");

ALTER TABLE "AP_Animal_Reports" ADD FOREIGN KEY ("form_id") REFERENCES "AP_Forms" ("id");

ALTER TABLE "AP_Animal_Reports" ADD FOREIGN KEY ("reporter_id") REFERENCES "AP_Users" ("id");

ALTER TABLE "AP_Animal_Reports" ADD FOREIGN KEY ("animal_id") REFERENCES "AP_Animals" ("id");

ALTER TABLE "AP_Field_Options" ADD FOREIGN KEY ("field_id") REFERENCES "AP_Fields" ("id");

ALTER TABLE "AP_Forms" ADD FOREIGN KEY ("shelter_id") REFERENCES "AP_Shelters" ("id");

ALTER TABLE "AP_Users" ADD CONSTRAINT user_username_key UNIQUE ("username");

ALTER TABLE "AP_Users" ADD CONSTRAINT user_email_key UNIQUE ("email");

ALTER TABLE "AP_User_Shelters" ADD FOREIGN KEY ("shelter_id") REFERENCES "AP_Shelters" ("id");

ALTER TABLE "AP_User_Shelters" ADD FOREIGN KEY ("user_id") REFERENCES "AP_Users" ("id");

ALTER TABLE "AP_Role_Shelters" ADD FOREIGN KEY ("shelter_id") REFERENCES "AP_Shelters" ("id");

ALTER TABLE "AP_Permissions_Role" ADD FOREIGN KEY ("role_id") REFERENCES "AP_Role_Shelters" ("id");

ALTER TABLE "AP_Permissions_Role" ADD FOREIGN KEY ("permission_id") REFERENCES "AP_Permissions" ("id");

ALTER TABLE "AP_User_Role" ADD FOREIGN KEY ("user_id") REFERENCES "AP_User_Shelters" ("id");

ALTER TABLE "AP_User_Role" ADD FOREIGN KEY ("role_id") REFERENCES "AP_Role_Shelters" ("id");
