DO $$
DECLARE var_userId BIGINT;
DECLARE var_shelterId BIGINT;
DECLARE var_roleXshelterId INTEGER;
DECLARE var_userXshelterId INTEGER;
DECLARE var_permissionId INTEGER;
DECLARE var_permissionXroleId INTEGER;
DECLARE var_userXroleId INTEGER;
DECLARE var_formType1 BIGINT;
DECLARE var_formType2 BIGINT;
DECLARE var_formId BIGINT;
BEGIN

    INSERT INTO public."AP_Users" (id, first_name, last_name, username, email, password, identification, sinac_registry)
    VALUES (1, 'Rodrigo', 'Navarro', 'rnavarro', 'rnavarro@akurey.com', 'rNavaroo7458', '304790505', '123')
        RETURNING id INTO var_userId;

    INSERT INTO public."AP_Management_Category"(id, name)
    VALUES(1, 'Zoologico');

    -- TODO: Fill these direction tables properly
    INSERT INTO public."AP_Provinces"(id, name)
    VALUES(1, 'San José');

    INSERT INTO public."AP_Cantons"(id, name, province_id)
    VALUES(1, 'Santa Ana', 1);

    INSERT INTO public."AP_Districts"(id, name, canton_id)
    VALUES(1, 'Santa Ana', 1);

    INSERT INTO public."AP_Directions"(id, exact_direction, district_id)
    VALUES(1, 'Alto de las Palomas', 1);

    INSERT INTO public."AP_Shelters"(name, trade_name, management_category_id, sinac_resolution_number, direction_id, phone, owner, regent_biologist, regent_vet)
    VALUES ('Refugio Animal Costa Rica', 'Refugio Santa Ana', 1, '123', 1, '22824614', 1, 1, 1)
        RETURNING id INTO var_shelterId;

    INSERT INTO public."AP_Role_Shelters"(name, description, shelter_id)
    VALUES ('Admin', 'Super usuario', var_shelterId)
        RETURNING id INTO var_roleXshelterId;

    INSERT INTO public."AP_User_Shelters"(user_id, shelter_id)
    VALUES (var_userId, var_shelterId)
        RETURNING id INTO var_userXshelterId;

    INSERT INTO public."AP_Permissions"(name, description)
    VALUES ('AnimalEntryRegisterer', 'Registra entradas de animales')
        RETURNING id INTO var_permissionId;

    INSERT INTO public."AP_Permissions_Role"(permission_id, role_id)
    VALUES (var_permissionId, var_roleXshelterId)
        RETURNING id INTO var_permissionXroleId;

    INSERT INTO public."AP_User_Role"(user_id, role_id)
    VALUES (var_userId, var_roleXshelterId)
        RETURNING id INTO var_userXroleId;

    INSERT INTO public."AP_Form_Types"(type)
    VALUES ('Report')
        RETURNING id INTO var_formType1;

    INSERT INTO public."AP_Form_Types"(type)
    VALUES ('Medic')
        RETURNING id INTO var_formType2;

    INSERT INTO public."AP_Forms"(shelter_id, type_id, name)
    VALUES (var_shelterId, var_formType1, 'Registro Animal')
        RETURNING id INTO var_formId;

    INSERT INTO public."AP_Field_Types"(id, name)
    VALUES (1, 'Textbox'), (2, 'Textarea'), (3, 'Numeric'), (4, 'Money'), (5, 'Cedula'),
           (6, 'Phone'), (7, 'Date'), (8, 'Dropdown'), (9, 'Radio'), (10, 'Checkbox'), (11, 'Address');

    INSERT INTO public."AP_Fields"(id, field_type_id, name, is_required)
    VALUES (1, 8, 'Sexo', B'1'),
           (2, 8, 'Color', B'1'),
           (3, 1, 'Número de identificación (CHIP)', B'1'),
           (4, 1, 'Procedencia', B'1'),
           (5, 2, 'Datos del padre', B'1'),
           (6, 2, 'Datos de la madre', B'1'),
           (7, 8, 'Razón de recepción', B'1'),
           (8, 8, 'Tipo de Caso', B'1'),
           (9, 11,'Dirección', B'1'),
           (10, 2, 'Observaciones', B'0'),
           (11, 1, 'Nombre', B'1'),
           (12, 1, 'Apellidos', B'1'),
           (13, 5, 'Cédula', B'1'),
           (14, 6, 'Teléfono', B'1'),
           (15, 2, 'Motivo de entrega', B'1'),
           (16, 2, 'Observaciones', B'0');

    INSERT INTO public."AP_Field_Options"(field_id, option)
    VALUES (1, '[Macho, Hembra, Indefinido]'),
           (2, '[Café, Verde, Negro, Azul, Rojo, Amarillo]'),
           (7, '[Emergencia Medica, Ingreso Temporal]'),
           (8, '[Rescate, Tenencia Irregular, Tenencia Legal, Decomiso]');

    INSERT INTO public."AP_Form_Fields"(form_id, field_id, is_public)
    VALUES (var_formId, 1, B'0'), (var_formId, 2, B'0'), (var_formId, 3, B'0'),
           (var_formId, 4, B'0'), (var_formId, 5, B'0'), (var_formId, 6, B'0'),
           (var_formId, 7, B'0'), (var_formId, 8, B'0'), (var_formId, 9, B'0'),
           (var_formId, 10, B'0'), (var_formId, 11, B'0'), (var_formId, 12, B'0'),
           (var_formId, 13, B'0'), (var_formId, 14, B'0'), (var_formId, 15, B'0'),
           (var_formId, 16, B'0');

    INSERT INTO public."AP_Conservation_Status"(id, name, abbreviation)
    VALUES (1, 'Preocupación menor', 'LC'), (2, 'Casi amenazada', 'NT'), (3, 'Vulnerable', 'VU'), (4, 'En peligro', 'EN'),
           (5, 'En peligro crítico', 'CR'), (6, 'Extinta en estado silvestre', 'EW'), (7, 'Extinta', 'EX'), (8, 'Datos Insuficientes', 'DD');

    INSERT INTO public."AP_Animal_Classification"(id, name)
    VALUES (1, 'Mamifero'), (2, 'Reptil'), (3, 'Ave'), (4, 'Anfibio'), (5, 'Pez'), (6, 'Arthropodo');

    INSERT INTO public."AP_Animals"(name, scientific_name, conservation_status_id, classification_id)
    VALUES ('Danta', 'Tapirus bairdii', 4, 1),
           ('Jaguar', 'Panthera onca', 2, 1),
           ('Lapa Roja', 'Ara macao', 1, 3),
           ('Terciopelo', 'Bothrops asper', 1, 2),
           ('Rana de Ojos Rojos', 'Agalychnis callidryas', 1, 4);

    INSERT INTO "AP_Form_Sections" (name, form_id)
    VALUES ('Datos del animal', 1), 
           ('Datos del rescate', 1), 
           ('Datos del rescatista', 1);

    UPDATE "AP_Form_Fields" SET form_section_id = 1 WHERE form_id=1 AND field_id <= 6;
    UPDATE "AP_Form_Fields" SET form_section_id = 2 WHERE form_id=1 AND field_id > 6 AND field_id <= 10;
    UPDATE "AP_Form_Fields" SET form_section_id = 3 WHERE form_id=1 AND field_id > 10;

END $$;

