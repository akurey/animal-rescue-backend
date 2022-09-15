CREATE OR REPLACE FUNCTION AFN_DeleteAnimalReport(
	pReportId BIGINT
) 
RETURNS TABLE(id BIGINT)
LANGUAGE 'plpgsql' 
AS $BODY$ 
BEGIN
	RETURN QUERY
	UPDATE "AP_Animal_Reports" AR SET is_deleted = '1'
	WHERE AR.id = pReportId
	RETURNING AR.id;
END;
$BODY$;
