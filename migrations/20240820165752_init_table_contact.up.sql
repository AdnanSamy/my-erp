CREATE TABLE IF NOT EXISTS contact(
    id serial NOT NULL,
    "name" varchar(250) not null,
    "type" int4 not null,
    company_name varchar(250),
    "address" text,
    vat varchar(100),
    job_position varchar(100),
    phone varchar(50),
    mobile varchar(50),
    email varchar(100),
    website varchar(100),
    title_id int,
    active boolean not null,
    created_at timestamp default now(),
    created_uid integer,
    updated_at timestamp default now(),
    updated_uid integer,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS contact_category(
    id serial NOT NULL,
    "name" varchar(50),
    active boolean not null,
    created_at timestamp default now(),
    created_uid integer,
    updated_at timestamp default now(),
    updated_uid integer,
    PRIMARY KEY (id)
);
create table if not exists contact_contact_category_rel (
    contact_id integer,
    contact_category_id integer
);