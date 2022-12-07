CREATE OR REPLACE FUNCTION AFN_UpdateUser(
	pusername varchar, 
	ptoken varchar,
	prefresh_token varchar
) 
RETURNS setof record
LANGUAGE 'plpgsql' 
AS $BODY$ 
BEGIN
	UPDATE "AP_Users"
	SET token = ptoken, refresh_token = prefresh_token, updated_at = NOW()
	WHERE username = pusername;
		
	RETURN QUERY 
        SELECT UT.first_name, UT.last_name, UT.username, UT.email, UT.identification, UT.sinac_registry, UT.token, UT.refresh_token
        FROM "AP_Users" UT
        WHERE UT.username = pusername;
END;
$BODY$;
