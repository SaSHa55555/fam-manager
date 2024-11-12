create table families
(
    id serial primary key,
    name varchar unique ,
    pswd varchar not null
);

create table members
(
    id serial unique,
    name varchar not null,
    family_id integer references families(id),
    constraint members_family_id_name_pk primary key(family_id, name)
);

create type priority as enum('high', 'medium', 'low');

create type status as enum('ready for work', 'in progress', 'done');

create table tasks
(
    id serial,
    name varchar not null,
    description varchar not null,
    points integer not null,
    priority priority not null,
    assignee integer references members(id),
    status status not null,
    family_id integer references families(id),
    constraint tasks_family_id_name_pk primary key(family_id, name)
);
