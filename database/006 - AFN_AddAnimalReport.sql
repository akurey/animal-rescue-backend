CREATE OR REPLACE FUNCTION AFN_AddAnimalReport(
	pAnimalId BIGINT, 
	pReporterId BIGINT,
	pFormId BIGINT,
	pFieldValues JSON
) 
RETURNS setof record
LANGUAGE 'plpgsql' 
AS $BODY$ 
DECLARE
	report_id BIGINT;
BEGIN
	INSERT INTO 
		"AP_Animal_Reports" (form_id, reporter_id, animal_id, is_public, is_approved)
	VALUES 
		(pFormId, pReporterId, pAnimalId, '0', '0')
	RETURNING "AP_Animal_Reports".id INTO report_id;

	WITH FV AS
	(
		SELECT 
			FI.id field_id, value::json#>>'{}' field_value
		FROM 
			json_each(pFieldValues) FV
		INNER JOIN 
			"AP_Fields" FI ON FV.key = FI.name
	)
	INSERT INTO 
		"AP_Report_Field_Values" (report_id, field_id, value)
	SELECT 
		report_id, field_id, field_value
	FROM FV;

	RETURN QUERY SELECT AR.ID, to_char(AR.created_at, 'DD/MM/YYYY'), AR.is_approved, AN.name
	FROM "AP_Animal_Reports" AR
	INNER JOIN "AP_Animals" AN ON AR.animal_id = AN.id
	WHERE AR.id = report_id;

END;
$BODY$;
