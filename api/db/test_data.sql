PGDMP                      |           erp    16.2    16.2 �    3           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            4           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            5           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            6           1262    17276    erp    DATABASE     i   CREATE DATABASE erp WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'pt_PT';
    DROP DATABASE erp;
                hestia_erp_internal    false                        2615    17278 	   companies    SCHEMA        CREATE SCHEMA companies;
    DROP SCHEMA companies;
                hestia_erp_internal    false                        2615    17283    general    SCHEMA        CREATE SCHEMA general;
    DROP SCHEMA general;
                hestia_erp_internal    false            	            2615    17279    hr    SCHEMA        CREATE SCHEMA hr;
    DROP SCHEMA hr;
                hestia_erp_internal    false                        2615    17281 	   inventory    SCHEMA        CREATE SCHEMA inventory;
    DROP SCHEMA inventory;
                hestia_erp_internal    false                        2615    17284    itv    SCHEMA        CREATE SCHEMA itv;
    DROP SCHEMA itv;
                hestia_erp_internal    false            
            2615    17280    sales    SCHEMA        CREATE SCHEMA sales;
    DROP SCHEMA sales;
                hestia_erp_internal    false                        2615    17282    subscriptions    SCHEMA        CREATE SCHEMA subscriptions;
    DROP SCHEMA subscriptions;
                hestia_erp_internal    false                        2615    17277    users    SCHEMA        CREATE SCHEMA users;
    DROP SCHEMA users;
                hestia_erp_internal    false                        3079    17390    pgcrypto 	   EXTENSION     <   CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;
    DROP EXTENSION pgcrypto;
                   false            7           0    0    EXTENSION pgcrypto    COMMENT     <   COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';
                        false    2            �            1259    17297    company    TABLE     ;  CREATE TABLE companies.company (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    commercial_name text,
    vat_id text NOT NULL,
    ssn text NOT NULL,
    street text NOT NULL,
    locality text NOT NULL,
    postal_code text NOT NULL,
    country_code character varying(2) NOT NULL
);
    DROP TABLE companies.company;
    	   companies         heap    postgres    false    8            8           0    0    COLUMN company.id    COMMENT     ?   COMMENT ON COLUMN companies.company.id IS 'Company ID (UUID)';
       	   companies          postgres    false    225            9           0    0    COLUMN company.commercial_name    COMMENT     e   COMMENT ON COLUMN companies.company.commercial_name IS 'Commercial name like Apple, Google, etc...';
       	   companies          postgres    false    225            �            1259    17460    location    TABLE     /  CREATE TABLE companies.location (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    name text NOT NULL,
    street text NOT NULL,
    locality text NOT NULL,
    postal_code text NOT NULL,
    country_code character varying(2) DEFAULT 'PT'::character varying NOT NULL
);
    DROP TABLE companies.location;
    	   companies         heap    hestia_erp_internal    false    8            :           0    0    COLUMN location.name    COMMENT     >   COMMENT ON COLUMN companies.location.name IS 'Location Name';
       	   companies          hestia_erp_internal    false    242            �            1259    17477 
   preference    TABLE     �   CREATE TABLE companies.preference (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    code text NOT NULL,
    value text NOT NULL
);
 !   DROP TABLE companies.preference;
    	   companies         heap    postgres    false    8            �            1259    17469    country    TABLE     �   CREATE TABLE general.country (
    code character varying(2) NOT NULL,
    name text NOT NULL,
    active boolean DEFAULT true NOT NULL
);
    DROP TABLE general.country;
       general         heap    hestia_erp_internal    false    13            ;           0    0    COLUMN country.code    COMMENT     E   COMMENT ON COLUMN general.country.code IS 'ISO 3166-1 alpha-2 code';
          general          hestia_erp_internal    false    243            <           0    0    COLUMN country.name    COMMENT     E   COMMENT ON COLUMN general.country.name IS 'English name of country';
          general          hestia_erp_internal    false    243            =           0    0    COLUMN country.active    COMMENT     d   COMMENT ON COLUMN general.country.active IS 'Is the Country active? Say the country was abolished';
          general          hestia_erp_internal    false    243            �            1259    17347    device    TABLE       CREATE TABLE hr.device (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid,
    active boolean DEFAULT true NOT NULL,
    name character varying(70) NOT NULL,
    key uuid DEFAULT gen_random_uuid() NOT NULL,
    last_date timestamp with time zone NOT NULL
);
    DROP TABLE hr.device;
       hr         heap    hestia_erp_internal    false    9            �            1259    17309    employee    TABLE     �   CREATE TABLE hr.employee (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    name text NOT NULL,
    location_id uuid NOT NULL
);
    DROP TABLE hr.employee;
       hr         heap    postgres    false    9            >           0    0    COLUMN employee.id    COMMENT     B   COMMENT ON COLUMN hr.employee.id IS 'HR''s employee internal id';
          hr          postgres    false    226            ?           0    0    COLUMN employee.location_id    COMMENT     I   COMMENT ON COLUMN hr.employee.location_id IS 'Does it need to be NULL?';
          hr          postgres    false    226            �            1259    17334    employee_identification_doc    TABLE     �   CREATE TABLE hr.employee_identification_doc (
    employee_id uuid NOT NULL,
    identification_doc_id uuid NOT NULL,
    doc_number character varying(50) NOT NULL
);
 +   DROP TABLE hr.employee_identification_doc;
       hr         heap    hestia_erp_internal    false    9            �            1259    17339    identification_doc    TABLE     �   CREATE TABLE hr.identification_doc (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    code character varying(20),
    name text NOT NULL,
    description text
);
 "   DROP TABLE hr.identification_doc;
       hr         heap    hestia_erp_internal    false    9            �            1259    17322 	   timesheet    TABLE     �   CREATE TABLE hr.timesheet (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    employee_id uuid NOT NULL,
    device uuid,
    type uuid NOT NULL,
    date timestamp with time zone NOT NULL
);
    DROP TABLE hr.timesheet;
       hr         heap    hestia_erp_internal    false    9            �            1259    17355    timesheet_type    TABLE     �   CREATE TABLE hr.timesheet_type (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    code character varying(20) NOT NULL,
    name character varying(100)
);
    DROP TABLE hr.timesheet_type;
       hr         heap    hestia_erp_internal    false    9            �            1259    17328    vacation    TABLE     �   CREATE TABLE hr.vacation (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    employee_id uuid NOT NULL,
    start_date timestamp with time zone NOT NULL,
    end_date timestamp with time zone NOT NULL,
    approved_by uuid
);
    DROP TABLE hr.vacation;
       hr         heap    hestia_erp_internal    false    9            �            1259    17384    category    TABLE     �   CREATE TABLE inventory.category (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    name character varying(40) NOT NULL,
    parent_id uuid
);
    DROP TABLE inventory.category;
    	   inventory         heap    hestia_erp_internal    false    11            �            1259    17367    movement    TABLE     T   CREATE TABLE inventory.movement (
    id uuid DEFAULT gen_random_uuid() NOT NULL
);
    DROP TABLE inventory.movement;
    	   inventory         heap    hestia_erp_internal    false    11            �            1259    17371    product    TABLE     R  CREATE TABLE inventory.product (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    active boolean DEFAULT true NOT NULL,
    category_id uuid NOT NULL,
    company_id uuid NOT NULL,
    name text NOT NULL,
    reference character varying(40) NOT NULL,
    barcode character varying(13),
    description text,
    stock numeric(14,5)
);
    DROP TABLE inventory.product;
    	   inventory         heap    hestia_erp_internal    false    11            @           0    0    COLUMN product.barcode    COMMENT     H   COMMENT ON COLUMN inventory.product.barcode IS 'Barcode type: EAN-13A';
       	   inventory          hestia_erp_internal    false    236            A           0    0    COLUMN product.stock    COMMENT     v   COMMENT ON COLUMN inventory.product.stock IS 'If stock is null then stock is not enabled
