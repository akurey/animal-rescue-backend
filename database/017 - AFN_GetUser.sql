CREATE OR REPLACE FUNCTION AFN_GetUser(pUserName varchar)
RETURNS SETOF record
      LANGUAGE 'plpgsql'
      AS $BODY$
BEGIN
RETURN QUERY

    SELECT AR.id ID, AR.first_name First_name, AR.last_name Last_name, AR.username username, 
    AR.email Email, AR.password Passwordd, AR.identification Identification, AR.sinac_registry Sinac, AR.token Token, AR.refresh_token RefreshToken
    FROM "AP_Users" AR
        WHERE AR.username = pUserName;
END;
$BODY$;