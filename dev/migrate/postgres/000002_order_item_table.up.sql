create table "order".items
(
    id          uuid                    not null
        constraint items_pk
            primary key,
    status      smallint                not null,
    ts_create   timestamp default now() not null,
    ts_modify   timestamp default now() not null,
    name        varchar(64)             not null,
    description varchar(256)            not null,
    user_id     uuid                    not null
);

alter table "order".items
    owner to krivenkov;

create unique index items_id_uindex
    on "order".items (id);
