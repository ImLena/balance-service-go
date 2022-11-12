create table if not exists balance
(
    id      varchar,
    balance float
);
create table if not exists transactions
(
    id         varchar,
    service_id int,
    order_id   int,
    price      int,
    user_id    varchar,
    verified   bool,
    comment    varchar,
    time TIMESTAMP
);
