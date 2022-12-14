CREATE OR REPLACE FUNCTION AFN_AddUser(
	pName VARCHAR(50), 
	pLastName VARCHAR(50),
	pUsername VARCHAR(100),
	pEmail VARCHAR(200),
    pPassword VARCHAR(500),
	pIdentification VARCHAR(20),
	pSinac VARCHAR(20),
	pToken VARCHAR(500),
	pRefreshToken VARCHAR(500)
) 
RETURNS setof record
LANGUAGE 'plpgsql' 
AS $BODY$ 
DECLARE
	userId BIGINT;
BEGIN
	INSERT INTO 
		"AP_Users" (first_name, last_name, username, email, password, identification, sinac_registry, token, refresh_token)
	VALUES 
		(pName,pLastName,pUsername,pEmail,pPassword,pIdentification,pSinac, pToken, pRefreshToken)
		RETURNING "AP_Users".id INTO userId;
   	RETURN QUERY
	SELECT AU.first_name,AU.last_name, AU.username, AU.email, AU.password, AU.identification, AU.sinac_registry, AU.token, AU.refresh_token
	FROM "AP_Users" AU
	WHERE AU.id = userId;
END;
$BODY$;