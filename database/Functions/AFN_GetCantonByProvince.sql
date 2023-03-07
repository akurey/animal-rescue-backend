CREATE OR REPLACE FUNCTION AFN_GetCantonByProvince(provinceID integer)
RETURNS TABLE(
    id BIGINT,
	Canton VARCHAR(50)
)
LANGUAGE 'plpgsql' 
AS $BODY$ 
BEGIN		
	RETURN QUERY 
		SELECT CA.id, CA.name as Canton 
		FROM "AP_Cantons" CA
        INNER JOIN "AP_Provinces" PR ON CA.province_id = PR.id
        WHERE CA.province_id = provinceID;
END;
$BODY$;