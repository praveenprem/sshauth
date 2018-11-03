CREATE SCHEMA IF NOT EXISTS sshAuth;

CREATE TABLE IF NOT EXISTS sshAuth.users (
    id          INT                                   NOT NULL,
    username    VARCHAR(255)                          NOT NULL,
    fullname    VARCHAR(255),
    publicKey   TEXT                                  NOT NULL,
    createdOn   TIMESTAMP DEFAULT CURRENT_TIMESTAMP   NOT NULL,
    PRIMARY KEY (id)
);
