CREATE OR REPLACE FUNCTION CAST_TO_INT(TEXT, INTEGER) RETURNS INTEGER
AS $$
BEGIN
    RETURN CAST($1 AS INTEGER);
    EXCEPTION
    	WHEN invalid_text_representation THEN
        	RETURN $2;
END;
$$ LANGUAGE plpgsql IMMUTABLE;

CREATE OR REPLACE FUNCTION AFN_GetAnimalReport(pReportId int)
RETURNS TABLE(ID BIGINT, IdAnimal BIGINT, AnimalName VARCHAR, ScientificName VARCHAR,
			  ConservationStatusName VARCHAR, Abbreviaton VARCHAR, ClassificationName VARCHAR, Fields JSONB)
      LANGUAGE 'plpgsql'
      AS $BODY$
BEGIN
RETURN QUERY

    WITH Directions AS (
        SELECT AD.id, JSONB_BUILD_OBJECT('Provincia', AP.name, 'Canton', AC.name, 'Distrito', D.name, 'Direccion', AD.exact_direction) AS value
        FROM "AP_Directions" AD
        INNER JOIN "AP_Districts" D ON D.id = AD.district_id
        INNER JOIN "AP_Cantons" AC ON AC.id = D.canton_id
        INNER JOIN "AP_Provinces" AP ON AP.id = AC.province_id
    )
    SELECT AAR.id ID, AA.id IdAnimal, AA.name AnimalName, AA.scientific_name ScientificName,
           ACS."name" ConservationStatusName, ACS.abbreviation Abbreviaton, AAC.name ClassificationName,
           JSONB_OBJECT_AGG(AF.name, COALESCE(D.value::TEXT, ARFV.value)) Fields
    FROM "AP_Animal_Reports" AAR
    INNER JOIN "AP_Animals" AA ON AAR.animal_id = AA.id
    INNER JOIN "AP_Conservation_Status" ACS ON AA.conservation_status_id = ACS.id
    INNER JOIN "AP_Animal_Classification" AAC ON AA.classification_id = AAC.id
    INNER JOIN "AP_Report_Field_Values" ARFV ON AAR.id = ARFV.report_id
    INNER JOIN "AP_Fields" AF ON AF.id = ARFV.field_id
    LEFT JOIN Directions D ON AF.name = 'Dirección' AND D.id = CAST_TO_INT(ARFV.value, 0)
    WHERE AAR.id = pReportId
    GROUP BY AAR.id, AA.id, AA.name, AA.scientific_name, ACS.name, ACS.abbreviation, AAC.name;

END;
$BODY$;

-- Usage:
-- SELECT public.AFN_GetAnimalReport(2);

CREATE OR REPLACE PROCEDURE ASP_GetAnimalReport(pReportId INT, INOUT RESULT refcursor)
LANGUAGE 'plpgsql'
AS $BODY$
BEGIN

OPEN RESULT FOR

    WITH Directions AS (
        SELECT AD.id, JSONB_BUILD_OBJECT('Provincia', AP.name, 'Canton', AC.name, 'Distrito', D.name, 'Direccion', AD.exact_direction) AS value
        FROM "AP_Directions" AD
        INNER JOIN "AP_Districts" D ON D.id = AD.district_id
        INNER JOIN "AP_Cantons" AC ON AC.id = D.canton_id
        INNER JOIN "AP_Provinces" AP ON AP.id = AC.province_id
    )
    SELECT AAR.id ID, AA.id IdAnimal, AA.name AnimalName, AA.scientific_name ScientificName,
           ACS."name" ConservationStatusName, ACS.abbreviation Abbreviaton, AAC.name ClassificationName,
           JSONB_OBJECT_AGG(AF.name, COALESCE(D.value::TEXT, ARFV.value)) Fields
    FROM "AP_Animal_Reports" AAR
    INNER JOIN "AP_Animals" AA ON AAR.animal_id = AA.id
    INNER JOIN "AP_Conservation_Status" ACS ON AA.conservation_status_id = ACS.id
    INNER JOIN "AP_Animal_Classification" AAC ON AA.classification_id = AAC.id
    INNER JOIN "AP_Report_Field_Values" ARFV ON AAR.id = ARFV.report_id
    INNER JOIN "AP_Fields" AF ON AF.id = ARFV.field_id
    LEFT JOIN Directions D ON AF.name = 'Dirección' AND D.id = CAST_TO_INT(ARFV.value, 0)
    WHERE AAR.id = pReportId
    GROUP BY AAR.id, AA.id, AA.name, AA.scientific_name, ACS.name, ACS.abbreviation, AAC.name;

END;
$BODY$;

-- Usage:
-- CALL public.ASP_GetAnimalReport(2,  'result');
-- FETCH ALL IN "result";