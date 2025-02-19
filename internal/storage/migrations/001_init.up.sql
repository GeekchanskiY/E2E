begin;

-- scoped access grants availability to see non-private operations
-- zero access grants only see current wallet amount
create type access_level as enum('owner', 'full', 'scoped', 'zero');

create type gender as enum('male', 'female');

-- USERS
create table users(
    id serial primary key,
    username varchar(255) not null unique,
    password_hash text not null,
    name varchar(255) not null,
    gender gender,
    birthday date,
    CONSTRAINT age_check CHECK (AGE(birthday) >= INTERVAL '18 years')
);

create table permission_groups(
    id serial primary key,
    name varchar(255) unique,
    created_at timestamp default current_timestamp,
    level access_level
);

create table user_permission(
    id serial primary key,
    permission_group int references permission_groups(id),
    user_id int references users(id),
    unique (permission_group, user_id)
);

--  wallet and finances
create table wallets(
    id serial primary key,
    name varchar(255) not null,
    description text,
    permission_group int references permission_groups(id),
    created_at timestamp,
    currency varchar(3), -- ISO 4217
    is_salary bool
);

-- need to add check if user has higher privileges on group than wallet
create table operation_groups(
    id serial primary key,
    name varchar(255),
    wallet int references wallets(id)
);

create table operations(
    id serial primary key,
    operation_group int references operation_groups(id),
    amount float,
    time timestamp default current_timestamp, -- may be in future
    is_confirmed bool default true,
    initiator int references users(id)
);

create table distributors(
    id serial primary key,
    name varchar(255),
    source_wallet int references wallets(id),
    target_wallet int references wallets(id),
    percent float default 5

);

end;
