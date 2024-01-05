-- +goose Up
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

create table hellos
(
    id                       uuid      default uuid_generate_v4() not null primary key,
    created_at               timestamp default now(),
    updated_at               timestamp default now(),
    hello_type               bigint,
    person_name              text
);

create table option_definitions
( id                         uuid default uuid_generate_v4() not null primary key
, created_at                 timestamp default now()
, updated_at                 timestamp default now()
, name                       text not null unique
, description                text not null
, default_value              text
, option_type                int not null
, schema                     text
);

create table groups
( id                         uuid default uuid_generate_v4() not null primary key
, created_at                 timestamp default now()
, updated_at                 timestamp default now()
, group_name                 text not null
, weight                     int not null default 0
);

create table option_overrides
( id                         uuid default uuid_generate_v4() not null primary key
, created_at                 timestamp default now()
, updated_at                 timestamp default now()
, option_definition_id       uuid not null references option_definitions(id) on delete cascade
, option_value               jsonb not null
, group_id                   uuid not null references groups(id) on delete cascade
);
