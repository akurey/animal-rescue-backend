CREATE OR REPLACE FUNCTION CAST_TO_INT(TEXT, INTEGER) RETURNS INTEGER
AS $$
BEGIN
RETURN CAST($1 AS INTEGER);
EXCEPTION
    	WHEN invalid_text_representation THEN
        	RETURN $2;
END;
$$ LANGUAGE plpgsql IMMUTABLE;