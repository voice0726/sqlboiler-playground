create table department
(
    id serial not null
        constraint department_pk
        primary key,
    name varchar(255) not null,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null
);

alter table department owner to root;



create table "user"
(
    id serial not null
        constraint user_pk
        primary key,
    name varchar(255) not null,
    department_id integer
        constraint user_department_id_fk
        references department
        on update cascade on delete cascade,
    created_at timestamp with time zone default CURRENT_TIMESTAMP not null,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP not null
);

alter table "user" owner to root;

