CREATE TABLE "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE "users" (
"id" TEXT PRIMARY KEY,
"first_name" TEXT NOT NULL,
"last_name" TEXT NOT NULL,
"phone_number" TEXT NOT NULL,
"google_id" TEXT NOT NULL,
"email" TEXT NOT NULL,
"gender" INTEGER NOT NULL,
"other_specify" TEXT NOT NULL,
"avg_rating" REAL NOT NULL,
"num_ratings" INTEGER NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE UNIQUE INDEX "users_email_idx" ON "users" (email);
CREATE TABLE "requests" (
"id" TEXT PRIMARY KEY,
"status" INTEGER NOT NULL,
"tutorid" char(36) NOT NULL,
"tuteeid" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE "reviews" (
"id" TEXT PRIMARY KEY,
"rating" INTEGER NOT NULL,
"description" TEXT NOT NULL,
"user" char(36) NOT NULL,
"astutor" INTEGER NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE "courses" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"instructor" TEXT NOT NULL,
"subject" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE "subjects" (
"id" TEXT PRIMARY KEY,
"name" INTEGER NOT NULL,
"prof" INTEGER NOT NULL,
"used_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE "userinfoes" (
"id" TEXT PRIMARY KEY,
"google_id" char(36) NOT NULL,
"languages" TEXT NOT NULL,
"subjects" TEXT NOT NULL,
"courses" TEXT NOT NULL,
"address" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE UNIQUE INDEX "userinfoes_google_id_idx" ON "userinfoes" (google_id);
