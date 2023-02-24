CREATE OR REPLACE FUNCTION AFN_GetAnimalReport(pReportId int)
RETURNS TABLE(
    id BIGINT,
    id_animal BIGINT,
    animal_name VARCHAR,
    scientific_name VARCHAR,
    conservation_status_name VARCHAR,
    abbreviation VARCHAR,
    classification_name VARCHAR,
    fields JSONB
)
LANGUAGE 'plpgsql'
AS $BODY$
BEGIN
RETURN QUERY

    WITH Directions AS (
        SELECT D.id, JSONB_BUILD_OBJECT('Provincia', AP.name, 'Canton', AC.name, 'Distrito', D.name) AS value
        FROM "AP_Districts" D
        INNER JOIN "AP_Cantons" AC ON AC.id = D.canton_id
        INNER JOIN "AP_Provinces" AP ON AP.id = AC.province_id
    )
    SELECT AAR.id ID, AA.id IdAnimal, AA.name AnimalName, AA.scientific_name ScientificName,
        ACS."name" ConservationStatusName, ACS.abbreviation Abbreviaton, AAC.name ClassificationName,
        JSONB_OBJECT_AGG(AF.name, COALESCE((D.value || JSONB_BUILD_OBJECT('Exacta', SPLIT_PART(ARFV.value, '▽', 1)))::TEXT, ARFV.value)) Fields
    FROM "AP_Animal_Reports" AAR
         INNER JOIN "AP_Animals" AA ON AAR.animal_id = AA.id
         INNER JOIN "AP_Conservation_Status" ACS ON AA.conservation_status_id = ACS.id
         INNER JOIN "AP_Animal_Classification" AAC ON AA.classification_id = AAC.id
         INNER JOIN "AP_Report_Field_Values" ARFV ON AAR.id = ARFV.report_id
         INNER JOIN "AP_Fields" AF ON AF.id = ARFV.field_id
         INNER JOIN "AP_Field_Types" AFT ON AF.field_type_id = AFT.id
         LEFT JOIN Directions D ON AFT.name = 'Address' AND D.id = CAST_TO_INT(SPLIT_PART(ARFV.value, '▽', 2), 0)
    WHERE AAR.id = pReportId
    GROUP BY AAR.id, AA.id, AA.name, AA.scientific_name, ACS.name, ACS.abbreviation, AAC.name;

END;
$BODY$;
