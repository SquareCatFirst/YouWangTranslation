ALTER TABLE "ywtrans"."chapter_work_images"
DROP CONSTRAINT IF EXISTS "chk_chapter_work_images_role";

ALTER TABLE "ywtrans"."chapter_work_images"
ADD CONSTRAINT "chk_chapter_work_images_role"
CHECK (role >= 0 AND role <= 5);
