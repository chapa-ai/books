CREATE TABLE IF NOT EXISTS "books"
(
    id SERIAL UNIQUE NOT NULL,
    bookname text,
    authorid int,
    CONSTRAINT "pk_Books" PRIMARY KEY ("id")
);


CREATE TABLE IF NOT EXISTS "authors"
(
    id SERIAL UNIQUE NOT NULL,
    authorname text,
    bookid int,
    CONSTRAINT "pk_Authors" PRIMARY KEY ("id")
);

-- ALTER TABLE "books" ADD CONSTRAINT "fk_Books_usd" FOREIGN KEY("authorid")
-- REFERENCES "authors" ("id") ON DELETE CASCADE;

-- ALTER TABLE "authors" ADD CONSTRAINT "fk_Books_usd" FOREIGN KEY("bookid")
-- REFERENCES "books" ("id") ON DELETE CASCADE;






