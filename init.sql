
-- DROP TABLE "users";
-- DROP TABLE "user_contact";

CREATE TABLE "users" (
    "id" INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "name" VARCHAR(40),
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE "user_contact" (
    "id_user" INT REFERENCES "users"("id"),
    "phone" VARCHAR(15),
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);
