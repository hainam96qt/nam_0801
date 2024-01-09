DROP TABLE IF EXISTS users;

create table users (
    id SERIAL primary key,
    name text not null,
    email text not null,
    password text not null,
    created_at timestamp default current_timestamp
);

CREATE TYPE bank_name AS ENUM ('VCB', 'ACB', 'VIB');

DROP TABLE IF EXISTS accounts;

create table accounts (
    id SERIAL primary key,
    user_id int not null,
    name text not null,
    bank bank_name not null ,
    balance FLOAT not null default 0,
    created_at timestamp default current_timestamp not null
);

alter table accounts add foreign key (user_id) references users(id);

CREATE TYPE transaction_type AS ENUM('withdraw','deposit');

DROP TABLE IF EXISTS transactions;

create table transactions(
    id SERIAL primary key,
    amount FLOAT not null,
    account_id int not null,
    transaction_type transaction_type not null ,
    created_at timestamp default current_timestamp not null
);

alter table transactions add foreign key (account_id) references accounts(id);