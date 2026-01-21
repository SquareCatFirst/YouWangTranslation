/*
 Navicat Premium Data Transfer

 Source Server         : postgresqloncentos7
 Source Server Type    : PostgreSQL
 Source Server Version : 120019 (120019)
 Source Host           : 192.168.137.139:5432
 Source Catalog        : YwTrans
 Source Schema         : ywtrans

 Target Server Type    : PostgreSQL
 Target Server Version : 120019 (120019)
 File Encoding         : 65001

 Date: 21/01/2026 16:55:12
*/


-- ----------------------------
-- Sequence structure for chapter_assignments_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "ywtrans"."chapter_assignments_id_seq";
CREATE SEQUENCE "ywtrans"."chapter_assignments_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for chapter_work_images_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "ywtrans"."chapter_work_images_id_seq";
CREATE SEQUENCE "ywtrans"."chapter_work_images_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for chapter_work_logs_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "ywtrans"."chapter_work_logs_id_seq";
CREATE SEQUENCE "ywtrans"."chapter_work_logs_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for chapter_works_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "ywtrans"."chapter_works_id_seq";
CREATE SEQUENCE "ywtrans"."chapter_works_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for messages_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "ywtrans"."messages_id_seq";
CREATE SEQUENCE "ywtrans"."messages_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for project_assignments_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "ywtrans"."project_assignments_id_seq";
CREATE SEQUENCE "ywtrans"."project_assignments_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for project_chapters_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "ywtrans"."project_chapters_id_seq";
CREATE SEQUENCE "ywtrans"."project_chapters_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for project_tags_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "ywtrans"."project_tags_id_seq";
CREATE SEQUENCE "ywtrans"."project_tags_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for projects_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "ywtrans"."projects_id_seq";
CREATE SEQUENCE "ywtrans"."projects_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for tags_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "ywtrans"."tags_id_seq";
CREATE SEQUENCE "ywtrans"."tags_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "ywtrans"."users_id_seq";
CREATE SEQUENCE "ywtrans"."users_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 10001
CACHE 1;

-- ----------------------------
-- Table structure for chapter_assignments
-- ----------------------------
DROP TABLE IF EXISTS "ywtrans"."chapter_assignments";
CREATE TABLE "ywtrans"."chapter_assignments" (
                                                 "id" int8 NOT NULL DEFAULT nextval('"ywtrans".chapter_assignments_id_seq'::regclass),
                                                 "chapter_id" int8 NOT NULL,
                                                 "user_id" int8 NOT NULL,
                                                 "role" int4 NOT NULL,
                                                 "created_at" timestamptz(6) NOT NULL DEFAULT now(),
                                                 "updated_at" timestamptz(6) NOT NULL DEFAULT now(),
                                                 "assigned_by" int8
)
;

-- ----------------------------
-- Table structure for chapter_work_images
-- ----------------------------
DROP TABLE IF EXISTS "ywtrans"."chapter_work_images";
CREATE TABLE "ywtrans"."chapter_work_images" (
                                                 "id" int8 NOT NULL DEFAULT nextval('"ywtrans".chapter_work_images_id_seq'::regclass),
                                                 "chapter_id" int8 NOT NULL,
                                                 "role" int2 NOT NULL,
                                                 "image_url" text COLLATE "pg_catalog"."default" NOT NULL,
                                                 "order_index" int4 NOT NULL,
                                                 "created_at" timestamptz(6) NOT NULL DEFAULT now(),
                                                 "updated_at" timestamptz(6) NOT NULL DEFAULT now()
)
;

-- ----------------------------
-- Table structure for chapter_work_logs
-- ----------------------------
DROP TABLE IF EXISTS "ywtrans"."chapter_work_logs";
CREATE TABLE "ywtrans"."chapter_work_logs" (
                                               "id" int8 NOT NULL DEFAULT nextval('"ywtrans".chapter_work_logs_id_seq'::regclass),
                                               "chapter_id" int8 NOT NULL,
                                               "order_index" int4 NOT NULL,
                                               "role" int2,
                                               "action_type" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
                                               "user_id" int8,
                                               "note" text COLLATE "pg_catalog"."default",
                                               "created_at" timestamptz(6) NOT NULL DEFAULT now(),
                                               "updated_at" timestamptz(6) NOT NULL DEFAULT now()
)
;
COMMENT ON COLUMN "ywtrans"."chapter_work_logs"."order_index" IS '代表这是章节的第几章图片';

