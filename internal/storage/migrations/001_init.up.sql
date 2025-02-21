begin;

-- scoped access grants availability to see non-private operations
-- zero access grants only see current wallet amount
create type access_level as enum('owner', 'full', 'read');

create type gender as enum('male', 'female');

-- USERS
create table users(
    id serial primary key,
    username varchar(255) not null unique,
    password_hash text not null,
    name varchar(255) not null,
    gender gender not null default 'male',
    birthday date not null,
    CONSTRAINT age_check CHECK (AGE(birthday) >= INTERVAL '18 years')
);

create table permission_groups(
    id serial primary key,
    name varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

create table user_permission(
    id serial primary key,
    permission_group_id int references permission_groups(id),
    user_id int references users(id),
    level access_level,
    created_at timestamp default current_timestamp,
    unique (permission_group_id, user_id)
);

--  wallet and finances
create table wallets(
    id serial primary key,
    name varchar(255) not null,
    description text,
    permission_group_id int references permission_groups(id),
    created_at timestamp,
    currency varchar(3), -- ISO 4217
    is_salary bool
);

-- need to add check if user has higher privileges on group than wallet
create table operation_groups(
    id serial primary key,
    name varchar(255),
    wallet_id int references wallets(id)
);

create table operations(
    id serial primary key,
    operation_group_id int references operation_groups(id),
    amount float,
    time timestamp default current_timestamp, -- may be in future
    is_confirmed bool default true,
    initiator_id int references users(id)
);

create table distributors(
    id serial primary key,
    name varchar(255),
    source_wallet_id int references wallets(id),
    target_wallet_id int references wallets(id),
    percent float default 5

);

end;
