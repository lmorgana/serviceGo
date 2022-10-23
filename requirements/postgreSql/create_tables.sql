CREATE TABLE IF NOT EXISTS Users (
    id_user int NOT NULL,
    balance int DEFAULT 0,
    res_balance int DEFAULT 0,
    PRIMARY KEY(id_user)
);

CREATE TABLE IF NOT EXISTS Orders (
    id_user int NOT NULL REFERENCES Users(id_user),
    id_service int NOT NULL,
    id_order int NOT NULL,
    cost int DEFAULT 0,
    status int2 DEFAULT 0 CHECK ( status < 3),
    PRIMARY KEY(id_order)
);