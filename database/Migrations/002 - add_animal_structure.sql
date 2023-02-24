ALTER TABLE "AP_Animals"
ADD COLUMN conservation_status_id INT NOT NULL,
ADD COLUMN classification_id INT NOT NULL;

ALTER TABLE "AP_Animal_Reports"
ADD COLUMN is_approved BIT NOT NULL;

CREATE TABLE "AP_Conservation_Status" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(30) NOT NULL,
  "abbreviation" varchar(2),
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "is_deleted" bit NOT NULL
);

CREATE TABLE "AP_Animal_Classification" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(30) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "is_deleted" bit NOT NULL
);

ALTER TABLE "AP_Animals" ADD FOREIGN KEY ("conservation_status_id") REFERENCES "AP_Conservation_Status" ("id");

ALTER TABLE "AP_Animals" ADD FOREIGN KEY ("classification_id") REFERENCES "AP_Animal_Classification" ("id");