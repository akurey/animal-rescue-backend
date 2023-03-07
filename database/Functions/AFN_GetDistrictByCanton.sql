CREATE OR REPLACE FUNCTION AFN_GetDistrictByCanton(cantonID integer)
RETURNS TABLE(
    id BIGINT,
	District VARCHAR(50)
)
LANGUAGE 'plpgsql' 
AS $BODY$ 
BEGIN		
	RETURN QUERY 
		SELECT DI.id, DI.name as District
        FROM "AP_Districts" DI
        INNER JOIN "AP_Cantons" CA ON DI.canton_id = CA.id
        WHERE CA.id = cantonID;
END;
$BODY$;