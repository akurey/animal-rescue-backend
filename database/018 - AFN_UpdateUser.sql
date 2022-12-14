CREATE OR REPLACE FUNCTION AFN_UpdateUser(
	pUsername VARCHAR(100), 
	pToken VARCHAR(500),
	pRefreshToken VARCHAR(500)
) 
RETURNS setof record
LANGUAGE 'plpgsql' 
AS $BODY$ 
BEGIN
	UPDATE "AP_Users"
	SET token = pToken, refresh_token = pRefreshToken, updated_at = NOW()
	WHERE username = pUsername;
		
	RETURN QUERY 
        SELECT AU.first_name, AU.last_name, AU.username, AU.email, AU.identification, AU.sinac_registry, AU.token, AU.refresh_token
        FROM "AP_Users" AU
        WHERE AU.username = pUsername;
END;
$BODY$;
