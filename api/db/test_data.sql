PGDMP  &                	    |           erp    16.4    16.4 �    z           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            {           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            |           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            }           1262    37136    erp    DATABASE     |   CREATE DATABASE erp WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Portuguese_Portugal.1252';
    DROP DATABASE erp;
                hestia_erp_internal    false                        2615    37609 
   accounting    SCHEMA        CREATE SCHEMA accounting;
    DROP SCHEMA accounting;
                hestia_erp_internal    false                        2615    37137 	   companies    SCHEMA        CREATE SCHEMA companies;
    DROP SCHEMA companies;
                hestia_erp_internal    false                        2615    37138    general    SCHEMA        CREATE SCHEMA general;
    DROP SCHEMA general;
                hestia_erp_internal    false            	            2615    37139    hr    SCHEMA        CREATE SCHEMA hr;
    DROP SCHEMA hr;
                hestia_erp_internal    false            
            2615    37140 	   inventory    SCHEMA        CREATE SCHEMA inventory;
    DROP SCHEMA inventory;
                hestia_erp_internal    false                        2615    37141    itv    SCHEMA        CREATE SCHEMA itv;
    DROP SCHEMA itv;
                hestia_erp_internal    false                        2615    37142    sales    SCHEMA        CREATE SCHEMA sales;
    DROP SCHEMA sales;
                hestia_erp_internal    false                        2615    37143    subscriptions    SCHEMA        CREATE SCHEMA subscriptions;
    DROP SCHEMA subscriptions;
                hestia_erp_internal    false                        2615    37144    users    SCHEMA        CREATE SCHEMA users;
    DROP SCHEMA users;
                hestia_erp_internal    false                        3079    37145    pgcrypto 	   EXTENSION     <   CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;
    DROP EXTENSION pgcrypto;
                   false            ~           0    0    EXTENSION pgcrypto    COMMENT     <   COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';
                        false    2                       1259    37662    vat_exemption    TABLE     �   CREATE TABLE accounting.vat_exemption (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    code text NOT NULL,
    description smallint NOT NULL,
    start_date timestamp with time zone NOT NULL,
    end_date timestamp with time zone
);
 %   DROP TABLE accounting.vat_exemption;
    
   accounting         heap    hestia_erp_internal    false    2    15                       0    0    TABLE vat_exemption    COMMENT     Q   COMMENT ON TABLE accounting.vat_exemption IS 'Exemption reasons provided by AT';
       
   accounting          hestia_erp_internal    false    259                       1259    37645    vat_rate    TABLE       CREATE TABLE accounting.vat_rate (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    start_date timestamp with time zone NOT NULL,
    end_date timestamp with time zone,
    rate integer NOT NULL,
    vat_type_id uuid NOT NULL,
    country_id uuid NOT NULL
);
     DROP TABLE accounting.vat_rate;
    
   accounting         heap    hestia_erp_internal    false    2    15                       1259    37670    vat_type    TABLE     �   CREATE TABLE accounting.vat_type (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    code text NOT NULL,
    description text
);
     DROP TABLE accounting.vat_type;
    
   accounting         heap    hestia_erp_internal    false    2    15            �           0    0    TABLE vat_type    COMMENT     W   COMMENT ON TABLE accounting.vat_type IS 'Column to define type such as
NOR, INT, RED';
       
   accounting          hestia_erp_internal    false    260            �            1259    37182    company    TABLE     )  CREATE TABLE companies.company (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    commercial_name text,
    vat_id text NOT NULL,
    ssn text NOT NULL,
    street text NOT NULL,
    locality text NOT NULL,
    postal_code text NOT NULL,
    country_id uuid NOT NULL
);
    DROP TABLE companies.company;
    	   companies         heap    postgres    false    7            �           0    0    COLUMN company.id    COMMENT     ?   COMMENT ON COLUMN companies.company.id IS 'Company ID (UUID)';
       	   companies          postgres    false    225            �           0    0    COLUMN company.commercial_name    COMMENT     e   COMMENT ON COLUMN companies.company.commercial_name IS 'Commercial name like Apple, Google, etc...';
       	   companies          postgres    false    225            �            1259    37188    location    TABLE     �   CREATE TABLE companies.location (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    name text NOT NULL,
    street text NOT NULL,
    locality text NOT NULL,
    postal_code text NOT NULL,
    country_id uuid NOT NULL
);
    DROP TABLE companies.location;
    	   companies         heap    hestia_erp_internal    false    7            �           0    0    COLUMN location.name    COMMENT     >   COMMENT ON COLUMN companies.location.name IS 'Location Name';
       	   companies          hestia_erp_internal    false    226            �            1259    37195 
   preference    TABLE     �   CREATE TABLE companies.preference (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    code text NOT NULL,
    value text NOT NULL
);
 !   DROP TABLE companies.preference;
    	   companies         heap    postgres    false    7            �            1259    37201    country    TABLE     �   CREATE TABLE general.country (
    code character varying(2) NOT NULL,
    name text NOT NULL,
    active boolean DEFAULT true NOT NULL,
    id uuid DEFAULT public.gen_random_uuid() NOT NULL
);
    DROP TABLE general.country;
       general         heap    hestia_erp_internal    false    2    8            �           0    0    COLUMN country.code    COMMENT     E   COMMENT ON COLUMN general.country.code IS 'ISO 3166-1 alpha-2 code';
          general          hestia_erp_internal    false    228            �           0    0    COLUMN country.name    COMMENT     E   COMMENT ON COLUMN general.country.name IS 'English name of country';
          general          hestia_erp_internal    false    228            �           0    0    COLUMN country.active    COMMENT     d   COMMENT ON COLUMN general.country.active IS 'Is the Country active? Say the country was abolished';
          general          hestia_erp_internal    false    228            �            1259    37207    device    TABLE       CREATE TABLE hr.device (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid,
    active boolean DEFAULT true NOT NULL,
    name character varying(70) NOT NULL,
    key uuid DEFAULT gen_random_uuid() NOT NULL,
    last_date timestamp with time zone NOT NULL
);
    DROP TABLE hr.device;
       hr         heap    hestia_erp_internal    false    9            �            1259    37213    employee    TABLE     �   CREATE TABLE hr.employee (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    name text NOT NULL,
    location_id uuid NOT NULL
);
    DROP TABLE hr.employee;
       hr         heap    postgres    false    9            �           0    0    COLUMN employee.id    COMMENT     B   COMMENT ON COLUMN hr.employee.id IS 'HR''s employee internal id';
          hr          postgres    false    230            �           0    0    COLUMN employee.location_id    COMMENT     I   COMMENT ON COLUMN hr.employee.location_id IS 'Does it need to be NULL?';
          hr          postgres    false    230            �            1259    37219    employee_identification_doc    TABLE     �   CREATE TABLE hr.employee_identification_doc (
    employee_id uuid NOT NULL,
    identification_doc_id uuid NOT NULL,
    doc_number character varying(50) NOT NULL
);
 +   DROP TABLE hr.employee_identification_doc;
       hr         heap    hestia_erp_internal    false    9            �            1259    37222    identification_doc    TABLE     �   CREATE TABLE hr.identification_doc (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    code character varying(20),
    name text NOT NULL,
    description text
);
 "   DROP TABLE hr.identification_doc;
       hr         heap    hestia_erp_internal    false    9            �            1259    37228 	   timesheet    TABLE     �   CREATE TABLE hr.timesheet (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    employee_id uuid NOT NULL,
    device uuid,
    type uuid NOT NULL,
    date timestamp with time zone NOT NULL
);
    DROP TABLE hr.timesheet;
       hr         heap    hestia_erp_internal    false    9            �            1259    37232    timesheet_type    TABLE     �   CREATE TABLE hr.timesheet_type (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    code character varying(20) NOT NULL,
    name character varying(100)
);
    DROP TABLE hr.timesheet_type;
       hr         heap    hestia_erp_internal    false    9            �            1259    37236    vacation    TABLE     �   CREATE TABLE hr.vacation (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    employee_id uuid NOT NULL,
    start_date timestamp with time zone NOT NULL,
    end_date timestamp with time zone NOT NULL,
    approved_by uuid
);
    DROP TABLE hr.vacation;
       hr         heap    hestia_erp_internal    false    9            �            1259    37240    category    TABLE     �   CREATE TABLE inventory.category (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    name character varying(40) NOT NULL,
    parent_id uuid
);
    DROP TABLE inventory.category;
    	   inventory         heap    hestia_erp_internal    false    10            �            1259    37244    movement    TABLE     T   CREATE TABLE inventory.movement (
    id uuid DEFAULT gen_random_uuid() NOT NULL
);
    DROP TABLE inventory.movement;
    	   inventory         heap    hestia_erp_internal    false    10            �            1259    37248    product    TABLE     L  CREATE TABLE inventory.product (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    active boolean DEFAULT true NOT NULL,
    category_id uuid NOT NULL,
    company_id uuid NOT NULL,
    name text NOT NULL,
    reference character varying(40) NOT NULL,
    barcode character varying(13),
    description text,
    stock numeric
);
    DROP TABLE inventory.product;
    	   inventory         heap    hestia_erp_internal    false    10            �           0    0    COLUMN product.barcode    COMMENT     H   COMMENT ON COLUMN inventory.product.barcode IS 'Barcode type: EAN-13A';
       	   inventory          hestia_erp_internal    false    238            �           0    0    COLUMN product.stock    COMMENT     v   COMMENT ON COLUMN inventory.product.stock IS 'If stock is null then stock is not enabled