-- ----------------------------
-- Table structure for chapter_works
-- ----------------------------
DROP TABLE IF EXISTS "ywtrans"."chapter_works";
CREATE TABLE "ywtrans"."chapter_works" (
                                           "id" int8 NOT NULL DEFAULT nextval('"ywtrans".chapter_works_id_seq'::regclass),
                                           "chapter_id" int8 NOT NULL,
                                           "translation_text" text COLLATE "pg_catalog"."default",
                                           "proofread_text" text COLLATE "pg_catalog"."default",
                                           "stage" int2 NOT NULL DEFAULT 1,
                                           "created_at" timestamptz(6) NOT NULL DEFAULT now(),
                                           "updated_at" timestamptz(6) NOT NULL DEFAULT now()
)
;

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS "ywtrans"."messages";
CREATE TABLE "ywtrans"."messages" (
                                      "id" int8 NOT NULL DEFAULT nextval('"ywtrans".messages_id_seq'::regclass),
                                      "user_id" int8 NOT NULL,
                                      "sender_id" int8,
                                      "type" int2 NOT NULL,
                                      "title" varchar(128) COLLATE "pg_catalog"."default",
                                      "content" text COLLATE "pg_catalog"."default",
                                      "ref_type" varchar(32) COLLATE "pg_catalog"."default",
                                      "ref_id" int8,
                                      "is_read" bool NOT NULL DEFAULT false,
                                      "created_at" timestamptz(6) NOT NULL DEFAULT now()
)
;

-- ----------------------------
-- Table structure for project_assignments
-- ----------------------------
DROP TABLE IF EXISTS "ywtrans"."project_assignments";
CREATE TABLE "ywtrans"."project_assignments" (
                                                 "id" int8 NOT NULL DEFAULT nextval('"ywtrans".project_assignments_id_seq'::regclass),
                                                 "project_id" int8 NOT NULL,
                                                 "user_id" int8 NOT NULL,
                                                 "role" int4 NOT NULL,
                                                 "created_at" timestamptz(6) NOT NULL DEFAULT now(),
                                                 "updated_at" timestamptz(6) NOT NULL DEFAULT now(),
                                                 "assigned_by" int8
)
;

-- ----------------------------
-- Table structure for project_chapters
-- ----------------------------
DROP TABLE IF EXISTS "ywtrans"."project_chapters";
CREATE TABLE "ywtrans"."project_chapters" (
                                              "id" int8 NOT NULL DEFAULT nextval('"ywtrans".project_chapters_id_seq'::regclass),
                                              "project_id" int8 NOT NULL,
                                              "chapter_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                              "is_visible" bool NOT NULL DEFAULT false,
                                              "order_index" int4 NOT NULL DEFAULT 1,
                                              "created_at" timestamptz(6) NOT NULL DEFAULT now(),
                                              "updated_at" timestamptz(6) NOT NULL DEFAULT now()
)
;

-- ----------------------------
-- Table structure for project_tags
-- ----------------------------
DROP TABLE IF EXISTS "ywtrans"."project_tags";
CREATE TABLE "ywtrans"."project_tags" (
                                          "id" int8 NOT NULL DEFAULT nextval('"ywtrans".project_tags_id_seq'::regclass),
                                          "project_id" int8 NOT NULL,
                                          "tag_id" int8 NOT NULL,
                                          "created_at" timestamptz(6) NOT NULL DEFAULT now(),
                                          "updated_at" timestamptz(6) NOT NULL DEFAULT now()
)
;

