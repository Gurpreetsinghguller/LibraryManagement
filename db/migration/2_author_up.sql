DO $$
DECLARE
    country_IN INTEGER;
    country_UK INTEGER;
    
    
    category_Action INTEGER;
    category_Romance INTEGER;

BEGIN

    SELECT "id" INTO country_IN FROM "country" WHERE "name"='India';
    SELECT "id" INTO country_IN FROM "country" WHERE "name"='United Kingdom';
    
    SELECT "id" INTO category_Action FROM "category" WHERE "name"='Action';
    SELECT "id" INTO category_Romance FROM "category" WHERE "name"='Romance';
END$$;

CREATE TABLE IF NOT EXISTS "authors" (

    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    "name" TEXT NOT NULL,

    "country" INTEGER REFERENCES country(id) NOT NULL, 

    "category" INTEGER REFERENCES category(id) NOT NULL
);

INSERT INTO "authors" ("name", "country", "category") VALUES
('Vikram Chandra', country_IN , category_Action),
('Rohinton Mistry', country_IN ,category_Romance ),
('Jane Austen',  country_UK,category_Action ),
('Charles Dickens',  country_UK, category_Romance);
