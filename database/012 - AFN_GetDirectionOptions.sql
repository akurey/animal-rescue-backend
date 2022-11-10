CREATE OR REPLACE FUNCTION AFN_GetDirectionOptions() 
RETURNS TABLE(
	province VARCHAR(50),
	canton VARCHAR(50),
	district VARCHAR(50)
)
LANGUAGE 'plpgsql' 
AS $BODY$ 
BEGIN		
	RETURN QUERY 
		SELECT PR.name, CA.name, DI.name FROM "AP_Districts" DI
		INNER JOIN "AP_Cantons" CA ON DI.canton_id = CA.id
		INNER JOIN "AP_Provinces" PR ON CA.province_id = PR.id;
END;
$BODY$;