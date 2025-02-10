begin;

-- scoped access grants availability to see non-private operations
-- zero access grants only see current wallet amount
create type access_level as enum('full', 'scoped', 'zero');

create type gender as enum('male', 'female');

create table users(
    id serial primary key,
    name varchar(255) not null,
    gender gender,
    age int,
    birthday timestamp,
    check ( age between 18 and 100 )
);

create table wallets(
    id serial primary key,
    name varchar(255) not null,
    owner int references users(id),
    created_at timestamp,
    currency varchar(3), -- ISO 4217
    unique (name, owner)
);

create table wallet_access(
    wallet_id int references wallets(id),
    user_id int references users(id),
    created_at timestamp,
    access_level access_level
);

create table operations(
    -- todo: finish
);

end;
