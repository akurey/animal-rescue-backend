CREATE OR REPLACE FUNCTION AFN_UpdateAnimalReport(
	pReportId BIGINT, 
	pFieldValues JSON default '{}'
) 
RETURNS TABLE(report_id BIGINT, field_id BIGINT, value TEXT)
LANGUAGE 'plpgsql' 
AS $BODY$ 
BEGIN	
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
		UPDATE SET value = EXCLUDED.value
		WHERE "AP_Report_Field_Values".report_id = EXCLUDED.report_id 
		AND "AP_Report_Field_Values".field_id = EXCLUDED.field_id;
		
	RETURN QUERY 
		SELECT RFV.report_id, RFV.field_id, RFV.value FROM "AP_Report_Field_Values" RFV 
		WHERE RFV.report_id = pReportId;
END;
$BODY$;
