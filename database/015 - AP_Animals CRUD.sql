-- CRUD for animal species (not report)

CREATE OR REPLACE FUNCTION AFN_AddAnimal(pName VARCHAR, pScientificName VARCHAR, 
                                         pConservationStatusId INTEGER, pClassificationId INTEGER) 
RETURNS TABLE(animal_id BIGINT, animal_name VARCHAR (100), scientificName VARCHAR (200),
              conservation_status VARCHAR (30), conservation_abbreviation VARCHAR (2), classification_name VARCHAR (30))
LANGUAGE 'plpgsql' 
AS $BODY$ 
DECLARE 
	animal_id BIGINT;
BEGIN
	INSERT INTO 
		"AP_Animals" (name, scientific_name, conservation_status_id, classification_id)
	VALUES 
		(pName, pScientificName, pConservationStatusId, pClassificationId)
	RETURNING "AP_Animals".id INTO animal_id;

	RETURN QUERY 
	SELECT AN.id animal_id, AN.name animal_name, scientific_name scientific_name, 
         CS.name conservation_status, CS.abbreviation, AC.name classification
	FROM "AP_Animals" AN
	INNER JOIN "AP_Conservation_Status" CS ON CS.id = AN.conservation_status_id
	INNER JOIN "AP_Animal_Classification" AC ON AC.id = AN.classification_id
	WHERE AN.id = animal_id;
END;
$BODY$;

CREATE OR REPLACE FUNCTION AFN_UpdateAnimal(pAnimalId BIGINT, pName VARCHAR(100), pScientificName VARCHAR (200),
                                            pConservationStatusId INTEGER, pClassificationId INTEGER) 
RETURNS TABLE(id BIGINT, name VARCHAR (100), scientific_name VARCHAR (200), conservation_status VARCHAR (30),
              conservation_abbreviation VARCHAR (2), classification_name VARCHAR (30))
LANGUAGE 'plpgsql' 
AS $BODY$ 
BEGIN
	UPDATE "AP_Animals" AN
	SET name = pName, scientific_name = pScientificName, conservation_status_id = pConservationStatusId, 
             classification_id = pClassificationId
	WHERE AN.id = pAnimalId;
		
	RETURN QUERY 
	SELECT AN.id, AN.name, AN.scientific_name, CS.name conservation_status, 
         CS.abbreviation, AC.name classification
	FROM "AP_Animals" AN
	INNER JOIN "AP_Conservation_Status" CS ON CS.id = AN.conservation_status_id
	INNER JOIN "AP_Animal_Classification" AC ON AC.id = AN.classification_id
	WHERE AN.id = pAnimalId;
END;
$BODY$;

CREATE OR REPLACE FUNCTION AFN_DeleteAnimal(pAnimalId BIGINT) 
RETURNS TABLE(id BIGINT, name VARCHAR (100), scientific_name VARCHAR (200),
	            conservation_status VARCHAR (30), conservation_abbreviation VARCHAR (2),
	            classification_name VARCHAR (30))
LANGUAGE 'plpgsql' 
AS $BODY$ 
BEGIN
	UPDATE "AP_Animals" AN SET is_deleted = '1' WHERE AN.id = pAnimalId;

	RETURN QUERY 
	SELECT AN.id, AN.name, AN.scientific_name, 
	       CS.name conservation_status, CS.abbreviation, AC.name classification
	FROM "AP_Animals" AN
	INNER JOIN "AP_Conservation_Status" CS ON CS.id = AN.conservation_status_id
	INNER JOIN "AP_Animal_Classification" AC ON AC.id = AN.classification_id
	WHERE AN.id = pAnimalId;
END;
$BODY$;
