CREATE OR REPLACE FUNCTION AFN_AddUser(
	pname varchar(50), 
	plastName varchar(50),
	pusername varchar(100),
	pemail varchar(200),
    ppassword varchar(500),
	pidentification varchar(20),
	psinac varchar(20),
	ptoken varchar(500),
	prefresh_token varchar(500)
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
		(pname,plastName,pusername,pemail,ppassword,pidentification,psinac, ptoken, prefresh_token)
		RETURNING "AP_Users".id INTO userId;
   	RETURN QUERY
	SELECT AR.first_name,AR.last_name, AR.username, AR.email, AR.password, AR.identification, AR.sinac_registry, AR.token, AR.refresh_token
	FROM "AP_Users" AR;
END;
$BODY$;