-- ----------------------------
-- Table structure for projects
-- ----------------------------
DROP TABLE IF EXISTS "ywtrans"."projects";
CREATE TABLE "ywtrans"."projects" (
                                      "id" int8 NOT NULL DEFAULT nextval('"ywtrans".projects_id_seq'::regclass),
                                      "cover_url" varchar(255) COLLATE "pg_catalog"."default",
                                      "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                      "translated_name" varchar(128) COLLATE "pg_catalog"."default",
                                      "author" varchar(128) COLLATE "pg_catalog"."default",
                                      "source_site" varchar(128) COLLATE "pg_catalog"."default",
                                      "description" text COLLATE "pg_catalog"."default",
                                      "translation_description" text COLLATE "pg_catalog"."default",
                                      "status" int4 NOT NULL DEFAULT 1,
                                      "created_at" timestamptz(6) NOT NULL DEFAULT now(),
                                      "updated_at" timestamptz(6) NOT NULL DEFAULT now(),
                                      "source_max" int4 NOT NULL DEFAULT 1,
                                      "translator_max" int4 NOT NULL DEFAULT 1,
                                      "proofreader_max" int4 NOT NULL DEFAULT 1,
                                      "typesetter_max" int4 NOT NULL DEFAULT 1,
                                      "supervisor_max" int4 NOT NULL DEFAULT 1
)
;
COMMENT ON COLUMN "ywtrans"."projects"."source_max" IS '图源最大人数';
COMMENT ON COLUMN "ywtrans"."projects"."translator_max" IS '翻译最大人数';
COMMENT ON COLUMN "ywtrans"."projects"."proofreader_max" IS '校对最大人数';
COMMENT ON COLUMN "ywtrans"."projects"."typesetter_max" IS '嵌字最大人数';
COMMENT ON COLUMN "ywtrans"."projects"."supervisor_max" IS '监督最大人数';

