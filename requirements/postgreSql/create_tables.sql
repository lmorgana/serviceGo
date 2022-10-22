CREATE TABLE IF NOT EXISTS users (
    id_user int NOT NULL,
    balance int DEFAULT 0,
    PRIMARY KEY(id_user)
);

CREATE TABLE IF NOT EXISTS orders (
    id_user int NOT NULL,
    id_service int NOT NULL,
    id_order int NOT NULL,
    cost int,
    PRIMARY KEY(id_user)
)