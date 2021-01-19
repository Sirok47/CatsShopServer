CREATE TABLE public.operations
(
    newownersnick character varying NOT NULL,
    catid integer PRIMARY KEY,
    purchasedate date NOT NULL,
    status character varying NOT NULL,
)