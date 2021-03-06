CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "users" (
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
"is_tutor" INTEGER NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE UNIQUE INDEX "users_email_idx" ON "users" (email);
CREATE UNIQUE INDEX "users_google_id_idx" ON "users" (google_id);
CREATE TABLE IF NOT EXISTS "requests" (
"id" TEXT PRIMARY KEY,
"status" INTEGER NOT NULL,
"senderid" char(36) NOT NULL,
"receiverid" char(36) NOT NULL,
"topic" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE INDEX "requests_senderid_idx" ON "requests" (senderid);
CREATE INDEX "requests_receiverid_idx" ON "requests" (receiverid);
CREATE TABLE IF NOT EXISTS "reviews" (
"id" TEXT PRIMARY KEY,
"rating" INTEGER NOT NULL,
"description" TEXT NOT NULL,
"reviewer" TEXT NOT NULL,
"reviewee" TEXT NOT NULL,
"astutor" INTEGER NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "courses" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"instructor" TEXT NOT NULL,
"subject" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "subjects" (
"id" TEXT PRIMARY KEY,
"name" INTEGER NOT NULL,
"prof" INTEGER NOT NULL,
"used_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "userinfoes" (
"id" TEXT PRIMARY KEY,
"google_id" TEXT NOT NULL,
"languages" TEXT NOT NULL,
"subjects" TEXT NOT NULL,
"courses" TEXT NOT NULL,
"address" TEXT NOT NULL,
"user_id" char(36) NOT NULL,
"tutors" TEXT NOT NULL,
"tutees" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
FOREIGN KEY (user_id) REFERENCES users (id)
);
CREATE UNIQUE INDEX "userinfoes_google_id_idx" ON "userinfoes" (google_id);
