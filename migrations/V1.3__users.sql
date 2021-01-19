CREATE TABLE public.users
(
    nickname character varying PRIMARY KEY,
    isadmin boolean NOT NULL,
    password character varying NOT NULL,
    address character varying,
)