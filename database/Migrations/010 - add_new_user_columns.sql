ALTER TABLE "AP_Users"
  ADD "token" character varying(500) NULL,
  ADD "refresh_token" character varying(500) NULL;

ALTER TABLE "AP_Users" 
ALTER COLUMN "sinac_registry" DROP NOT NULL;