100 digits good enough no?';
       	   inventory          hestia_erp_internal    false    238            �            1259    37255 
   department    TABLE     �   CREATE TABLE itv.department (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    name text NOT NULL
);
    DROP TABLE itv.department;
       itv         heap    hestia_erp_internal    false    11            �            1259    37261    family    TABLE     �   CREATE TABLE itv.family (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    name text NOT NULL
);
    DROP TABLE itv.family;
       itv         heap    hestia_erp_internal    false    11            �           0    0    TABLE family    COMMENT     c   COMMENT ON TABLE itv.family IS 'Family of the product (like Pijamas, Sweater, Sweater w/ Hoodie)';
          itv          hestia_erp_internal    false    240            �            1259    37267    gender    TABLE     �   CREATE TABLE itv.gender (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    gender text NOT NULL
);
    DROP TABLE itv.gender;
       itv         heap    hestia_erp_internal    false    11            �            1259    37273    raw_material    TABLE     �   CREATE TABLE itv.raw_material (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    code text,
    material text NOT NULL
);
    DROP TABLE itv.raw_material;
       itv         heap    hestia_erp_internal    false    11            �            1259    37279    size    TABLE     �   CREATE TABLE itv.size (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    size text NOT NULL
);
    DROP TABLE itv.size;
       itv         heap    hestia_erp_internal    false    11            �           0    0    COLUMN size.size    COMMENT     B   COMMENT ON COLUMN itv.size.size IS 'Size (XS, S, M, L, XL,....)';
          itv          hestia_erp_internal    false    243            �            1259    37285    technical_file    TABLE     �   CREATE TABLE itv.technical_file (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    code text NOT NULL,
    client_id uuid NOT NULL,
    department_id uuid NOT NULL
);
    DROP TABLE itv.technical_file;
       itv         heap    hestia_erp_internal    false    11            �            1259    37291    technical_file_size    TABLE     {   CREATE TABLE itv.technical_file_size (
    technical_file_id uuid NOT NULL,
    size_id uuid NOT NULL,
    comment text
);
 $   DROP TABLE itv.technical_file_size;
       itv         heap    hestia_erp_internal    false    11            �            1259    37296    client    TABLE     )  CREATE TABLE sales.client (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    name text NOT NULL,
    code text NOT NULL,
    vat_id text NOT NULL,
    street text NOT NULL,
    postal_code text NOT NULL,
    locality text NOT NULL,
    country_id uuid NOT NULL
);
    DROP TABLE sales.client;
       sales         heap    postgres    false    12            �            1259    37309    document    TABLE     U  CREATE TABLE sales.document (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    series_id uuid NOT NULL,
    num integer NOT NULL,
    date timestamp with time zone NOT NULL,
    year integer GENERATED ALWAYS AS (EXTRACT(year FROM (date AT TIME ZONE 'Europe/Lisbon'::text))) STORED,
    hash text NOT NULL
);
    DROP TABLE sales.document;
       sales         heap    hestia_erp_internal    false    12            �           0    0    COLUMN document.num    COMMENT     A   COMMENT ON COLUMN sales.document.num IS 'number of the invoice';
          sales          hestia_erp_internal    false    247                       1259    37653    document_line    TABLE     �  CREATE TABLE sales.document_line (
    document_id uuid NOT NULL,
    product_id uuid NOT NULL,
    vat_rate_id uuid,
    vat_exemption_id uuid,
    quantity numeric NOT NULL,
    price numeric NOT NULL,
    net_total numeric GENERATED ALWAYS AS ((quantity * price)) STORED,
    gross_total numeric NOT NULL,
    CONSTRAINT ck_document_line_vat_rate_vat_exemption CHECK ((((vat_rate_id IS NOT NULL) AND (vat_exemption_id IS NULL)) OR ((vat_rate_id IS NULL) AND (vat_exemption_id IS NOT NULL))))
);
     DROP TABLE sales.document_line;
       sales         heap    hestia_erp_internal    false    12            �           0    0    COLUMN document_line.price    COMMENT     A   COMMENT ON COLUMN sales.document_line.price IS 'price per unit';
          sales          hestia_erp_internal    false    258            �           0    0     COLUMN document_line.gross_total    COMMENT     ^   COMMENT ON COLUMN sales.document_line.gross_total IS 'Needs to be calculated by the backend';
          sales          hestia_erp_internal    false    258            �            1259    37617    series    TABLE     �   CREATE TABLE sales.series (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    code text DEFAULT 'H'::text NOT NULL,
    description text
);
    DROP TABLE sales.series;
       sales         heap    hestia_erp_internal    false    2    12            �           0    0    COLUMN series.code    COMMENT     9   COMMENT ON COLUMN sales.series.code IS 'FT code/number';
          sales          hestia_erp_internal    false    254                        1259    37638    series_type_code    TABLE     x   CREATE TABLE sales.series_type_code (
    series_id uuid NOT NULL,
    type_id uuid NOT NULL,
    code text NOT NULL
);
 #   DROP TABLE sales.series_type_code;
       sales         heap    hestia_erp_internal    false    12            �           0    0    COLUMN series_type_code.code    COMMENT     S   COMMENT ON COLUMN sales.series_type_code.code IS 'Validation code provided by AT';
          sales          hestia_erp_internal    false    256            �            1259    37628    type    TABLE     �   CREATE TABLE sales.type (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    company_id uuid,
    code text NOT NULL,
    type uuid
);
    DROP TABLE sales.type;
       sales         heap    hestia_erp_internal    false    2    12            �           0    0    COLUMN type.type    COMMENT     M   COMMENT ON COLUMN sales.type.type IS 'Example:
