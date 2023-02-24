ALTER TABLE "AP_Conservation_Status" 
ALTER COLUMN "created_at" SET DEFAULT NOW(),
ALTER COLUMN "updated_at" SET DEFAULT NOW(),
ALTER COLUMN "is_deleted" SET DEFAULT '0';

ALTER TABLE "AP_Animal_Classification" 
ALTER COLUMN "created_at" SET DEFAULT NOW(),
ALTER COLUMN "updated_at" SET DEFAULT NOW(),
ALTER COLUMN "is_deleted" SET DEFAULT '0';