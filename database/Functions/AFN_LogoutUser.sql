CREATE OR REPLACE FUNCTION AFN_LogoutUser(pToken character varying(500))
RETURNS TEXT
      LANGUAGE 'plpgsql'
      AS $BODY$
DECLARE
    success_message TEXT := 'User sign out properly';
BEGIN
	UPDATE "AP_Users"
	SET token = NULL,
		refresh_token = NULL
	WHERE token = pToken;
	
	RETURN success_message;
END;
$BODY$;