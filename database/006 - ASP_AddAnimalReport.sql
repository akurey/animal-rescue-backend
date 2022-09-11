CREATE OR REPLACE PROCEDURE ASP_AddAnimalReport(
	_animal_id BIGINT, 
	_reporter_id BIGINT,
	_form_id BIGINT,
	_field_values JSON,
	INOUT result refcursor
)
LANGUAGE 'plpgsql'
AS $BODY$
DECLARE
	report_id BIGINT;
BEGIN
	
	INSERT INTO 
		"AP_Animal_Reports" (form_id, reporter_id, animal_id, is_public, created_at, updated_at, is_deleted, is_approved)
	VALUES 
		(_form_id, _reporter_id, _animal_id, '0', NOW(), NOW(), '0', '0')
	RETURNING id INTO report_id;

	WITH FV AS
	(
		SELECT 
			FI.id field_id, value::json#>>'{}' field_value
		FROM 
			json_each(_field_values) FV
		INNER JOIN 
			"AP_Fields" FI ON FV.key = FI.name
	)
	INSERT INTO 
		"AP_Report_Field_Values" (report_id, field_id, value, created_at, updated_at)
	SELECT 
		report_id, field_id, field_value, NOW(), NOW()
	FROM FV;
	
	OPEN result FOR SELECT report_id as ID;

END;
$BODY$;

CREATE OR REPLACE FUNCTION AFN_AddAnimalReport(
	_animal_id BIGINT, 
	_reporter_id BIGINT,
	_form_id BIGINT,
	_field_values JSON
) 
RETURNS setof record
LANGUAGE 'plpgsql' 
AS $BODY$ 
DECLARE
	report_id BIGINT;
BEGIN
	INSERT INTO 
		"AP_Animal_Reports" (form_id, reporter_id, animal_id, is_public, created_at, updated_at, is_deleted, is_approved)
	VALUES 
		(_form_id, _reporter_id, _animal_id, '0', NOW(), NOW(), '0', '0')
	RETURNING "AP_Animal_Reports".id INTO report_id;

	WITH FV AS
	(
		SELECT 
			FI.id field_id, value::json#>>'{}' field_value
		FROM 
			json_each(_field_values) FV
		INNER JOIN 
			"AP_Fields" FI ON FV.key = FI.name
	)
	INSERT INTO 
		"AP_Report_Field_Values" (report_id, field_id, value, created_at, updated_at)
	SELECT 
		report_id, field_id, field_value, NOW(), NOW()
	FROM FV;

	RETURN QUERY SELECT AR.ID, to_char(AR.created_at, 'DD/MM/YYYY'), AR.is_approved, AN.name
	FROM "AP_Animal_Reports" AR
	INNER JOIN "AP_Animals" AN ON AR.animal_id = AN.id
	WHERE AR.id = report_id;

END;
$BODY$;
