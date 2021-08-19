-- +migrate Up

CREATE SEQUENCE IF NOT EXISTS user_id_seq;

CREATE TABLE IF NOT EXISTS t_user(
    id         int not null unique primary key default nextval('user_id_seq'),
    name       varchar(255) NOT NULL,
    age        int      NOT NULL,
    user_type  int      NOT NULL,
    created_at timestamp    NOT NULL,
    updated_at timestamp    NOT NULL
);

CREATE SEQUENCE IF NOT EXISTS item_id_seq;

CREATE TABLE IF NOT EXISTS t_item(
    id         int not null unique primary key default nextval('item_id_seq'),
    name       varchar(255) NOT NULL,
    user_id    int NOT NULL,
    created_at timestamp    NOT NULL,
    updated_at timestamp    NOT NULL,
    CONSTRAINT fk_t_user
        FOREIGN KEY(user_id)
        REFERENCES t_user(id)
        ON DELETE SET NULL
);



-- +migrate Down
DROP TABLE t_user;
DROP TABLE t_item;