begin;

create type event_type as enum('log', 'event');

create table registry (
    id serial primary key,
    name varchar(255),
    event event_type,
    content text,
    time timestamp default current_timestamp
);

create index on registry(event);

end;
