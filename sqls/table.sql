create extension btree_gist;

create table if not exists codes
(
    id serial primary key not null,
    value char(10) not null,
    owner varchar(150) not null,
    domain varchar(150) not null,
    valid_between tstzrange not null,

    constraint unique_code_between_time_window exclude using gist (valid_between with &&, value with =)
);

create table if not exists user_subscriptions(
    id serial primary key not null,
    owner varchar(150) not null,
    domain varchar(150) not null
);
