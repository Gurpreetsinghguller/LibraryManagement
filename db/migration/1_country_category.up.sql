CREATE TABLE IF NOT EXISTS "country" (

    "id" SERIAL PRIMARY KEY,

    "name" TEXT NOT NULL,

    "currency" TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS "category" (

    "id" SERIAL PRIMARY KEY,

    "name" TEXT NOT NULL
);

INSERT INTO "country" ("name", "currency") VALUES
('India',  'INR'),
('United Kingdom', 'GBP');

INSERT INTO "category" ( "name") VALUES
( 'Action'),
( 'Classics'),
(  'Horror'),
( 'Fantasy'),
( 'Romance'),
(  'Adventure');