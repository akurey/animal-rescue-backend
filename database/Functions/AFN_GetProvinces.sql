CREATE OR REPLACE FUNCTION AFN_GetProvinces()
RETURNS TABLE(
    id BIGINT,
	Province VARCHAR(50)
)
LANGUAGE 'plpgsql' 
AS $BODY$ 
BEGIN		
	RETURN QUERY 
		SELECT PR.id, PR.name as Province
		FROM "AP_Provinces" PR;
END;
$BODY$;