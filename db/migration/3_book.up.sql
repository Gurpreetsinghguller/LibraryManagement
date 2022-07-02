CREATE TABLE IF NOT EXISTS "book" (

    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    "name" TEXT NOT NULL,

    "price" FLOAT   NOT NULL,

    "author_id" UUID REFERENCES authors(id) NOT NULL,

    "published_country" INTEGER REFERENCES country(id) NOT NULL,

    "category" INTEGER REFERENCES category(id) NOT NULL
);