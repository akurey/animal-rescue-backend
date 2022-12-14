CREATE OR REPLACE FUNCTION AFN_GetUser(pUserName VARCHAR(100))
RETURNS SETOF record
      LANGUAGE 'plpgsql'
      AS $BODY$
BEGIN
RETURN QUERY

    SELECT AU.id ID, AU.first_name First_name, AU.last_name Last_name, AU.username username, 
    AU.email pEmail, AU.password pPassword, AU.identification pIdentification, AU.sinac_registry pSinac, AU.token pToken, AU.refresh_token pRefreshToken
    FROM "AP_Users" AU
        WHERE AU.username = pUserName;
END;
$BODY$;