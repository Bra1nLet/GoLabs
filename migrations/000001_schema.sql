CREATE SCHEMA IF NOT EXISTS gowebapp;

-- ************************************** webapp.users

CREATE TABLE gowebapp.users
(
    User_ID        bigserial NOT NULL,
    User_Name      text NOT NULL,
    Pass_Word_Hash text NOT NULL,
    Name           text NOT NULL,
    Config         jsonb NOT NULL DEFAULT '{}'::JSONB,
    Created_At     timestamp NOT NULL DEFAULT NOW(),
    Is_Enabled     boolean NOT NULL DEFAULT TRUE,
    CONSTRAINT PK_users PRIMARY KEY ( User_ID )
);