-- ----------------------------
-- Table structure for tags
-- ----------------------------
DROP TABLE IF EXISTS "ywtrans"."tags";
CREATE TABLE "ywtrans"."tags" (
                                  "id" int8 NOT NULL DEFAULT nextval('"ywtrans".tags_id_seq'::regclass),
                                  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                  "color" varchar(7) COLLATE "pg_catalog"."default" NOT NULL DEFAULT '#FFFFFF'::character varying,
                                  "created_at" timestamptz(6) NOT NULL DEFAULT now(),
                                  "updated_at" timestamptz(6) NOT NULL DEFAULT now()
)
;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "ywtrans"."users";
CREATE TABLE "ywtrans"."users" (
                                   "id" int8 NOT NULL DEFAULT nextval('"ywtrans".users_id_seq'::regclass),
                                   "avatar_url" varchar(255) COLLATE "pg_catalog"."default",
                                   "nickname" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                   "role" int4 NOT NULL DEFAULT 1,
                                   "qq_openid" varchar(64) COLLATE "pg_catalog"."default",
                                   "status" int4 NOT NULL DEFAULT 1,
                                   "registered_at" timestamptz(6) NOT NULL DEFAULT now(),
                                   "last_login_at" timestamptz(6),
                                   "username" varchar(64) COLLATE "pg_catalog"."default",
                                   "password_hash" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "ywtrans"."chapter_assignments_id_seq"
    OWNED BY "ywtrans"."chapter_assignments"."id";
SELECT setval('"ywtrans"."chapter_assignments_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "ywtrans"."chapter_work_images_id_seq"
    OWNED BY "ywtrans"."chapter_work_images"."id";
SELECT setval('"ywtrans"."chapter_work_images_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "ywtrans"."chapter_work_logs_id_seq"
    OWNED BY "ywtrans"."chapter_work_logs"."id";
SELECT setval('"ywtrans"."chapter_work_logs_id_seq"', 2, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "ywtrans"."chapter_works_id_seq"
    OWNED BY "ywtrans"."chapter_works"."id";
SELECT setval('"ywtrans"."chapter_works_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "ywtrans"."messages_id_seq"
    OWNED BY "ywtrans"."messages"."id";
SELECT setval('"ywtrans"."messages_id_seq"', 2, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "ywtrans"."project_assignments_id_seq"
    OWNED BY "ywtrans"."project_assignments"."id";
SELECT setval('"ywtrans"."project_assignments_id_seq"', 3, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "ywtrans"."project_chapters_id_seq"
    OWNED BY "ywtrans"."project_chapters"."id";
SELECT setval('"ywtrans"."project_chapters_id_seq"', 3, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "ywtrans"."project_tags_id_seq"
    OWNED BY "ywtrans"."project_tags"."id";
SELECT setval('"ywtrans"."project_tags_id_seq"', 2, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "ywtrans"."projects_id_seq"
    OWNED BY "ywtrans"."projects"."id";
SELECT setval('"ywtrans"."projects_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
SELECT setval('"ywtrans"."tags_id_seq"', 3, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
SELECT setval('"ywtrans"."users_id_seq"', 10001, true);

-- ----------------------------
-- Uniques structure for table chapter_assignments
-- ----------------------------
ALTER TABLE "ywtrans"."chapter_assignments" ADD CONSTRAINT "uq_chapter_user_role" UNIQUE ("chapter_id", "user_id", "role");

-- ----------------------------
-- Primary Key structure for table chapter_assignments
-- ----------------------------
ALTER TABLE "ywtrans"."chapter_assignments" ADD CONSTRAINT "chapter_assignments_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table chapter_work_images
-- ----------------------------
CREATE INDEX "idx_chapter_work_images_lookup" ON "ywtrans"."chapter_work_images" USING btree (
    "chapter_id" "pg_catalog"."int8_ops" ASC NULLS LAST,
    "role" "pg_catalog"."int2_ops" ASC NULLS LAST,
    "order_index" "pg_catalog"."int4_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Uniques structure for table chapter_work_images
-- ----------------------------
ALTER TABLE "ywtrans"."chapter_work_images" ADD CONSTRAINT "uq_chapter_work_images_role_order" UNIQUE ("chapter_id", "role", "order_index");

-- ----------------------------
-- Checks structure for table chapter_work_images
-- ----------------------------
ALTER TABLE "ywtrans"."chapter_work_images" ADD CONSTRAINT "chk_chapter_work_images_role" CHECK (role = ANY (ARRAY[1, 4, 5]));

-- ----------------------------
-- Primary Key structure for table chapter_work_images
-- ----------------------------
ALTER TABLE "ywtrans"."chapter_work_images" ADD CONSTRAINT "chapter_work_images_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table chapter_work_logs
-- ----------------------------
ALTER TABLE "ywtrans"."chapter_work_logs" ADD CONSTRAINT "chapter_work_logs_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table chapter_works
-- ----------------------------
ALTER TABLE "ywtrans"."chapter_works" ADD CONSTRAINT "chapter_works_chapter_id_key" UNIQUE ("chapter_id");

-- ----------------------------
-- Checks structure for table chapter_works
-- ----------------------------
ALTER TABLE "ywtrans"."chapter_works" ADD CONSTRAINT "chk_chapter_works_stage" CHECK (stage >= 1 AND stage <= 5);

-- ----------------------------
-- Primary Key structure for table chapter_works
-- ----------------------------
ALTER TABLE "ywtrans"."chapter_works" ADD CONSTRAINT "chapter_works_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table messages
-- ----------------------------
CREATE INDEX "idx_messages_created" ON "ywtrans"."messages" USING btree (
    "created_at" "pg_catalog"."timestamptz_ops" DESC NULLS FIRST
    );
CREATE INDEX "idx_messages_user_read" ON "ywtrans"."messages" USING btree (
    "user_id" "pg_catalog"."int8_ops" ASC NULLS LAST,
    "is_read" "pg_catalog"."bool_ops" ASC NULLS LAST
    );
CREATE INDEX "idx_messages_user_type" ON "ywtrans"."messages" USING btree (
    "user_id" "pg_catalog"."int8_ops" ASC NULLS LAST,
    "type" "pg_catalog"."int2_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Checks structure for table messages
-- ----------------------------
ALTER TABLE "ywtrans"."messages" ADD CONSTRAINT "chk_messages_type" CHECK (type >= 1 AND type <= 5);

-- ----------------------------
-- Primary Key structure for table messages
-- ----------------------------
ALTER TABLE "ywtrans"."messages" ADD CONSTRAINT "messages_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table project_assignments
-- ----------------------------
ALTER TABLE "ywtrans"."project_assignments" ADD CONSTRAINT "uq_project_user_role" UNIQUE ("project_id", "user_id", "role");

-- ----------------------------
-- Primary Key structure for table project_assignments
-- ----------------------------
ALTER TABLE "ywtrans"."project_assignments" ADD CONSTRAINT "project_assignments_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table project_chapters
-- ----------------------------
ALTER TABLE "ywtrans"."project_chapters" ADD CONSTRAINT "uq_project_chapter" UNIQUE ("project_id", "chapter_name");

-- ----------------------------
-- Primary Key structure for table project_chapters
-- ----------------------------
ALTER TABLE "ywtrans"."project_chapters" ADD CONSTRAINT "project_chapters_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table project_tags
-- ----------------------------
ALTER TABLE "ywtrans"."project_tags" ADD CONSTRAINT "uq_project_tag" UNIQUE ("project_id", "tag_id");

-- ----------------------------
-- Primary Key structure for table project_tags
-- ----------------------------
ALTER TABLE "ywtrans"."project_tags" ADD CONSTRAINT "project_tags_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table projects
-- ----------------------------
ALTER TABLE "ywtrans"."projects" ADD CONSTRAINT "projects_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table tags
-- ----------------------------
ALTER TABLE "ywtrans"."tags" ADD CONSTRAINT "tags_name_key" UNIQUE ("name");

-- ----------------------------
-- Primary Key structure for table tags
-- ----------------------------
ALTER TABLE "ywtrans"."tags" ADD CONSTRAINT "tags_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table users
-- ----------------------------
ALTER TABLE "ywtrans"."users" ADD CONSTRAINT "users_qq_openid_key" UNIQUE ("qq_openid");
ALTER TABLE "ywtrans"."users" ADD CONSTRAINT "users_username_key" UNIQUE ("username");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "ywtrans"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table chapter_assignments
-- ----------------------------
ALTER TABLE "ywtrans"."chapter_assignments" ADD CONSTRAINT "fk_chapter" FOREIGN KEY ("chapter_id") REFERENCES "ywtrans"."project_chapters" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "ywtrans"."chapter_assignments" ADD CONSTRAINT "fk_chapter_assignments_assigned_by" FOREIGN KEY ("assigned_by") REFERENCES "ywtrans"."users" ("id") ON DELETE SET NULL ON UPDATE NO ACTION;
ALTER TABLE "ywtrans"."chapter_assignments" ADD CONSTRAINT "fk_user" FOREIGN KEY ("user_id") REFERENCES "ywtrans"."users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table chapter_work_images
-- ----------------------------
ALTER TABLE "ywtrans"."chapter_work_images" ADD CONSTRAINT "fk_chapter_work_images_chapter" FOREIGN KEY ("chapter_id") REFERENCES "ywtrans"."project_chapters" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table chapter_work_logs
-- ----------------------------
ALTER TABLE "ywtrans"."chapter_work_logs" ADD CONSTRAINT "fk_chapter_work_logs_chapter" FOREIGN KEY ("chapter_id") REFERENCES "ywtrans"."project_chapters" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "ywtrans"."chapter_work_logs" ADD CONSTRAINT "fk_chapter_work_logs_user" FOREIGN KEY ("user_id") REFERENCES "ywtrans"."users" ("id") ON DELETE SET NULL ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table chapter_works
-- ----------------------------
ALTER TABLE "ywtrans"."chapter_works" ADD CONSTRAINT "fk_chapter_works_chapter" FOREIGN KEY ("chapter_id") REFERENCES "ywtrans"."project_chapters" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table messages
-- ----------------------------
ALTER TABLE "ywtrans"."messages" ADD CONSTRAINT "fk_messages_sender" FOREIGN KEY ("sender_id") REFERENCES "ywtrans"."users" ("id") ON DELETE SET NULL ON UPDATE NO ACTION;
ALTER TABLE "ywtrans"."messages" ADD CONSTRAINT "fk_messages_user" FOREIGN KEY ("user_id") REFERENCES "ywtrans"."users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table project_assignments
-- ----------------------------
ALTER TABLE "ywtrans"."project_assignments" ADD CONSTRAINT "fk_project" FOREIGN KEY ("project_id") REFERENCES "ywtrans"."projects" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "ywtrans"."project_assignments" ADD CONSTRAINT "fk_project_assignments_assigned_by" FOREIGN KEY ("assigned_by") REFERENCES "ywtrans"."users" ("id") ON DELETE SET NULL ON UPDATE NO ACTION;
ALTER TABLE "ywtrans"."project_assignments" ADD CONSTRAINT "fk_user" FOREIGN KEY ("user_id") REFERENCES "ywtrans"."users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table project_chapters
-- ----------------------------
ALTER TABLE "ywtrans"."project_chapters" ADD CONSTRAINT "fk_project" FOREIGN KEY ("project_id") REFERENCES "ywtrans"."projects" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table project_tags
-- ----------------------------
ALTER TABLE "ywtrans"."project_tags" ADD CONSTRAINT "fk_project" FOREIGN KEY ("project_id") REFERENCES "ywtrans"."projects" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "ywtrans"."project_tags" ADD CONSTRAINT "fk_tag" FOREIGN KEY ("tag_id") REFERENCES "ywtrans"."tags" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