100 digits good enough no?';
       	   inventory          hestia_erp_internal    false    236            �            1259    17524 
   department    TABLE     �   CREATE TABLE itv.department (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    name text NOT NULL
);
    DROP TABLE itv.department;
       itv         heap    hestia_erp_internal    false    14            �            1259    17547    family    TABLE     �   CREATE TABLE itv.family (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    name text NOT NULL
);
    DROP TABLE itv.family;
       itv         heap    hestia_erp_internal    false    14            B           0    0    TABLE family    COMMENT     c   COMMENT ON TABLE itv.family IS 'Family of the product (like Pijamas, Sweater, Sweater w/ Hoodie)';
          itv          hestia_erp_internal    false    252            �            1259    17555    gender    TABLE     �   CREATE TABLE itv.gender (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    gender text NOT NULL
);
    DROP TABLE itv.gender;
       itv         heap    hestia_erp_internal    false    14            �            1259    17539    raw_material    TABLE     �   CREATE TABLE itv.raw_material (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    code text,
    material text NOT NULL
);
    DROP TABLE itv.raw_material;
       itv         heap    hestia_erp_internal    false    14            �            1259    17509    size    TABLE     �   CREATE TABLE itv.size (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    size text NOT NULL
);
    DROP TABLE itv.size;
       itv         heap    hestia_erp_internal    false    14            C           0    0    COLUMN size.size    COMMENT     B   COMMENT ON COLUMN itv.size.size IS 'Size (XS, S, M, L, XL,....)';
          itv          hestia_erp_internal    false    247            �            1259    17499    technical_file    TABLE     �   CREATE TABLE itv.technical_file (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    code text NOT NULL,
    client_id uuid NOT NULL,
    department_id uuid NOT NULL
);
    DROP TABLE itv.technical_file;
       itv         heap    hestia_erp_internal    false    14            �            1259    17532    technical_file_size    TABLE     {   CREATE TABLE itv.technical_file_size (
    technical_file_id uuid NOT NULL,
    size_id uuid NOT NULL,
    comment text
);
 $   DROP TABLE itv.technical_file_size;
       itv         heap    hestia_erp_internal    false    14            �            1259    17487    client    TABLE     6  CREATE TABLE sales.client (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    name text NOT NULL,
    code text NOT NULL,
    vat_id text NOT NULL,
    street text NOT NULL,
    postal_code text NOT NULL,
    locality text NOT NULL,
    country character varying(2) NOT NULL
);
    DROP TABLE sales.client;
       sales         heap    postgres    false    10            �            1259    17451    doccument_series    TABLE     �   CREATE TABLE sales.doccument_series (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    code character varying(80) DEFAULT 'H'::character varying NOT NULL,
    description text
);
 #   DROP TABLE sales.doccument_series;
       sales         heap    hestia_erp_internal    false    10            �            1259    17361    document    TABLE     n   CREATE TABLE sales.document (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL
);
    DROP TABLE sales.document;
       sales         heap    hestia_erp_internal    false    10            �            1259    17436    module    TABLE     �   CREATE TABLE subscriptions.module (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    code character varying(20) NOT NULL,
    name text NOT NULL,
    default_price numeric NOT NULL
);
 !   DROP TABLE subscriptions.module;
       subscriptions         heap    hestia_erp_internal    false    12            �            1259    17427    subscription    TABLE     �  CREATE TABLE subscriptions.subscription (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    user_limit integer NOT NULL,
    user_price numeric NOT NULL,
    start_date timestamp with time zone NOT NULL,
    end_date timestamp with time zone NOT NULL,
    creation_date timestamp with time zone DEFAULT now() NOT NULL,
    update_date timestamp with time zone NOT NULL
);
 '   DROP TABLE subscriptions.subscription;
       subscriptions         heap    hestia_erp_internal    false    12            �            1259    17444    subscription_module    TABLE     �   CREATE TABLE subscriptions.subscription_module (
    subscription_id uuid NOT NULL,
    module_id uuid NOT NULL,
    price numeric NOT NULL
);
 .   DROP TABLE subscriptions.subscription_module;
       subscriptions         heap    hestia_erp_internal    false    12            �            1259    17317    user_company    TABLE     s   CREATE TABLE users.user_company (
    user_id uuid NOT NULL,
    company_id uuid NOT NULL,
    employee_id uuid
);
    DROP TABLE users.user_company;
       users         heap    hestia_erp_internal    false    7            �            1259    17285    users    TABLE     I  CREATE TABLE users.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    salt text NOT NULL,
    timezone text DEFAULT 'Etc/GMT'::character varying NOT NULL,
    CONSTRAINT ck_users_timezone CHECK (((now() AT TIME ZONE timezone) IS NOT NULL))
);
    DROP TABLE users.users;
       users         heap    hestia_erp_internal    false    7            D           0    0    TABLE users    COMMENT     j   COMMENT ON TABLE users.users IS 'This is the only table named in plural because USER is a reserved word';
          users          hestia_erp_internal    false    224            E           0    0    COLUMN users.id    COMMENT     L   COMMENT ON COLUMN users.users.id IS 'User''s id (uuid) generated randomly';
          users          hestia_erp_internal    false    224            F           0    0    COLUMN users.name    COMMENT     �   COMMENT ON COLUMN users.users.name IS 'This is the user full name used in the Subscription Management Portal, if used in atual ERP it should use the employee name, if employee name is not available use the user name';
          users          hestia_erp_internal    false    224            G           0    0    COLUMN users.email    COMMENT     8   COMMENT ON COLUMN users.users.email IS 'User''s email';
          users          hestia_erp_internal    false    224            H           0    0    COLUMN users.password    COMMENT     J   COMMENT ON COLUMN users.users.password IS 'User''s password (in SHA512)';
          users          hestia_erp_internal    false    224            I           0    0    COLUMN users.salt    COMMENT     ]   COMMENT ON COLUMN users.users.salt IS 'User''s password salt in string format (maybe hex?)';
          users          hestia_erp_internal    false    224            J           0    0    COLUMN users.timezone    COMMENT     >   COMMENT ON COLUMN users.users.timezone IS 'User''s timezone';
          users          hestia_erp_internal    false    224            K           0    0 %   CONSTRAINT ck_users_timezone ON users    COMMENT     d   COMMENT ON CONSTRAINT ck_users_timezone ON users.users IS 'Check if the inputed timezone is valid';
          users          hestia_erp_internal    false    224            �            1259    17517    users_session    TABLE     �   CREATE TABLE users.users_session (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    expiry_date timestamp with time zone DEFAULT (now() + '08:00:00'::interval) NOT NULL
);
     DROP TABLE users.users_session;
       users         heap    postgres    false    7                      0    17297    company 
   TABLE DATA           y   COPY companies.company (id, name, commercial_name, vat_id, ssn, street, locality, postal_code, country_code) FROM stdin;
 	   companies          postgres    false    225   �       %          0    17460    location 
   TABLE DATA           h   COPY companies.location (id, company_id, name, street, locality, postal_code, country_code) FROM stdin;
 	   companies          hestia_erp_internal    false    242   ��       '          0    17477 
   preference 
   TABLE DATA           D   COPY companies.preference (id, company_id, code, value) FROM stdin;
 	   companies          postgres    false    244   ��       &          0    17469    country 
   TABLE DATA           6   COPY general.country (code, name, active) FROM stdin;
    general          hestia_erp_internal    false    243   ��                 0    17347    device 
   TABLE DATA           J   COPY hr.device (id, company_id, active, name, key, last_date) FROM stdin;
    hr          hestia_erp_internal    false    232   ��                 0    17309    employee 
   TABLE DATA           A   COPY hr.employee (id, company_id, name, location_id) FROM stdin;
    hr          postgres    false    226   �                 0    17334    employee_identification_doc 
   TABLE DATA           a   COPY hr.employee_identification_doc (employee_id, identification_doc_id, doc_number) FROM stdin;
    hr          hestia_erp_internal    false    230   3�                 0    17339    identification_doc 
   TABLE DATA           E   COPY hr.identification_doc (id, code, name, description) FROM stdin;
    hr          hestia_erp_internal    false    231   P�                 0    17322 	   timesheet 
   TABLE DATA           D   COPY hr.timesheet (id, employee_id, device, type, date) FROM stdin;
    hr          hestia_erp_internal    false    228   m�                 0    17355    timesheet_type 
   TABLE DATA           4   COPY hr.timesheet_type (id, code, name) FROM stdin;
    hr          hestia_erp_internal    false    233   ��                 0    17328    vacation 
   TABLE DATA           R   COPY hr.vacation (id, employee_id, start_date, end_date, approved_by) FROM stdin;
    hr          hestia_erp_internal    false    229   ��                  0    17384    category 
   TABLE DATA           F   COPY inventory.category (id, company_id, name, parent_id) FROM stdin;
 	   inventory          hestia_erp_internal    false    237   ��                 0    17367    movement 
   TABLE DATA           )   COPY inventory.movement (id) FROM stdin;
 	   inventory          hestia_erp_internal    false    235   ��                 0    17371    product 
   TABLE DATA           w   COPY inventory.product (id, active, category_id, company_id, name, reference, barcode, description, stock) FROM stdin;
 	   inventory          hestia_erp_internal    false    236   ��       ,          0    17524 
   department 
   TABLE DATA           7   COPY itv.department (id, company_id, name) FROM stdin;
    itv          hestia_erp_internal    false    249   �       /          0    17547    family 
   TABLE DATA           3   COPY itv.family (id, company_id, name) FROM stdin;
    itv          hestia_erp_internal    false    252   8�       0          0    17555    gender 
   TABLE DATA           5   COPY itv.gender (id, company_id, gender) FROM stdin;
    itv          hestia_erp_internal    false    253   U�       .          0    17539    raw_material 
   TABLE DATA           C   COPY itv.raw_material (id, company_id, code, material) FROM stdin;
    itv          hestia_erp_internal    false    251   r�       *          0    17509    size 
   TABLE DATA           1   COPY itv.size (id, company_id, size) FROM stdin;
    itv          hestia_erp_internal    false    247   ��       )          0    17499    technical_file 
   TABLE DATA           U   COPY itv.technical_file (id, company_id, code, client_id, department_id) FROM stdin;
    itv          hestia_erp_internal    false    246   ��       -          0    17532    technical_file_size 
   TABLE DATA           O   COPY itv.technical_file_size (technical_file_id, size_id, comment) FROM stdin;
    itv          hestia_erp_internal    false    250   ��       (          0    17487    client 
   TABLE DATA           k   COPY sales.client (id, company_id, name, code, vat_id, street, postal_code, locality, country) FROM stdin;
    sales          postgres    false    245   ��       $          0    17451    doccument_series 
   TABLE DATA           @   COPY sales.doccument_series (id, code, description) FROM stdin;
    sales          hestia_erp_internal    false    241   �                 0    17361    document 
   TABLE DATA           1   COPY sales.document (id, company_id) FROM stdin;
    sales          hestia_erp_internal    false    234    �       "          0    17436    module 
   TABLE DATA           F   COPY subscriptions.module (id, code, name, default_price) FROM stdin;
    subscriptions          hestia_erp_internal    false    239   =�       !          0    17427    subscription 
   TABLE DATA           �   COPY subscriptions.subscription (id, company_id, user_limit, user_price, start_date, end_date, creation_date, update_date) FROM stdin;
    subscriptions          hestia_erp_internal    false    238   Z�       #          0    17444    subscription_module 
   TABLE DATA           W   COPY subscriptions.subscription_module (subscription_id, module_id, price) FROM stdin;
    subscriptions          hestia_erp_internal    false    240   w�                 0    17317    user_company 
   TABLE DATA           G   COPY users.user_company (user_id, company_id, employee_id) FROM stdin;
    users          hestia_erp_internal    false    227   ��                 0    17285    users 
   TABLE DATA           I   COPY users.users (id, name, email, password, salt, timezone) FROM stdin;
    users          hestia_erp_internal    false    224   ��       +          0    17517    users_session 
   TABLE DATA           @   COPY users.users_session (id, user_id, expiry_date) FROM stdin;
    users          postgres    false    248   ��                  2606    17304    company pk_company_id 
   CONSTRAINT     V   ALTER TABLE ONLY companies.company
    ADD CONSTRAINT pk_company_id PRIMARY KEY (id);
 B   ALTER TABLE ONLY companies.company DROP CONSTRAINT pk_company_id;
    	   companies            postgres    false    225            @           2606    17468    location pk_location 
   CONSTRAINT     U   ALTER TABLE ONLY companies.location
    ADD CONSTRAINT pk_location PRIMARY KEY (id);
 A   ALTER TABLE ONLY companies.location DROP CONSTRAINT pk_location;
    	   companies            hestia_erp_internal    false    242            D           2606    17484    preference pk_preference 
   CONSTRAINT     Y   ALTER TABLE ONLY companies.preference
    ADD CONSTRAINT pk_preference PRIMARY KEY (id);
 E   ALTER TABLE ONLY companies.preference DROP CONSTRAINT pk_preference;
    	   companies            postgres    false    244                       2606    17308    company uq_company_ssn 
   CONSTRAINT     S   ALTER TABLE ONLY companies.company
    ADD CONSTRAINT uq_company_ssn UNIQUE (ssn);
 C   ALTER TABLE ONLY companies.company DROP CONSTRAINT uq_company_ssn;
    	   companies            postgres    false    225                       2606    17306    company uq_company_vat_id 
   CONSTRAINT     Y   ALTER TABLE ONLY companies.company
    ADD CONSTRAINT uq_company_vat_id UNIQUE (vat_id);
 F   ALTER TABLE ONLY companies.company DROP CONSTRAINT uq_company_vat_id;
    	   companies            postgres    false    225            F           2606    17486    preference uq_preference_code 
   CONSTRAINT     g   ALTER TABLE ONLY companies.preference
    ADD CONSTRAINT uq_preference_code UNIQUE (company_id, code);
 J   ALTER TABLE ONLY companies.preference DROP CONSTRAINT uq_preference_code;
    	   companies            postgres    false    244    244            B           2606    17476    country pk_country 
   CONSTRAINT     S   ALTER TABLE ONLY general.country
    ADD CONSTRAINT pk_country PRIMARY KEY (code);
 =   ALTER TABLE ONLY general.country DROP CONSTRAINT pk_country;
       general            hestia_erp_internal    false    243            *           2606    17354    device pk_device 
   CONSTRAINT     J   ALTER TABLE ONLY hr.device
    ADD CONSTRAINT pk_device PRIMARY KEY (id);
 6   ALTER TABLE ONLY hr.device DROP CONSTRAINT pk_device;
       hr            hestia_erp_internal    false    232                       2606    17316    employee pk_employee 
   CONSTRAINT     N   ALTER TABLE ONLY hr.employee
    ADD CONSTRAINT pk_employee PRIMARY KEY (id);
 :   ALTER TABLE ONLY hr.employee DROP CONSTRAINT pk_employee;
       hr            postgres    false    226            &           2606    17338 .   employee_identification_doc pk_employee_doc_id 
   CONSTRAINT     �   ALTER TABLE ONLY hr.employee_identification_doc
    ADD CONSTRAINT pk_employee_doc_id PRIMARY KEY (employee_id, identification_doc_id);
 T   ALTER TABLE ONLY hr.employee_identification_doc DROP CONSTRAINT pk_employee_doc_id;
       hr            hestia_erp_internal    false    230    230            (           2606    17346 $   identification_doc pk_identification 
   CONSTRAINT     ^   ALTER TABLE ONLY hr.identification_doc
    ADD CONSTRAINT pk_identification PRIMARY KEY (id);
 J   ALTER TABLE ONLY hr.identification_doc DROP CONSTRAINT pk_identification;
       hr            hestia_erp_internal    false    231            "           2606    17327    timesheet pk_timesheet 
   CONSTRAINT     P   ALTER TABLE ONLY hr.timesheet
    ADD CONSTRAINT pk_timesheet PRIMARY KEY (id);
 <   ALTER TABLE ONLY hr.timesheet DROP CONSTRAINT pk_timesheet;
       hr            hestia_erp_internal    false    228            ,           2606    17360     timesheet_type pk_timesheet_type 
   CONSTRAINT     Z   ALTER TABLE ONLY hr.timesheet_type
    ADD CONSTRAINT pk_timesheet_type PRIMARY KEY (id);
 F   ALTER TABLE ONLY hr.timesheet_type DROP CONSTRAINT pk_timesheet_type;
       hr            hestia_erp_internal    false    233            $           2606    17333    vacation pk_vacation 
   CONSTRAINT     N   ALTER TABLE ONLY hr.vacation
    ADD CONSTRAINT pk_vacation PRIMARY KEY (id);
 :   ALTER TABLE ONLY hr.vacation DROP CONSTRAINT pk_vacation;
       hr            hestia_erp_internal    false    229            6           2606    17389    category pk_category 
   CONSTRAINT     U   ALTER TABLE ONLY inventory.category
    ADD CONSTRAINT pk_category PRIMARY KEY (id);
 A   ALTER TABLE ONLY inventory.category DROP CONSTRAINT pk_category;
    	   inventory            hestia_erp_internal    false    237            0           2606    17379    product pk_product 
   CONSTRAINT     S   ALTER TABLE ONLY inventory.product
    ADD CONSTRAINT pk_product PRIMARY KEY (id);
 ?   ALTER TABLE ONLY inventory.product DROP CONSTRAINT pk_product;
    	   inventory            hestia_erp_internal    false    236            2           2606    17383    product uq_product_barcode 
   CONSTRAINT     [   ALTER TABLE ONLY inventory.product
    ADD CONSTRAINT uq_product_barcode UNIQUE (barcode);
 G   ALTER TABLE ONLY inventory.product DROP CONSTRAINT uq_product_barcode;
    	   inventory            hestia_erp_internal    false    236            4           2606    17381    product uq_product_reference 
   CONSTRAINT     _   ALTER TABLE ONLY inventory.product
    ADD CONSTRAINT uq_product_reference UNIQUE (reference);
 I   ALTER TABLE ONLY inventory.product DROP CONSTRAINT uq_product_reference;
    	   inventory            hestia_erp_internal    false    236            V           2606    17531    department pk_department 
   CONSTRAINT     S   ALTER TABLE ONLY itv.department
    ADD CONSTRAINT pk_department PRIMARY KEY (id);
 ?   ALTER TABLE ONLY itv.department DROP CONSTRAINT pk_department;
       itv            hestia_erp_internal    false    249            \           2606    17554    family pk_family 
   CONSTRAINT     K   ALTER TABLE ONLY itv.family
    ADD CONSTRAINT pk_family PRIMARY KEY (id);
 7   ALTER TABLE ONLY itv.family DROP CONSTRAINT pk_family;
       itv            hestia_erp_internal    false    252            ^           2606    17562    gender pk_gender 
   CONSTRAINT     K   ALTER TABLE ONLY itv.gender
    ADD CONSTRAINT pk_gender PRIMARY KEY (id);
 7   ALTER TABLE ONLY itv.gender DROP CONSTRAINT pk_gender;
       itv            hestia_erp_internal    false    253            Z           2606    17546    raw_material pk_raw_material 
   CONSTRAINT     W   ALTER TABLE ONLY itv.raw_material
    ADD CONSTRAINT pk_raw_material PRIMARY KEY (id);
 C   ALTER TABLE ONLY itv.raw_material DROP CONSTRAINT pk_raw_material;
       itv            hestia_erp_internal    false    251            R           2606    17516    size pk_size 
   CONSTRAINT     G   ALTER TABLE ONLY itv.size
    ADD CONSTRAINT pk_size PRIMARY KEY (id);
 3   ALTER TABLE ONLY itv.size DROP CONSTRAINT pk_size;
       itv            hestia_erp_internal    false    247            N           2606    17506     technical_file pk_technical_file 
   CONSTRAINT     [   ALTER TABLE ONLY itv.technical_file
    ADD CONSTRAINT pk_technical_file PRIMARY KEY (id);
 G   ALTER TABLE ONLY itv.technical_file DROP CONSTRAINT pk_technical_file;
       itv            hestia_erp_internal    false    246            X           2606    17538 *   technical_file_size pk_technical_file_size 
   CONSTRAINT     }   ALTER TABLE ONLY itv.technical_file_size
    ADD CONSTRAINT pk_technical_file_size PRIMARY KEY (technical_file_id, size_id);
 Q   ALTER TABLE ONLY itv.technical_file_size DROP CONSTRAINT pk_technical_file_size;
       itv            hestia_erp_internal    false    250    250            P           2606    17508 %   technical_file uq_technical_file_code 
   CONSTRAINT     i   ALTER TABLE ONLY itv.technical_file
    ADD CONSTRAINT uq_technical_file_code UNIQUE (company_id, code);
 L   ALTER TABLE ONLY itv.technical_file DROP CONSTRAINT uq_technical_file_code;
       itv            hestia_erp_internal    false    246    246            .           2606    17366    document pk_document 
   CONSTRAINT     Q   ALTER TABLE ONLY sales.document
    ADD CONSTRAINT pk_document PRIMARY KEY (id);
 =   ALTER TABLE ONLY sales.document DROP CONSTRAINT pk_document;
       sales            hestia_erp_internal    false    234            >           2606    17459 #   doccument_series pk_document_series 
   CONSTRAINT     `   ALTER TABLE ONLY sales.doccument_series
    ADD CONSTRAINT pk_document_series PRIMARY KEY (id);
 L   ALTER TABLE ONLY sales.doccument_series DROP CONSTRAINT pk_document_series;
       sales            hestia_erp_internal    false    241            H           2606    17494    client pk_supplier 
   CONSTRAINT     O   ALTER TABLE ONLY sales.client
    ADD CONSTRAINT pk_supplier PRIMARY KEY (id);
 ;   ALTER TABLE ONLY sales.client DROP CONSTRAINT pk_supplier;
       sales            postgres    false    245            J           2606    17496    client uq_supplier_code 
   CONSTRAINT     ]   ALTER TABLE ONLY sales.client
    ADD CONSTRAINT uq_supplier_code UNIQUE (company_id, code);
 @   ALTER TABLE ONLY sales.client DROP CONSTRAINT uq_supplier_code;
       sales            postgres    false    245    245            L           2606    17498    client uq_supplier_vat_id 
   CONSTRAINT     a   ALTER TABLE ONLY sales.client
    ADD CONSTRAINT uq_supplier_vat_id UNIQUE (company_id, vat_id);
 B   ALTER TABLE ONLY sales.client DROP CONSTRAINT uq_supplier_vat_id;
       sales            postgres    false    245    245            :           2606    17443    module pk_module 
   CONSTRAINT     U   ALTER TABLE ONLY subscriptions.module
    ADD CONSTRAINT pk_module PRIMARY KEY (id);
 A   ALTER TABLE ONLY subscriptions.module DROP CONSTRAINT pk_module;
       subscriptions            hestia_erp_internal    false    239            <           2606    17450 *   subscription_module pk_subscription_module 
   CONSTRAINT     �   ALTER TABLE ONLY subscriptions.subscription_module
    ADD CONSTRAINT pk_subscription_module PRIMARY KEY (subscription_id, module_id);
 [   ALTER TABLE ONLY subscriptions.subscription_module DROP CONSTRAINT pk_subscription_module;
       subscriptions            hestia_erp_internal    false    240    240            8           2606    17435    subscription pk_subscriptions 
   CONSTRAINT     b   ALTER TABLE ONLY subscriptions.subscription
    ADD CONSTRAINT pk_subscriptions PRIMARY KEY (id);
 N   ALTER TABLE ONLY subscriptions.subscription DROP CONSTRAINT pk_subscriptions;
       subscriptions            hestia_erp_internal    false    238                       2606    17294    users pk_user 
   CONSTRAINT     J   ALTER TABLE ONLY users.users
    ADD CONSTRAINT pk_user PRIMARY KEY (id);
 6   ALTER TABLE ONLY users.users DROP CONSTRAINT pk_user;
       users            hestia_erp_internal    false    224                        2606    17321    user_company pk_user_company 
   CONSTRAINT     j   ALTER TABLE ONLY users.user_company
    ADD CONSTRAINT pk_user_company PRIMARY KEY (user_id, company_id);
 E   ALTER TABLE ONLY users.user_company DROP CONSTRAINT pk_user_company;
       users            hestia_erp_internal    false    227    227            T           2606    17523    users_session pk_users_session 
   CONSTRAINT     [   ALTER TABLE ONLY users.users_session
    ADD CONSTRAINT pk_users_session PRIMARY KEY (id);
 G   ALTER TABLE ONLY users.users_session DROP CONSTRAINT pk_users_session;
       users            postgres    false    248                       2606    18223    users uq_users_email 
   CONSTRAINT     O   ALTER TABLE ONLY users.users
    ADD CONSTRAINT uq_users_email UNIQUE (email);
 =   ALTER TABLE ONLY users.users DROP CONSTRAINT uq_users_email;
       users            hestia_erp_internal    false    224            _           2606    17563    company fk_company_country    FK CONSTRAINT     �   ALTER TABLE ONLY companies.company
    ADD CONSTRAINT fk_company_country FOREIGN KEY (country_code) REFERENCES general.country(code) ON UPDATE RESTRICT ON DELETE RESTRICT;
 G   ALTER TABLE ONLY companies.company DROP CONSTRAINT fk_company_country;
    	   companies          postgres    false    4930    243    225            t           2606    17668    location fk_location_company    FK CONSTRAINT     �   ALTER TABLE ONLY companies.location
    ADD CONSTRAINT fk_location_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 I   ALTER TABLE ONLY companies.location DROP CONSTRAINT fk_location_company;
    	   companies          hestia_erp_internal    false    4888    225    242            u           2606    17673    location fk_location_country    FK CONSTRAINT     �   ALTER TABLE ONLY companies.location
    ADD CONSTRAINT fk_location_country FOREIGN KEY (country_code) REFERENCES general.country(code) ON UPDATE RESTRICT ON DELETE RESTRICT;
 I   ALTER TABLE ONLY companies.location DROP CONSTRAINT fk_location_country;
    	   companies          hestia_erp_internal    false    242    243    4930            v           2606    17678     preference fk_preference_company    FK CONSTRAINT     �   ALTER TABLE ONLY companies.preference
    ADD CONSTRAINT fk_preference_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 M   ALTER TABLE ONLY companies.preference DROP CONSTRAINT fk_preference_company;
    	   companies          postgres    false    4888    244    225            l           2606    17628    device fk_device_company    FK CONSTRAINT     �   ALTER TABLE ONLY hr.device
    ADD CONSTRAINT fk_device_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 >   ALTER TABLE ONLY hr.device DROP CONSTRAINT fk_device_company;
       hr          hestia_erp_internal    false    232    4888    225            `           2606    17568    employee fk_employee_company    FK CONSTRAINT     �   ALTER TABLE ONLY hr.employee
    ADD CONSTRAINT fk_employee_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 B   ALTER TABLE ONLY hr.employee DROP CONSTRAINT fk_employee_company;
       hr          postgres    false    4888    225    226            j           2606    17618 C   employee_identification_doc fk_employee_identification_doc_employee    FK CONSTRAINT     �   ALTER TABLE ONLY hr.employee_identification_doc
    ADD CONSTRAINT fk_employee_identification_doc_employee FOREIGN KEY (employee_id) REFERENCES hr.employee(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 i   ALTER TABLE ONLY hr.employee_identification_doc DROP CONSTRAINT fk_employee_identification_doc_employee;
       hr          hestia_erp_internal    false    226    230    4894            k           2606    17623 M   employee_identification_doc fk_employee_identification_doc_identification_doc    FK CONSTRAINT     �   ALTER TABLE ONLY hr.employee_identification_doc
    ADD CONSTRAINT fk_employee_identification_doc_identification_doc FOREIGN KEY (identification_doc_id) REFERENCES hr.identification_doc(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 s   ALTER TABLE ONLY hr.employee_identification_doc DROP CONSTRAINT fk_employee_identification_doc_identification_doc;
       hr          hestia_erp_internal    false    231    230    4904            a           2606    17573    employee fk_employee_location    FK CONSTRAINT     �   ALTER TABLE ONLY hr.employee
    ADD CONSTRAINT fk_employee_location FOREIGN KEY (location_id) REFERENCES companies.location(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 C   ALTER TABLE ONLY hr.employee DROP CONSTRAINT fk_employee_location;
       hr          postgres    false    4928    242    226            e           2606    17598    timesheet fk_timesheet_device    FK CONSTRAINT     �   ALTER TABLE ONLY hr.timesheet
    ADD CONSTRAINT fk_timesheet_device FOREIGN KEY (device) REFERENCES hr.device(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 C   ALTER TABLE ONLY hr.timesheet DROP CONSTRAINT fk_timesheet_device;
       hr          hestia_erp_internal    false    232    4906    228            f           2606    17593    timesheet fk_timesheet_employee    FK CONSTRAINT     �   ALTER TABLE ONLY hr.timesheet
    ADD CONSTRAINT fk_timesheet_employee FOREIGN KEY (employee_id) REFERENCES hr.employee(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 E   ALTER TABLE ONLY hr.timesheet DROP CONSTRAINT fk_timesheet_employee;
       hr          hestia_erp_internal    false    4894    228    226            g           2606    17603 %   timesheet fk_timesheet_timesheet_type    FK CONSTRAINT     �   ALTER TABLE ONLY hr.timesheet
    ADD CONSTRAINT fk_timesheet_timesheet_type FOREIGN KEY (type) REFERENCES hr.timesheet_type(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 K   ALTER TABLE ONLY hr.timesheet DROP CONSTRAINT fk_timesheet_timesheet_type;
       hr          hestia_erp_internal    false    4908    233    228            h           2606    17608    vacation fk_vacation_employee    FK CONSTRAINT     �   ALTER TABLE ONLY hr.vacation
    ADD CONSTRAINT fk_vacation_employee FOREIGN KEY (employee_id) REFERENCES hr.employee(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 C   ALTER TABLE ONLY hr.vacation DROP CONSTRAINT fk_vacation_employee;
       hr          hestia_erp_internal    false    226    229    4894            i           2606    17613    vacation fk_vacation_employee_1    FK CONSTRAINT     �   ALTER TABLE ONLY hr.vacation
    ADD CONSTRAINT fk_vacation_employee_1 FOREIGN KEY (approved_by) REFERENCES hr.employee(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 E   ALTER TABLE ONLY hr.vacation DROP CONSTRAINT fk_vacation_employee_1;
       hr          hestia_erp_internal    false    4894    226    229            o           2606    17648    category fk_category_category    FK CONSTRAINT     �   ALTER TABLE ONLY inventory.category
    ADD CONSTRAINT fk_category_category FOREIGN KEY (parent_id) REFERENCES inventory.category(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 J   ALTER TABLE ONLY inventory.category DROP CONSTRAINT fk_category_category;
    	   inventory          hestia_erp_internal    false    237    4918    237            p           2606    17643    category fk_category_company    FK CONSTRAINT     �   ALTER TABLE ONLY inventory.category
    ADD CONSTRAINT fk_category_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 I   ALTER TABLE ONLY inventory.category DROP CONSTRAINT fk_category_company;
    	   inventory          hestia_erp_internal    false    225    4888    237            m           2606    17633    product fk_product_category    FK CONSTRAINT     �   ALTER TABLE ONLY inventory.product
    ADD CONSTRAINT fk_product_category FOREIGN KEY (category_id) REFERENCES inventory.category(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 H   ALTER TABLE ONLY inventory.product DROP CONSTRAINT fk_product_category;
    	   inventory          hestia_erp_internal    false    236    237    4918            n           2606    17638    product fk_product_company    FK CONSTRAINT     �   ALTER TABLE ONLY inventory.product
    ADD CONSTRAINT fk_product_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 G   ALTER TABLE ONLY inventory.product DROP CONSTRAINT fk_product_company;
    	   inventory          hestia_erp_internal    false    4888    236    225            ~           2606    17718     department fk_department_company    FK CONSTRAINT     �   ALTER TABLE ONLY itv.department
    ADD CONSTRAINT fk_department_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 G   ALTER TABLE ONLY itv.department DROP CONSTRAINT fk_department_company;
       itv          hestia_erp_internal    false    249    225    4888            �           2606    17738    family fk_family_company    FK CONSTRAINT     �   ALTER TABLE ONLY itv.family
    ADD CONSTRAINT fk_family_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 ?   ALTER TABLE ONLY itv.family DROP CONSTRAINT fk_family_company;
       itv          hestia_erp_internal    false    252    4888    225            �           2606    17743    gender fk_gender_company    FK CONSTRAINT     �   ALTER TABLE ONLY itv.gender
    ADD CONSTRAINT fk_gender_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 ?   ALTER TABLE ONLY itv.gender DROP CONSTRAINT fk_gender_company;
       itv          hestia_erp_internal    false    253    4888    225            �           2606    17733 $   raw_material fk_raw_material_company    FK CONSTRAINT     �   ALTER TABLE ONLY itv.raw_material
    ADD CONSTRAINT fk_raw_material_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 K   ALTER TABLE ONLY itv.raw_material DROP CONSTRAINT fk_raw_material_company;
       itv          hestia_erp_internal    false    4888    251    225            |           2606    17708    size fk_size_company    FK CONSTRAINT     �   ALTER TABLE ONLY itv.size
    ADD CONSTRAINT fk_size_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 ;   ALTER TABLE ONLY itv.size DROP CONSTRAINT fk_size_company;
       itv          hestia_erp_internal    false    225    4888    247            y           2606    17698 '   technical_file fk_technical_file_client    FK CONSTRAINT     �   ALTER TABLE ONLY itv.technical_file
    ADD CONSTRAINT fk_technical_file_client FOREIGN KEY (client_id) REFERENCES sales.client(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 N   ALTER TABLE ONLY itv.technical_file DROP CONSTRAINT fk_technical_file_client;
       itv          hestia_erp_internal    false    4936    246    245            z           2606    17693 (   technical_file fk_technical_file_company    FK CONSTRAINT     �   ALTER TABLE ONLY itv.technical_file
    ADD CONSTRAINT fk_technical_file_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 O   ALTER TABLE ONLY itv.technical_file DROP CONSTRAINT fk_technical_file_company;
       itv          hestia_erp_internal    false    246    4888    225            {           2606    17703 +   technical_file fk_technical_file_department    FK CONSTRAINT     �   ALTER TABLE ONLY itv.technical_file
    ADD CONSTRAINT fk_technical_file_department FOREIGN KEY (department_id) REFERENCES itv.department(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 R   ALTER TABLE ONLY itv.technical_file DROP CONSTRAINT fk_technical_file_department;
       itv          hestia_erp_internal    false    249    246    4950                       2606    17728 /   technical_file_size fk_technical_file_size_size    FK CONSTRAINT     �   ALTER TABLE ONLY itv.technical_file_size
    ADD CONSTRAINT fk_technical_file_size_size FOREIGN KEY (size_id) REFERENCES itv.size(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 V   ALTER TABLE ONLY itv.technical_file_size DROP CONSTRAINT fk_technical_file_size_size;
       itv          hestia_erp_internal    false    247    4946    250            �           2606    17723 9   technical_file_size fk_technical_file_size_technical_file    FK CONSTRAINT     �   ALTER TABLE ONLY itv.technical_file_size
    ADD CONSTRAINT fk_technical_file_size_technical_file FOREIGN KEY (technical_file_id) REFERENCES itv.technical_file(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 `   ALTER TABLE ONLY itv.technical_file_size DROP CONSTRAINT fk_technical_file_size_technical_file;
       itv          hestia_erp_internal    false    250    246    4942            w           2606    17683    client fk_supplier_company    FK CONSTRAINT     �   ALTER TABLE ONLY sales.client
    ADD CONSTRAINT fk_supplier_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 C   ALTER TABLE ONLY sales.client DROP CONSTRAINT fk_supplier_company;
       sales          postgres    false    225    4888    245            x           2606    17688    client fk_supplier_country    FK CONSTRAINT     �   ALTER TABLE ONLY sales.client
    ADD CONSTRAINT fk_supplier_country FOREIGN KEY (country) REFERENCES general.country(code) ON UPDATE RESTRICT ON DELETE RESTRICT;
 C   ALTER TABLE ONLY sales.client DROP CONSTRAINT fk_supplier_country;
       sales          postgres    false    4930    245    243            q           2606    17653 $   subscription fk_subscription_company    FK CONSTRAINT     �   ALTER TABLE ONLY subscriptions.subscription
    ADD CONSTRAINT fk_subscription_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 U   ALTER TABLE ONLY subscriptions.subscription DROP CONSTRAINT fk_subscription_company;
       subscriptions          hestia_erp_internal    false    238    225    4888            r           2606    17663 1   subscription_module fk_subscription_module_module    FK CONSTRAINT     �   ALTER TABLE ONLY subscriptions.subscription_module
    ADD CONSTRAINT fk_subscription_module_module FOREIGN KEY (module_id) REFERENCES subscriptions.module(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 b   ALTER TABLE ONLY subscriptions.subscription_module DROP CONSTRAINT fk_subscription_module_module;
       subscriptions          hestia_erp_internal    false    4922    240    239            s           2606    17658 7   subscription_module fk_subscription_module_subscription    FK CONSTRAINT     �   ALTER TABLE ONLY subscriptions.subscription_module
    ADD CONSTRAINT fk_subscription_module_subscription FOREIGN KEY (subscription_id) REFERENCES subscriptions.subscription(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 h   ALTER TABLE ONLY subscriptions.subscription_module DROP CONSTRAINT fk_subscription_module_subscription;
       subscriptions          hestia_erp_internal    false    238    240    4920            b           2606    17583 $   user_company fk_user_company_company    FK CONSTRAINT     �   ALTER TABLE ONLY users.user_company
    ADD CONSTRAINT fk_user_company_company FOREIGN KEY (company_id) REFERENCES companies.company(id) MATCH FULL ON UPDATE RESTRICT ON DELETE RESTRICT;
 M   ALTER TABLE ONLY users.user_company DROP CONSTRAINT fk_user_company_company;
       users          hestia_erp_internal    false    227    4888    225            c           2606    17588 %   user_company fk_user_company_employee    FK CONSTRAINT     �   ALTER TABLE ONLY users.user_company
    ADD CONSTRAINT fk_user_company_employee FOREIGN KEY (employee_id) REFERENCES hr.employee(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 N   ALTER TABLE ONLY users.user_company DROP CONSTRAINT fk_user_company_employee;
       users          hestia_erp_internal    false    227    4894    226            d           2606    17578 "   user_company fk_user_company_users    FK CONSTRAINT     �   ALTER TABLE ONLY users.user_company
    ADD CONSTRAINT fk_user_company_users FOREIGN KEY (user_id) REFERENCES users.users(id) MATCH FULL ON UPDATE RESTRICT ON DELETE RESTRICT;
 K   ALTER TABLE ONLY users.user_company DROP CONSTRAINT fk_user_company_users;
       users          hestia_erp_internal    false    4884    224    227            }           2606    17713 $   users_session fk_users_session_users    FK CONSTRAINT     �   ALTER TABLE ONLY users.users_session
    ADD CONSTRAINT fk_users_session_users FOREIGN KEY (user_id) REFERENCES users.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 M   ALTER TABLE ONLY users.users_session DROP CONSTRAINT fk_users_session_users;
       users          postgres    false    4884    248    224               i   x�3K�H31MJյ0�H�5143ҵL3I�M�4�0N25M1�H�I-.Q�qq�t��M-J�L�ᴄμҜΠ�DCN�Ģ�Ԝ�bNsS]C΀�=... �       %      x������ � �      '      x������ � �      &      x���/*)MO��,����� <+            x������ � �            x������ � �            x������ � �            x������ � �            x������ � �            x������ � �            x������ � �             x������ � �            x������ � �            x������ � �      ,      x������ � �      /      x������ � �      0      x������ � �      .      x������ � �      *      x������ � �      )      x������ � �      -      x������ � �      (      x������ � �      $      x������ � �            x������ � �      "      x������ � �      !      x������ � �      #      x������ � �         L   x����  ���B��~���#4�9���J�����Y�$Q��]�(0�&�����ݶ߾���J��         �   x�͏1N1Ek�)r'���]
H��ό���l�Y��=�4��������k�Xm�Rm��Y%�(�&vh��6�u����+�M��2b�M�r[���g�Ll��b#��N�h﹧ԁ
3t
`������9�	����a/c<�<���5��Mz��VN9	{L��J�%Q���.w=�Ώ�ܦQ�\����m�!Zdj���XȜ���㡫���,�?��?�����v�>��4�&�B      +   u  x���9��HD��S�?��=I�e�\���WuY��!��[�g暠���	�ɠ��c��$�E�D�<���H��#�ַM�_������o�[څA��A��:z*��U!Q!��|l/�_a�f��	%?��H'C�`��?��٥J7�)�-G�8MV��Bt}�N!�zZ�ҫ�ȷ�(��?Bo�m�k�<��)���9{�|�1F�G��Vu��M�8�O��yI?t(�='�P���qن��������O���쌝,o9~��7�pfI?3l�ڍ!�{��w'��ki���ne�����;�<��8F�}�W�0y���t֘N��Zr�҇ÑW{��mr9���Y�Df�㙟]*�V�k���������$|�&g�j���@w��sh����y[^�H��yK�Պ�c�����E�-x�^͒��'���pl��7!Ug�F��^�~��7�������3�ch5���#�J��7E�����P�o���f,���s`lO�5��Zg�[�S���i�5?�������Z�}�)�6|�V�]B��	v�Yo������W������/B�d���m��!`�ڠ:�Lݥ�g:ڔic���؍qymB������?_+�9     