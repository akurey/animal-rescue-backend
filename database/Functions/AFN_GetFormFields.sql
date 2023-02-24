CREATE OR REPLACE FUNCTION AFN_GetFormFields(
	pFormId BIGINT
) 
RETURNS TABLE(
	form_name VARCHAR(100), 
	form_section VARCHAR(100),
	field_id BIGINT,
	field_name VARCHAR(100), 
	is_required BIT, 
	field_type VARCHAR(50), 
	field_options TEXT
)
LANGUAGE 'plpgsql' 
AS $BODY$ 
BEGIN	
	RETURN QUERY 
		SELECT AF.name form_name, AFS.name form_section, AFI.id field_id, AFI.name field_name, 
		       AFI.is_required, AFT.name field_type, AFO.option
	    FROM "AP_Forms" AF
		INNER JOIN "AP_Form_Fields" AFF ON AFF.form_id = AF.id
		INNER JOIN "AP_Form_Sections" AFS ON AFS.id = AFF.form_section_id
		INNER JOIN "AP_Fields" AFI ON AFF.field_id = AFI.id
		INNER JOIN "AP_Field_Types" AFT ON AFT.id = AFI.field_type_id
		LEFT JOIN  "AP_Field_Options" AFO ON AFO.field_id = AFI.id
		WHERE AF.is_deleted = '0' AND AF.id = pFormId AND AFF.is_deleted = '0' 
		 AND AFI.is_deleted = '0' AND AFT.is_deleted = '0'
		 AND (AFO.is_deleted IS NULL OR AFO.is_deleted = '0')
		ORDER BY form_section, field_id;
END;
$BODY$;
