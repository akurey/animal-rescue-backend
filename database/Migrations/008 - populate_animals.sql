-- the csv in \animal-rescue-backend\database was cleaned 
-- beforehand to delete duplicate values and other wrong values

INSERT INTO "AP_Conservation_Status" (name, abbreviation)
VALUES ('N/A', 'NA');

CREATE TEMPORARY TABLE "temp_animals"(
	scientific_name VARCHAR(100),
	common_name VARCHAR(100),
	classification VARCHAR(50)
);

-- load data from csv file to temp table (PSQL command)
\copy temp_animals(scientific_name, common_name, classification) FROM '-path_to_repo_goes_here_\animal-rescue-backend\database\AP_Animals.csv' DELIMITER ',' CSV HEADER;


INSERT INTO "AP_Animals" (name, scientific_name, conservation_status_id, classification_id)
SELECT common_name, scientific_name, CS.id, AC.id
FROM "temp_animals" TA
INNER JOIN "AP_Conservation_Status" CS ON CS.name = 'N/A'
INNER JOIN "AP_Animal_Classification" AC ON AC.name = TA.classification; 
