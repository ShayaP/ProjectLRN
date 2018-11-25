CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "users" (
"id" TEXT PRIMARY KEY,
"first_name" TEXT NOT NULL,
"last_name" TEXT NOT NULL,
"phone_number" TEXT NOT NULL,
"tutor_id" INTEGER NOT NULL,
"tutee_id" INTEGER NOT NULL,
"google_id" TEXT NOT NULL,
"email" TEXT NOT NULL,
"gender" INTEGER NOT NULL,
"other_specify" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE UNIQUE INDEX "users_tutor_id_idx" ON "users" (tutor_id);
CREATE UNIQUE INDEX "users_tutee_id_idx" ON "users" (tutee_id);