type: FTR
FTR is of type FT';
          sales          hestia_erp_internal    false    255            �            1259    37313    module    TABLE     �   CREATE TABLE subscriptions.module (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    code character varying(20) NOT NULL,
    name text NOT NULL,
    default_price numeric NOT NULL
);
 !   DROP TABLE subscriptions.module;
       subscriptions         heap    hestia_erp_internal    false    13            �            1259    37319    subscription    TABLE     �  CREATE TABLE subscriptions.subscription (
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
       subscriptions         heap    hestia_erp_internal    false    13            �            1259    37326    subscription_module    TABLE     �   CREATE TABLE subscriptions.subscription_module (
    subscription_id uuid NOT NULL,
    module_id uuid NOT NULL,
    price numeric NOT NULL
);
 .   DROP TABLE subscriptions.subscription_module;
       subscriptions         heap    hestia_erp_internal    false    13            �            1259    37331    user_company    TABLE     s   CREATE TABLE users.user_company (
    user_id uuid NOT NULL,
    company_id uuid NOT NULL,
    employee_id uuid
);
    DROP TABLE users.user_company;
       users         heap    hestia_erp_internal    false    14            �            1259    37334    users    TABLE     <  CREATE TABLE users.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    salt text NOT NULL,
    timezone text DEFAULT 'Etc/GMT'::text NOT NULL,
    CONSTRAINT ck_users_timezone CHECK (((now() AT TIME ZONE timezone) IS NOT NULL))
);
    DROP TABLE users.users;
       users         heap    hestia_erp_internal    false    14            �           0    0    TABLE users    COMMENT     j   COMMENT ON TABLE users.users IS 'This is the only table named in plural because USER is a reserved word';
          users          hestia_erp_internal    false    252            �           0    0    COLUMN users.id    COMMENT     L   COMMENT ON COLUMN users.users.id IS 'User''s id (uuid) generated randomly';
          users          hestia_erp_internal    false    252            �           0    0    COLUMN users.name    COMMENT     �   COMMENT ON COLUMN users.users.name IS 'This is the user full name used in the Subscription Management Portal, if used in atual ERP it should use the employee name, if employee name is not available use the user name';
          users          hestia_erp_internal    false    252            �           0    0    COLUMN users.email    COMMENT     8   COMMENT ON COLUMN users.users.email IS 'User''s email';
          users          hestia_erp_internal    false    252            �           0    0    COLUMN users.password    COMMENT     J   COMMENT ON COLUMN users.users.password IS 'User''s password (in SHA512)';
          users          hestia_erp_internal    false    252            �           0    0    COLUMN users.salt    COMMENT     ]   COMMENT ON COLUMN users.users.salt IS 'User''s password salt in string format (maybe hex?)';
          users          hestia_erp_internal    false    252            �           0    0    COLUMN users.timezone    COMMENT     >   COMMENT ON COLUMN users.users.timezone IS 'User''s timezone';
          users          hestia_erp_internal    false    252            �           0    0 %   CONSTRAINT ck_users_timezone ON users    COMMENT     d   COMMENT ON CONSTRAINT ck_users_timezone ON users.users IS 'Check if the inputed timezone is valid';
          users          hestia_erp_internal    false    252            �            1259    37342    users_session    TABLE     �   CREATE TABLE users.users_session (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    expiry_date timestamp with time zone DEFAULT (now() + '08:00:00'::interval) NOT NULL
);
     DROP TABLE users.users_session;
       users         heap    postgres    false    14            v          0    37662    vat_exemption 
   TABLE DATA           X   COPY accounting.vat_exemption (id, code, description, start_date, end_date) FROM stdin;
 
   accounting          hestia_erp_internal    false    259   �      t          0    37645    vat_rate 
   TABLE DATA           _   COPY accounting.vat_rate (id, start_date, end_date, rate, vat_type_id, country_id) FROM stdin;
 
   accounting          hestia_erp_internal    false    257   �      w          0    37670    vat_type 
   TABLE DATA           =   COPY accounting.vat_type (id, code, description) FROM stdin;
 
   accounting          hestia_erp_internal    false    260   �      T          0    37182    company 
   TABLE DATA           w   COPY companies.company (id, name, commercial_name, vat_id, ssn, street, locality, postal_code, country_id) FROM stdin;
 	   companies          postgres    false    225   �      U          0    37188    location 
   TABLE DATA           f   COPY companies.location (id, company_id, name, street, locality, postal_code, country_id) FROM stdin;
 	   companies          hestia_erp_internal    false    226   p      V          0    37195 
   preference 
   TABLE DATA           D   COPY companies.preference (id, company_id, code, value) FROM stdin;
 	   companies          postgres    false    227   �      W          0    37201    country 
   TABLE DATA           :   COPY general.country (code, name, active, id) FROM stdin;
    general          hestia_erp_internal    false    228   �      X          0    37207    device 
   TABLE DATA           J   COPY hr.device (id, company_id, active, name, key, last_date) FROM stdin;
    hr          hestia_erp_internal    false    229   �      Y          0    37213    employee 
   TABLE DATA           A   COPY hr.employee (id, company_id, name, location_id) FROM stdin;
    hr          postgres    false    230         Z          0    37219    employee_identification_doc 
   TABLE DATA           a   COPY hr.employee_identification_doc (employee_id, identification_doc_id, doc_number) FROM stdin;
    hr          hestia_erp_internal    false    231   4      [          0    37222    identification_doc 
   TABLE DATA           E   COPY hr.identification_doc (id, code, name, description) FROM stdin;
    hr          hestia_erp_internal    false    232   Q      \          0    37228 	   timesheet 
   TABLE DATA           D   COPY hr.timesheet (id, employee_id, device, type, date) FROM stdin;
    hr          hestia_erp_internal    false    233   n      ]          0    37232    timesheet_type 
   TABLE DATA           4   COPY hr.timesheet_type (id, code, name) FROM stdin;
    hr          hestia_erp_internal    false    234   �      ^          0    37236    vacation 
   TABLE DATA           R   COPY hr.vacation (id, employee_id, start_date, end_date, approved_by) FROM stdin;
    hr          hestia_erp_internal    false    235   �      _          0    37240    category 
   TABLE DATA           F   COPY inventory.category (id, company_id, name, parent_id) FROM stdin;
 	   inventory          hestia_erp_internal    false    236   �      `          0    37244    movement 
   TABLE DATA           )   COPY inventory.movement (id) FROM stdin;
 	   inventory          hestia_erp_internal    false    237   �      a          0    37248    product 
   TABLE DATA           w   COPY inventory.product (id, active, category_id, company_id, name, reference, barcode, description, stock) FROM stdin;
 	   inventory          hestia_erp_internal    false    238   �      b          0    37255 
   department 
   TABLE DATA           7   COPY itv.department (id, company_id, name) FROM stdin;
    itv          hestia_erp_internal    false    239         c          0    37261    family 
   TABLE DATA           3   COPY itv.family (id, company_id, name) FROM stdin;
    itv          hestia_erp_internal    false    240   9      d          0    37267    gender 
   TABLE DATA           5   COPY itv.gender (id, company_id, gender) FROM stdin;
    itv          hestia_erp_internal    false    241   V      e          0    37273    raw_material 
   TABLE DATA           C   COPY itv.raw_material (id, company_id, code, material) FROM stdin;
    itv          hestia_erp_internal    false    242   s      f          0    37279    size 
   TABLE DATA           1   COPY itv.size (id, company_id, size) FROM stdin;
    itv          hestia_erp_internal    false    243   �      g          0    37285    technical_file 
   TABLE DATA           U   COPY itv.technical_file (id, company_id, code, client_id, department_id) FROM stdin;
    itv          hestia_erp_internal    false    244   �      h          0    37291    technical_file_size 
   TABLE DATA           O   COPY itv.technical_file_size (technical_file_id, size_id, comment) FROM stdin;
    itv          hestia_erp_internal    false    245   �      i          0    37296    client 
   TABLE DATA           n   COPY sales.client (id, company_id, name, code, vat_id, street, postal_code, locality, country_id) FROM stdin;
    sales          postgres    false    246   �      j          0    37309    document 
   TABLE DATA           M   COPY sales.document (id, company_id, series_id, num, date, hash) FROM stdin;
    sales          hestia_erp_internal    false    247         u          0    37653    document_line 
   TABLE DATA           |   COPY sales.document_line (document_id, product_id, vat_rate_id, vat_exemption_id, quantity, price, gross_total) FROM stdin;
    sales          hestia_erp_internal    false    258   !      q          0    37617    series 
   TABLE DATA           B   COPY sales.series (id, company_id, code, description) FROM stdin;
    sales          hestia_erp_internal    false    254   >      s          0    37638    series_type_code 
   TABLE DATA           C   COPY sales.series_type_code (series_id, type_id, code) FROM stdin;
    sales          hestia_erp_internal    false    256   [      r          0    37628    type 
   TABLE DATA           9   COPY sales.type (id, company_id, code, type) FROM stdin;
    sales          hestia_erp_internal    false    255   x      k          0    37313    module 
   TABLE DATA           F   COPY subscriptions.module (id, code, name, default_price) FROM stdin;
    subscriptions          hestia_erp_internal    false    248   �      l          0    37319    subscription 
   TABLE DATA           �   COPY subscriptions.subscription (id, company_id, user_limit, user_price, start_date, end_date, creation_date, update_date) FROM stdin;
    subscriptions          hestia_erp_internal    false    249   �      m          0    37326    subscription_module 
   TABLE DATA           W   COPY subscriptions.subscription_module (subscription_id, module_id, price) FROM stdin;
    subscriptions          hestia_erp_internal    false    250   �      n          0    37331    user_company 
   TABLE DATA           G   COPY users.user_company (user_id, company_id, employee_id) FROM stdin;
    users          hestia_erp_internal    false    251   �      o          0    37334    users 
   TABLE DATA           I   COPY users.users (id, name, email, password, salt, timezone) FROM stdin;
    users          hestia_erp_internal    false    252   H      p          0    37342    users_session 
   TABLE DATA           @   COPY users.users_session (id, user_id, expiry_date) FROM stdin;
    users          postgres    false    253   m      �           2606    37669    vat_exemption pk_vat_exemption 
   CONSTRAINT     `   ALTER TABLE ONLY accounting.vat_exemption
    ADD CONSTRAINT pk_vat_exemption PRIMARY KEY (id);
 L   ALTER TABLE ONLY accounting.vat_exemption DROP CONSTRAINT pk_vat_exemption;
    
   accounting            hestia_erp_internal    false    259            �           2606    37652    vat_rate pk_vat_rate 
   CONSTRAINT     V   ALTER TABLE ONLY accounting.vat_rate
    ADD CONSTRAINT pk_vat_rate PRIMARY KEY (id);
 B   ALTER TABLE ONLY accounting.vat_rate DROP CONSTRAINT pk_vat_rate;
    
   accounting            hestia_erp_internal    false    257            �           2606    37677    vat_type pk_vat_type 
   CONSTRAINT     V   ALTER TABLE ONLY accounting.vat_type
    ADD CONSTRAINT pk_vat_type PRIMARY KEY (id);
 B   ALTER TABLE ONLY accounting.vat_type DROP CONSTRAINT pk_vat_type;
    
   accounting            hestia_erp_internal    false    260            4           2606    37348    company pk_company_id 
   CONSTRAINT     V   ALTER TABLE ONLY companies.company
    ADD CONSTRAINT pk_company_id PRIMARY KEY (id);
 B   ALTER TABLE ONLY companies.company DROP CONSTRAINT pk_company_id;
    	   companies            postgres    false    225            :           2606    37350    location pk_location 
   CONSTRAINT     U   ALTER TABLE ONLY companies.location
    ADD CONSTRAINT pk_location PRIMARY KEY (id);
 A   ALTER TABLE ONLY companies.location DROP CONSTRAINT pk_location;
    	   companies            hestia_erp_internal    false    226            <           2606    37352    preference pk_preference 
   CONSTRAINT     Y   ALTER TABLE ONLY companies.preference
    ADD CONSTRAINT pk_preference PRIMARY KEY (id);
 E   ALTER TABLE ONLY companies.preference DROP CONSTRAINT pk_preference;
    	   companies            postgres    false    227            6           2606    37354    company uq_company_ssn 
   CONSTRAINT     S   ALTER TABLE ONLY companies.company
    ADD CONSTRAINT uq_company_ssn UNIQUE (ssn);
 C   ALTER TABLE ONLY companies.company DROP CONSTRAINT uq_company_ssn;
    	   companies            postgres    false    225            8           2606    37356    company uq_company_vat_id 
   CONSTRAINT     Y   ALTER TABLE ONLY companies.company
    ADD CONSTRAINT uq_company_vat_id UNIQUE (vat_id);
 F   ALTER TABLE ONLY companies.company DROP CONSTRAINT uq_company_vat_id;
    	   companies            postgres    false    225            >           2606    37358    preference uq_preference_code 
   CONSTRAINT     g   ALTER TABLE ONLY companies.preference
    ADD CONSTRAINT uq_preference_code UNIQUE (company_id, code);
 J   ALTER TABLE ONLY companies.preference DROP CONSTRAINT uq_preference_code;
    	   companies            postgres    false    227    227            @           2606    40203    country pk_country 
   CONSTRAINT     Q   ALTER TABLE ONLY general.country
    ADD CONSTRAINT pk_country PRIMARY KEY (id);
 =   ALTER TABLE ONLY general.country DROP CONSTRAINT pk_country;
       general            hestia_erp_internal    false    228            B           2606    40179    country uq_country_code 
   CONSTRAINT     S   ALTER TABLE ONLY general.country
    ADD CONSTRAINT uq_country_code UNIQUE (code);
 B   ALTER TABLE ONLY general.country DROP CONSTRAINT uq_country_code;
       general            hestia_erp_internal    false    228            D           2606    37362    device pk_device 
   CONSTRAINT     J   ALTER TABLE ONLY hr.device
    ADD CONSTRAINT pk_device PRIMARY KEY (id);
 6   ALTER TABLE ONLY hr.device DROP CONSTRAINT pk_device;
       hr            hestia_erp_internal    false    229            F           2606    37364    employee pk_employee 
   CONSTRAINT     N   ALTER TABLE ONLY hr.employee
    ADD CONSTRAINT pk_employee PRIMARY KEY (id);
 :   ALTER TABLE ONLY hr.employee DROP CONSTRAINT pk_employee;
       hr            postgres    false    230            H           2606    37366 .   employee_identification_doc pk_employee_doc_id 
   CONSTRAINT     �   ALTER TABLE ONLY hr.employee_identification_doc
    ADD CONSTRAINT pk_employee_doc_id PRIMARY KEY (employee_id, identification_doc_id);
 T   ALTER TABLE ONLY hr.employee_identification_doc DROP CONSTRAINT pk_employee_doc_id;
       hr            hestia_erp_internal    false    231    231            J           2606    37368 $   identification_doc pk_identification 
   CONSTRAINT     ^   ALTER TABLE ONLY hr.identification_doc
    ADD CONSTRAINT pk_identification PRIMARY KEY (id);
 J   ALTER TABLE ONLY hr.identification_doc DROP CONSTRAINT pk_identification;
       hr            hestia_erp_internal    false    232            L           2606    37370    timesheet pk_timesheet 
   CONSTRAINT     P   ALTER TABLE ONLY hr.timesheet
    ADD CONSTRAINT pk_timesheet PRIMARY KEY (id);
 <   ALTER TABLE ONLY hr.timesheet DROP CONSTRAINT pk_timesheet;
       hr            hestia_erp_internal    false    233            N           2606    37372     timesheet_type pk_timesheet_type 
   CONSTRAINT     Z   ALTER TABLE ONLY hr.timesheet_type
    ADD CONSTRAINT pk_timesheet_type PRIMARY KEY (id);
 F   ALTER TABLE ONLY hr.timesheet_type DROP CONSTRAINT pk_timesheet_type;
       hr            hestia_erp_internal    false    234            P           2606    37374    vacation pk_vacation 
   CONSTRAINT     N   ALTER TABLE ONLY hr.vacation
    ADD CONSTRAINT pk_vacation PRIMARY KEY (id);
 :   ALTER TABLE ONLY hr.vacation DROP CONSTRAINT pk_vacation;
       hr            hestia_erp_internal    false    235            R           2606    37376    category pk_category 
   CONSTRAINT     U   ALTER TABLE ONLY inventory.category
    ADD CONSTRAINT pk_category PRIMARY KEY (id);
 A   ALTER TABLE ONLY inventory.category DROP CONSTRAINT pk_category;
    	   inventory            hestia_erp_internal    false    236            T           2606    37378    product pk_product 
   CONSTRAINT     S   ALTER TABLE ONLY inventory.product
    ADD CONSTRAINT pk_product PRIMARY KEY (id);
 ?   ALTER TABLE ONLY inventory.product DROP CONSTRAINT pk_product;
    	   inventory            hestia_erp_internal    false    238            V           2606    37380    product uq_product_barcode 
   CONSTRAINT     [   ALTER TABLE ONLY inventory.product
    ADD CONSTRAINT uq_product_barcode UNIQUE (barcode);
 G   ALTER TABLE ONLY inventory.product DROP CONSTRAINT uq_product_barcode;
    	   inventory            hestia_erp_internal    false    238            X           2606    37382    product uq_product_reference 
   CONSTRAINT     _   ALTER TABLE ONLY inventory.product
    ADD CONSTRAINT uq_product_reference UNIQUE (reference);
 I   ALTER TABLE ONLY inventory.product DROP CONSTRAINT uq_product_reference;
    	   inventory            hestia_erp_internal    false    238            Z           2606    37384    department pk_department 
   CONSTRAINT     S   ALTER TABLE ONLY itv.department
    ADD CONSTRAINT pk_department PRIMARY KEY (id);
 ?   ALTER TABLE ONLY itv.department DROP CONSTRAINT pk_department;
       itv            hestia_erp_internal    false    239            \           2606    37386    family pk_family 
   CONSTRAINT     K   ALTER TABLE ONLY itv.family
    ADD CONSTRAINT pk_family PRIMARY KEY (id);
 7   ALTER TABLE ONLY itv.family DROP CONSTRAINT pk_family;
       itv            hestia_erp_internal    false    240            ^           2606    37388    gender pk_gender 
   CONSTRAINT     K   ALTER TABLE ONLY itv.gender
    ADD CONSTRAINT pk_gender PRIMARY KEY (id);
 7   ALTER TABLE ONLY itv.gender DROP CONSTRAINT pk_gender;
       itv            hestia_erp_internal    false    241            `           2606    37390    raw_material pk_raw_material 
   CONSTRAINT     W   ALTER TABLE ONLY itv.raw_material
    ADD CONSTRAINT pk_raw_material PRIMARY KEY (id);
 C   ALTER TABLE ONLY itv.raw_material DROP CONSTRAINT pk_raw_material;
       itv            hestia_erp_internal    false    242            b           2606    37392    size pk_size 
   CONSTRAINT     G   ALTER TABLE ONLY itv.size
    ADD CONSTRAINT pk_size PRIMARY KEY (id);
 3   ALTER TABLE ONLY itv.size DROP CONSTRAINT pk_size;
       itv            hestia_erp_internal    false    243            d           2606    37394     technical_file pk_technical_file 
   CONSTRAINT     [   ALTER TABLE ONLY itv.technical_file
    ADD CONSTRAINT pk_technical_file PRIMARY KEY (id);
 G   ALTER TABLE ONLY itv.technical_file DROP CONSTRAINT pk_technical_file;
       itv            hestia_erp_internal    false    244            h           2606    37396 *   technical_file_size pk_technical_file_size 
   CONSTRAINT     }   ALTER TABLE ONLY itv.technical_file_size
    ADD CONSTRAINT pk_technical_file_size PRIMARY KEY (technical_file_id, size_id);
 Q   ALTER TABLE ONLY itv.technical_file_size DROP CONSTRAINT pk_technical_file_size;
       itv            hestia_erp_internal    false    245    245            f           2606    37398 %   technical_file uq_technical_file_code 
   CONSTRAINT     i   ALTER TABLE ONLY itv.technical_file
    ADD CONSTRAINT uq_technical_file_code UNIQUE (company_id, code);
 L   ALTER TABLE ONLY itv.technical_file DROP CONSTRAINT uq_technical_file_code;
       itv            hestia_erp_internal    false    244    244            j           2606    37683    client pk_client 
   CONSTRAINT     M   ALTER TABLE ONLY sales.client
    ADD CONSTRAINT pk_client PRIMARY KEY (id);
 9   ALTER TABLE ONLY sales.client DROP CONSTRAINT pk_client;
       sales            postgres    false    246            p           2606    37400    document pk_document 
   CONSTRAINT     Q   ALTER TABLE ONLY sales.document
    ADD CONSTRAINT pk_document PRIMARY KEY (id);
 =   ALTER TABLE ONLY sales.document DROP CONSTRAINT pk_document;
       sales            hestia_erp_internal    false    247            �           2606    37661    document_line pk_document_line 
   CONSTRAINT     p   ALTER TABLE ONLY sales.document_line
    ADD CONSTRAINT pk_document_line PRIMARY KEY (document_id, product_id);
 G   ALTER TABLE ONLY sales.document_line DROP CONSTRAINT pk_document_line;
       sales            hestia_erp_internal    false    258    258            �           2606    37625    series pk_series 
   CONSTRAINT     M   ALTER TABLE ONLY sales.series
    ADD CONSTRAINT pk_series PRIMARY KEY (id);
 9   ALTER TABLE ONLY sales.series DROP CONSTRAINT pk_series;
       sales            hestia_erp_internal    false    254            �           2606    37644 $   series_type_code pk_series_type_code 
   CONSTRAINT     q   ALTER TABLE ONLY sales.series_type_code
    ADD CONSTRAINT pk_series_type_code PRIMARY KEY (series_id, type_id);
 M   ALTER TABLE ONLY sales.series_type_code DROP CONSTRAINT pk_series_type_code;
       sales            hestia_erp_internal    false    256    256            �           2606    37635    type type_pk 
   CONSTRAINT     I   ALTER TABLE ONLY sales.type
    ADD CONSTRAINT type_pk PRIMARY KEY (id);
 5   ALTER TABLE ONLY sales.type DROP CONSTRAINT type_pk;
       sales            hestia_erp_internal    false    255            l           2606    37685    client uq_client_code 
   CONSTRAINT     [   ALTER TABLE ONLY sales.client
    ADD CONSTRAINT uq_client_code UNIQUE (company_id, code);
 >   ALTER TABLE ONLY sales.client DROP CONSTRAINT uq_client_code;
       sales            postgres    false    246    246            n           2606    37687    client uq_client_vat_id 
   CONSTRAINT     _   ALTER TABLE ONLY sales.client
    ADD CONSTRAINT uq_client_vat_id UNIQUE (company_id, vat_id);
 @   ALTER TABLE ONLY sales.client DROP CONSTRAINT uq_client_vat_id;
       sales            postgres    false    246    246            r           2606    37681 $   document uq_document_series_num_year 
   CONSTRAINT     z   ALTER TABLE ONLY sales.document
    ADD CONSTRAINT uq_document_series_num_year UNIQUE (company_id, series_id, num, year);
 M   ALTER TABLE ONLY sales.document DROP CONSTRAINT uq_document_series_num_year;
       sales            hestia_erp_internal    false    247    247    247    247            �           0    0 2   CONSTRAINT uq_document_series_num_year ON document    COMMENT     �   COMMENT ON CONSTRAINT uq_document_series_num_year ON sales.document IS 'Because the YEAR column is generated using the Europe/Lisbon timezone people in the Azores timnezone will have a problem when creating new series at 23:00 of 31st of december';
          sales          hestia_erp_internal    false    4978            �           2606    37627    series uq_series_code 
   CONSTRAINT     O   ALTER TABLE ONLY sales.series
    ADD CONSTRAINT uq_series_code UNIQUE (code);
 >   ALTER TABLE ONLY sales.series DROP CONSTRAINT uq_series_code;
       sales            hestia_erp_internal    false    254            �           2606    37637    type uq_type_code 
   CONSTRAINT     W   ALTER TABLE ONLY sales.type
    ADD CONSTRAINT uq_type_code UNIQUE (company_id, code);
 :   ALTER TABLE ONLY sales.type DROP CONSTRAINT uq_type_code;
       sales            hestia_erp_internal    false    255    255            t           2606    37410    module pk_module 
   CONSTRAINT     U   ALTER TABLE ONLY subscriptions.module
    ADD CONSTRAINT pk_module PRIMARY KEY (id);
 A   ALTER TABLE ONLY subscriptions.module DROP CONSTRAINT pk_module;
       subscriptions            hestia_erp_internal    false    248            x           2606    37412 *   subscription_module pk_subscription_module 
   CONSTRAINT     �   ALTER TABLE ONLY subscriptions.subscription_module
    ADD CONSTRAINT pk_subscription_module PRIMARY KEY (subscription_id, module_id);
 [   ALTER TABLE ONLY subscriptions.subscription_module DROP CONSTRAINT pk_subscription_module;
       subscriptions            hestia_erp_internal    false    250    250            v           2606    37414    subscription pk_subscriptions 
   CONSTRAINT     b   ALTER TABLE ONLY subscriptions.subscription
    ADD CONSTRAINT pk_subscriptions PRIMARY KEY (id);
 N   ALTER TABLE ONLY subscriptions.subscription DROP CONSTRAINT pk_subscriptions;
       subscriptions            hestia_erp_internal    false    249            |           2606    37416    users pk_user 
   CONSTRAINT     J   ALTER TABLE ONLY users.users
    ADD CONSTRAINT pk_user PRIMARY KEY (id);
 6   ALTER TABLE ONLY users.users DROP CONSTRAINT pk_user;
       users            hestia_erp_internal    false    252            z           2606    37418    user_company pk_user_company 
   CONSTRAINT     j   ALTER TABLE ONLY users.user_company
    ADD CONSTRAINT pk_user_company PRIMARY KEY (user_id, company_id);
 E   ALTER TABLE ONLY users.user_company DROP CONSTRAINT pk_user_company;
       users            hestia_erp_internal    false    251    251            �           2606    37420    users_session pk_users_session 
   CONSTRAINT     [   ALTER TABLE ONLY users.users_session
    ADD CONSTRAINT pk_users_session PRIMARY KEY (id);
 G   ALTER TABLE ONLY users.users_session DROP CONSTRAINT pk_users_session;
       users            postgres    false    253            ~           2606    37422    users uq_users_email 
   CONSTRAINT     O   ALTER TABLE ONLY users.users
    ADD CONSTRAINT uq_users_email UNIQUE (email);
 =   ALTER TABLE ONLY users.users DROP CONSTRAINT uq_users_email;
       users            hestia_erp_internal    false    252            �           2606    40219    vat_rate fk_vat_rate_country    FK CONSTRAINT     �   ALTER TABLE ONLY accounting.vat_rate
    ADD CONSTRAINT fk_vat_rate_country FOREIGN KEY (country_id) REFERENCES general.country(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 J   ALTER TABLE ONLY accounting.vat_rate DROP CONSTRAINT fk_vat_rate_country;
    
   accounting          hestia_erp_internal    false    257    228    4928            �           2606    37738    vat_rate fk_vat_rate_vat_type    FK CONSTRAINT     �   ALTER TABLE ONLY accounting.vat_rate
    ADD CONSTRAINT fk_vat_rate_vat_type FOREIGN KEY (vat_type_id) REFERENCES accounting.vat_type(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 K   ALTER TABLE ONLY accounting.vat_rate DROP CONSTRAINT fk_vat_rate_vat_type;
    
   accounting          hestia_erp_internal    false    257    260    5010            �           2606    40204    company fk_company_country    FK CONSTRAINT     �   ALTER TABLE ONLY companies.company
    ADD CONSTRAINT fk_company_country FOREIGN KEY (country_id) REFERENCES general.country(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 G   ALTER TABLE ONLY companies.company DROP CONSTRAINT fk_company_country;
    	   companies          postgres    false    225    4928    228            �           2606    37428    location fk_location_company    FK CONSTRAINT     �   ALTER TABLE ONLY companies.location
    ADD CONSTRAINT fk_location_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 I   ALTER TABLE ONLY companies.location DROP CONSTRAINT fk_location_company;
    	   companies          hestia_erp_internal    false    225    226    4916            �           2606    40209    location fk_location_country    FK CONSTRAINT     �   ALTER TABLE ONLY companies.location
    ADD CONSTRAINT fk_location_country FOREIGN KEY (country_id) REFERENCES general.country(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 I   ALTER TABLE ONLY companies.location DROP CONSTRAINT fk_location_country;
    	   companies          hestia_erp_internal    false    4928    226    228            �           2606    37438     preference fk_preference_company    FK CONSTRAINT     �   ALTER TABLE ONLY companies.preference
    ADD CONSTRAINT fk_preference_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 M   ALTER TABLE ONLY companies.preference DROP CONSTRAINT fk_preference_company;
    	   companies          postgres    false    4916    225    227            �           2606    37443    device fk_device_company    FK CONSTRAINT     �   ALTER TABLE ONLY hr.device
    ADD CONSTRAINT fk_device_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 >   ALTER TABLE ONLY hr.device DROP CONSTRAINT fk_device_company;
       hr          hestia_erp_internal    false    229    4916    225            �           2606    37448    employee fk_employee_company    FK CONSTRAINT     �   ALTER TABLE ONLY hr.employee
    ADD CONSTRAINT fk_employee_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 B   ALTER TABLE ONLY hr.employee DROP CONSTRAINT fk_employee_company;
       hr          postgres    false    4916    230    225            �           2606    37453 C   employee_identification_doc fk_employee_identification_doc_employee    FK CONSTRAINT     �   ALTER TABLE ONLY hr.employee_identification_doc
    ADD CONSTRAINT fk_employee_identification_doc_employee FOREIGN KEY (employee_id) REFERENCES hr.employee(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 i   ALTER TABLE ONLY hr.employee_identification_doc DROP CONSTRAINT fk_employee_identification_doc_employee;
       hr          hestia_erp_internal    false    231    230    4934            �           2606    37458 M   employee_identification_doc fk_employee_identification_doc_identification_doc    FK CONSTRAINT     �   ALTER TABLE ONLY hr.employee_identification_doc
    ADD CONSTRAINT fk_employee_identification_doc_identification_doc FOREIGN KEY (identification_doc_id) REFERENCES hr.identification_doc(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 s   ALTER TABLE ONLY hr.employee_identification_doc DROP CONSTRAINT fk_employee_identification_doc_identification_doc;
       hr          hestia_erp_internal    false    4938    231    232            �           2606    37463    employee fk_employee_location    FK CONSTRAINT     �   ALTER TABLE ONLY hr.employee
    ADD CONSTRAINT fk_employee_location FOREIGN KEY (location_id) REFERENCES companies.location(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 C   ALTER TABLE ONLY hr.employee DROP CONSTRAINT fk_employee_location;
       hr          postgres    false    230    226    4922            �           2606    37468    timesheet fk_timesheet_device    FK CONSTRAINT     �   ALTER TABLE ONLY hr.timesheet
    ADD CONSTRAINT fk_timesheet_device FOREIGN KEY (device) REFERENCES hr.device(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 C   ALTER TABLE ONLY hr.timesheet DROP CONSTRAINT fk_timesheet_device;
       hr          hestia_erp_internal    false    233    4932    229            �           2606    37473    timesheet fk_timesheet_employee    FK CONSTRAINT     �   ALTER TABLE ONLY hr.timesheet
    ADD CONSTRAINT fk_timesheet_employee FOREIGN KEY (employee_id) REFERENCES hr.employee(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 E   ALTER TABLE ONLY hr.timesheet DROP CONSTRAINT fk_timesheet_employee;
       hr          hestia_erp_internal    false    230    4934    233            �           2606    37478 %   timesheet fk_timesheet_timesheet_type    FK CONSTRAINT     �   ALTER TABLE ONLY hr.timesheet
    ADD CONSTRAINT fk_timesheet_timesheet_type FOREIGN KEY (type) REFERENCES hr.timesheet_type(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 K   ALTER TABLE ONLY hr.timesheet DROP CONSTRAINT fk_timesheet_timesheet_type;
       hr          hestia_erp_internal    false    233    4942    234            �           2606    37483    vacation fk_vacation_employee    FK CONSTRAINT     �   ALTER TABLE ONLY hr.vacation
    ADD CONSTRAINT fk_vacation_employee FOREIGN KEY (employee_id) REFERENCES hr.employee(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 C   ALTER TABLE ONLY hr.vacation DROP CONSTRAINT fk_vacation_employee;
       hr          hestia_erp_internal    false    230    4934    235            �           2606    37488    vacation fk_vacation_employee_1    FK CONSTRAINT     �   ALTER TABLE ONLY hr.vacation
    ADD CONSTRAINT fk_vacation_employee_1 FOREIGN KEY (approved_by) REFERENCES hr.employee(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 E   ALTER TABLE ONLY hr.vacation DROP CONSTRAINT fk_vacation_employee_1;
       hr          hestia_erp_internal    false    235    230    4934            �           2606    37493    category fk_category_category    FK CONSTRAINT     �   ALTER TABLE ONLY inventory.category
    ADD CONSTRAINT fk_category_category FOREIGN KEY (parent_id) REFERENCES inventory.category(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 J   ALTER TABLE ONLY inventory.category DROP CONSTRAINT fk_category_category;
    	   inventory          hestia_erp_internal    false    236    4946    236            �           2606    37498    category fk_category_company    FK CONSTRAINT     �   ALTER TABLE ONLY inventory.category
    ADD CONSTRAINT fk_category_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 I   ALTER TABLE ONLY inventory.category DROP CONSTRAINT fk_category_company;
    	   inventory          hestia_erp_internal    false    4916    225    236            �           2606    37503    product fk_product_category    FK CONSTRAINT     �   ALTER TABLE ONLY inventory.product
    ADD CONSTRAINT fk_product_category FOREIGN KEY (category_id) REFERENCES inventory.category(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 H   ALTER TABLE ONLY inventory.product DROP CONSTRAINT fk_product_category;
    	   inventory          hestia_erp_internal    false    238    4946    236            �           2606    37508    product fk_product_company    FK CONSTRAINT     �   ALTER TABLE ONLY inventory.product
    ADD CONSTRAINT fk_product_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 G   ALTER TABLE ONLY inventory.product DROP CONSTRAINT fk_product_company;
    	   inventory          hestia_erp_internal    false    238    4916    225            �           2606    37513     department fk_department_company    FK CONSTRAINT     �   ALTER TABLE ONLY itv.department
    ADD CONSTRAINT fk_department_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 G   ALTER TABLE ONLY itv.department DROP CONSTRAINT fk_department_company;
       itv          hestia_erp_internal    false    225    4916    239            �           2606    37518    family fk_family_company    FK CONSTRAINT     �   ALTER TABLE ONLY itv.family
    ADD CONSTRAINT fk_family_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 ?   ALTER TABLE ONLY itv.family DROP CONSTRAINT fk_family_company;
       itv          hestia_erp_internal    false    225    4916    240            �           2606    37523    gender fk_gender_company    FK CONSTRAINT     �   ALTER TABLE ONLY itv.gender
    ADD CONSTRAINT fk_gender_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 ?   ALTER TABLE ONLY itv.gender DROP CONSTRAINT fk_gender_company;
       itv          hestia_erp_internal    false    225    4916    241            �           2606    37528 $   raw_material fk_raw_material_company    FK CONSTRAINT     �   ALTER TABLE ONLY itv.raw_material
    ADD CONSTRAINT fk_raw_material_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 K   ALTER TABLE ONLY itv.raw_material DROP CONSTRAINT fk_raw_material_company;
       itv          hestia_erp_internal    false    225    4916    242            �           2606    37533    size fk_size_company    FK CONSTRAINT     �   ALTER TABLE ONLY itv.size
    ADD CONSTRAINT fk_size_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 ;   ALTER TABLE ONLY itv.size DROP CONSTRAINT fk_size_company;
       itv          hestia_erp_internal    false    4916    225    243            �           2606    38365 '   technical_file fk_technical_file_client    FK CONSTRAINT     �   ALTER TABLE ONLY itv.technical_file
    ADD CONSTRAINT fk_technical_file_client FOREIGN KEY (client_id) REFERENCES sales.client(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 N   ALTER TABLE ONLY itv.technical_file DROP CONSTRAINT fk_technical_file_client;
       itv          hestia_erp_internal    false    244    246    4970            �           2606    37543 (   technical_file fk_technical_file_company    FK CONSTRAINT     �   ALTER TABLE ONLY itv.technical_file
    ADD CONSTRAINT fk_technical_file_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 O   ALTER TABLE ONLY itv.technical_file DROP CONSTRAINT fk_technical_file_company;
       itv          hestia_erp_internal    false    244    4916    225            �           2606    37548 +   technical_file fk_technical_file_department    FK CONSTRAINT     �   ALTER TABLE ONLY itv.technical_file
    ADD CONSTRAINT fk_technical_file_department FOREIGN KEY (department_id) REFERENCES itv.department(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 R   ALTER TABLE ONLY itv.technical_file DROP CONSTRAINT fk_technical_file_department;
       itv          hestia_erp_internal    false    239    244    4954            �           2606    37553 /   technical_file_size fk_technical_file_size_size    FK CONSTRAINT     �   ALTER TABLE ONLY itv.technical_file_size
    ADD CONSTRAINT fk_technical_file_size_size FOREIGN KEY (size_id) REFERENCES itv.size(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 V   ALTER TABLE ONLY itv.technical_file_size DROP CONSTRAINT fk_technical_file_size_size;
       itv          hestia_erp_internal    false    4962    245    243            �           2606    37558 9   technical_file_size fk_technical_file_size_technical_file    FK CONSTRAINT     �   ALTER TABLE ONLY itv.technical_file_size
    ADD CONSTRAINT fk_technical_file_size_technical_file FOREIGN KEY (technical_file_id) REFERENCES itv.technical_file(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 `   ALTER TABLE ONLY itv.technical_file_size DROP CONSTRAINT fk_technical_file_size_technical_file;
       itv          hestia_erp_internal    false    4964    245    244            �           2606    37703    client fk_client_company    FK CONSTRAINT     �   ALTER TABLE ONLY sales.client
    ADD CONSTRAINT fk_client_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 A   ALTER TABLE ONLY sales.client DROP CONSTRAINT fk_client_company;
       sales          postgres    false    4916    225    246            �           2606    40214    client fk_client_country    FK CONSTRAINT     �   ALTER TABLE ONLY sales.client
    ADD CONSTRAINT fk_client_country FOREIGN KEY (country_id) REFERENCES general.country(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 A   ALTER TABLE ONLY sales.client DROP CONSTRAINT fk_client_country;
       sales          postgres    false    228    246    4928            �           2606    37688    document fk_document_company    FK CONSTRAINT     �   ALTER TABLE ONLY sales.document
    ADD CONSTRAINT fk_document_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 E   ALTER TABLE ONLY sales.document DROP CONSTRAINT fk_document_company;
       sales          hestia_erp_internal    false    4916    247    225            �           2606    37743 '   document_line fk_document_line_document    FK CONSTRAINT     �   ALTER TABLE ONLY sales.document_line
    ADD CONSTRAINT fk_document_line_document FOREIGN KEY (document_id) REFERENCES sales.document(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 P   ALTER TABLE ONLY sales.document_line DROP CONSTRAINT fk_document_line_document;
       sales          hestia_erp_internal    false    258    247    4976            �           2606    37748 &   document_line fk_document_line_product    FK CONSTRAINT     �   ALTER TABLE ONLY sales.document_line
    ADD CONSTRAINT fk_document_line_product FOREIGN KEY (product_id) REFERENCES inventory.product(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 O   ALTER TABLE ONLY sales.document_line DROP CONSTRAINT fk_document_line_product;
       sales          hestia_erp_internal    false    4948    238    258            �           2606    37758 ,   document_line fk_document_line_vat_exemption    FK CONSTRAINT     �   ALTER TABLE ONLY sales.document_line
    ADD CONSTRAINT fk_document_line_vat_exemption FOREIGN KEY (vat_exemption_id) REFERENCES accounting.vat_exemption(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 U   ALTER TABLE ONLY sales.document_line DROP CONSTRAINT fk_document_line_vat_exemption;
       sales          hestia_erp_internal    false    259    258    5008            �           2606    37753 '   document_line fk_document_line_vat_rate    FK CONSTRAINT     �   ALTER TABLE ONLY sales.document_line
    ADD CONSTRAINT fk_document_line_vat_rate FOREIGN KEY (vat_rate_id) REFERENCES accounting.vat_rate(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 P   ALTER TABLE ONLY sales.document_line DROP CONSTRAINT fk_document_line_vat_rate;
       sales          hestia_erp_internal    false    258    257    5004            �           2606    37693    document fk_document_series    FK CONSTRAINT     �   ALTER TABLE ONLY sales.document
    ADD CONSTRAINT fk_document_series FOREIGN KEY (series_id) REFERENCES sales.series(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 D   ALTER TABLE ONLY sales.document DROP CONSTRAINT fk_document_series;
       sales          hestia_erp_internal    false    247    4994    254            �           2606    37698    series fk_series_company    FK CONSTRAINT     �   ALTER TABLE ONLY sales.series
    ADD CONSTRAINT fk_series_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 A   ALTER TABLE ONLY sales.series DROP CONSTRAINT fk_series_company;
       sales          hestia_erp_internal    false    4916    225    254            �           2606    37723 +   series_type_code fk_series_type_code_series    FK CONSTRAINT     �   ALTER TABLE ONLY sales.series_type_code
    ADD CONSTRAINT fk_series_type_code_series FOREIGN KEY (series_id) REFERENCES sales.series(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 T   ALTER TABLE ONLY sales.series_type_code DROP CONSTRAINT fk_series_type_code_series;
       sales          hestia_erp_internal    false    256    4994    254            �           2606    37728 )   series_type_code fk_series_type_code_type    FK CONSTRAINT     �   ALTER TABLE ONLY sales.series_type_code
    ADD CONSTRAINT fk_series_type_code_type FOREIGN KEY (type_id) REFERENCES sales.type(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 R   ALTER TABLE ONLY sales.series_type_code DROP CONSTRAINT fk_series_type_code_type;
       sales          hestia_erp_internal    false    256    4998    255            �           2606    37713    type fk_type_company    FK CONSTRAINT     �   ALTER TABLE ONLY sales.type
    ADD CONSTRAINT fk_type_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 =   ALTER TABLE ONLY sales.type DROP CONSTRAINT fk_type_company;
       sales          hestia_erp_internal    false    225    255    4916            �           2606    37718    type fk_type_type    FK CONSTRAINT     �   ALTER TABLE ONLY sales.type
    ADD CONSTRAINT fk_type_type FOREIGN KEY (type) REFERENCES sales.type(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 :   ALTER TABLE ONLY sales.type DROP CONSTRAINT fk_type_type;
       sales          hestia_erp_internal    false    4998    255    255            �           2606    37573 $   subscription fk_subscription_company    FK CONSTRAINT     �   ALTER TABLE ONLY subscriptions.subscription
    ADD CONSTRAINT fk_subscription_company FOREIGN KEY (company_id) REFERENCES companies.company(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 U   ALTER TABLE ONLY subscriptions.subscription DROP CONSTRAINT fk_subscription_company;
       subscriptions          hestia_erp_internal    false    4916    249    225            �           2606    37578 1   subscription_module fk_subscription_module_module    FK CONSTRAINT     �   ALTER TABLE ONLY subscriptions.subscription_module
    ADD CONSTRAINT fk_subscription_module_module FOREIGN KEY (module_id) REFERENCES subscriptions.module(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 b   ALTER TABLE ONLY subscriptions.subscription_module DROP CONSTRAINT fk_subscription_module_module;
       subscriptions          hestia_erp_internal    false    4980    250    248            �           2606    37583 7   subscription_module fk_subscription_module_subscription    FK CONSTRAINT     �   ALTER TABLE ONLY subscriptions.subscription_module
    ADD CONSTRAINT fk_subscription_module_subscription FOREIGN KEY (subscription_id) REFERENCES subscriptions.subscription(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 h   ALTER TABLE ONLY subscriptions.subscription_module DROP CONSTRAINT fk_subscription_module_subscription;
       subscriptions          hestia_erp_internal    false    250    4982    249            �           2606    37588 $   user_company fk_user_company_company    FK CONSTRAINT     �   ALTER TABLE ONLY users.user_company
    ADD CONSTRAINT fk_user_company_company FOREIGN KEY (company_id) REFERENCES companies.company(id) MATCH FULL ON UPDATE RESTRICT ON DELETE RESTRICT;
 M   ALTER TABLE ONLY users.user_company DROP CONSTRAINT fk_user_company_company;
       users          hestia_erp_internal    false    251    225    4916            �           2606    37593 %   user_company fk_user_company_employee    FK CONSTRAINT     �   ALTER TABLE ONLY users.user_company
    ADD CONSTRAINT fk_user_company_employee FOREIGN KEY (employee_id) REFERENCES hr.employee(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 N   ALTER TABLE ONLY users.user_company DROP CONSTRAINT fk_user_company_employee;
       users          hestia_erp_internal    false    230    251    4934            �           2606    37598 "   user_company fk_user_company_users    FK CONSTRAINT     �   ALTER TABLE ONLY users.user_company
    ADD CONSTRAINT fk_user_company_users FOREIGN KEY (user_id) REFERENCES users.users(id) MATCH FULL ON UPDATE RESTRICT ON DELETE RESTRICT;
 K   ALTER TABLE ONLY users.user_company DROP CONSTRAINT fk_user_company_users;
       users          hestia_erp_internal    false    251    252    4988            �           2606    37603 $   users_session fk_users_session_users    FK CONSTRAINT     �   ALTER TABLE ONLY users.users_session
    ADD CONSTRAINT fk_users_session_users FOREIGN KEY (user_id) REFERENCES users.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
 M   ALTER TABLE ONLY users.users_session DROP CONSTRAINT fk_users_session_users;
       users          postgres    false    253    252    4988            v      x������ � �      t      x������ � �      w      x������ � �      T   }   x�-�;�0 ��9E/`�q�|F&���GBJ������;�KUZ�b(^*FJ�+�)�X��A�m�����lk}i����{�n�']��e��٣�ܚ*Q@O�~}Q,)j
ƙXFa�88��J#`      U      x������ � �      V      x������ � �      W   @   x���/*)MO��,�4MKKL444�50L5�51LJ�M232�M43J5574�0�0����� �%6      X      x������ � �      Y      x������ � �      Z      x������ � �      [      x������ � �      \      x������ � �      ]      x������ � �      ^      x������ � �      _      x������ � �      `      x������ � �      a      x������ � �      b      x������ � �      c      x������ � �      d      x������ � �      e      x������ � �      f      x������ � �      g      x������ � �      h      x������ � �      i      x������ � �      j      x������ � �      u      x������ � �      q      x������ � �      s      x������ � �      r      x������ � �      k      x������ � �      l      x������ � �      m      x������ � �      n   L   x����  ���B��~���#4�9���J�����Y�$Q��]�(0�&�����ݶ߾���J��      o     x����NA�뻧�l����ӥ рD:��&'E��r ��\zJ��#��_�B��R��PJu�w��sf%�ix�ˤ�ݛ.:-<�z[��&^UN��<���;l�f�cc��|�EͲp��c�d!�H�i�lJ)A0ԭ�B��q�y{/T��Sj�*&�
e��ڳR�L]���e����tk�eܢX�G�	�!Xr�FtI�U� Xx8n ���.�7��~�����+$�sC�K!�HΖ���!7HH!KT��Yl�!��*ϯ��c?��/�A~L      p   �  x���9r$9E��)�O  @ȳ����G����nc���,)��
�_�,^�3��(d�o�!I�M�;�.��ץ�l�j9d�4�Fm���H��?8��}�rK�K��Y����UW�Ѹ���	ۓ���lZ׬Ǒ`&v*����z�_�Q���rK����� ß�ʂvk^}��G+\�)�t-wmWc��P_�t�jF�J����XZ��cf|��[�R)M���#�;�ͅq�˳乥8���?���qY�|8���U�ø��*&�N���q�!q��T��>�=W.�t!s b�As�m�ط�O8T�j�Ty8Ew&�����l��o�my�%�G�%\nQ���3�5�y�~�ǠG�|P~���~�n������v������t���M��r'���os�ɛ��f�yJN,��3�G VdNQ��>��+_��Eyq6�٩�g?�
� 3(_ʪnq������e�ٙ})����8SbR�c�1���F��S�V^6�[y�aښ�ن�Y��״��D��n΋=��d��D�U����∜p��c[W_l���S�Z��O����L8d�����]�-��g2���%͸���x*�1��[��H�b"Yǳ�78���[�<q�R�'��֊����z���=���]TPOo��o�B��H����/}��|j��d\Zm͖���A-p��Pa/]��ޜ�m�$����=��hY���٩��Լj3$�K�.��:
	:���L3Z[��VX��<�����?��k��n��SG���4�������<�ͯd�!'s@d2������7�G�ފ�?�|j	����
H�	Xo�7����Z���~s�7��f�?�zf7t�!ȃ�Q��N�6���s��o8��q���^u:�,c�J���y��D�2�2���?��Vqo��˧�^___���U     