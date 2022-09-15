CREATE OR REPLACE FUNCTION AFN_UpdateAnimalReport(
	pReportId BIGINT, 
	pNewAnimalId BIGINT,
	pFieldValues JSON default '{}'
) 
RETURNS TABLE(report_id BIGINT, animal_id BIGINT, field_id BIGINT, value TEXT)
LANGUAGE 'plpgsql' 
AS $BODY$ 
BEGIN
	IF pNewAnimalId != 0 THEN
		UPDATE "AP_Animal_Reports" 
		SET animal_id = pNewAnimalId, updated_at = NOW()
		WHERE id = pReportId;
	END IF;
	
	WITH FV AS
	(
		SELECT pReportId report_id, FI.id field_id, FI.name field_name, FV.value::json#>>'{}' field_value
		FROM json_each(pFieldValues) FV
		INNER JOIN "AP_Fields" FI ON FV.key = FI.name
	)
		
	INSERT INTO "AP_Report_Field_Values" (report_id, field_id, value)
		SELECT FV.report_id, FV.field_id, FV.field_value
		FROM FV
	ON CONFLICT 
		ON CONSTRAINT "AP_Report_Field_Values_report_id_field_id_key" 
	DO 
		UPDATE SET value = EXCLUDED.value, updated_at = NOW()
		WHERE "AP_Report_Field_Values".report_id = EXCLUDED.report_id 
		AND "AP_Report_Field_Values".field_id = EXCLUDED.field_id;
		
	RETURN QUERY 
		SELECT RFV.report_id, AR.animal_id, RFV.field_id, RFV.value 
		FROM "AP_Report_Field_Values" RFV 
		INNER JOIN "AP_Animal_Reports" AR ON RFV.report_id = AR.id
		WHERE RFV.report_id = pReportId;
END;
$BODY$;
