package util

import (
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/model"
	"gorm.io/gorm"
)

func ChapterAssignmentInsert(tx *gorm.DB, chapterid int64, userid int64, role int, assignedby int64) bool {
	row := model.ChapterAssignment{
		ChapterID:  chapterid,
		UserID:     userid,
		Role:       role,
		AssignedBy: &assignedby,
	}

	if tx.Create(&row).Error != nil {
		return false
	}
	return true
}

func ChapterWorkImageInsert(tx *gorm.DB, chapterid int64, role int16, imageurl string) error {
	var maxOrderIndex int
	if err := tx.Model(&model.ChapterWorkImage{}).
		Where("chapter_id = ?", chapterid).
		Select("COALESCE(MAX(order_index), 0)").
		Scan(&maxOrderIndex).Error; err != nil {
		return err
	}

	row := model.ChapterWorkImage{
		ChapterID:  chapterid,
		Role:       role,
		ImageURL:   imageurl,
		OrderIndex: maxOrderIndex + 1,
	}

	if err := tx.Create(&row).Error; err != nil {
		return err
	}
	return nil
}

func ChapterWorkLogInsert(tx *gorm.DB, chapterid int64, role *int16, actiontype string, userid *int64, note *string) bool {
	var maxOrderIndex int
	if tx.Model(&model.ChapterWorkLog{}).
		Where("chapter_id = ?", chapterid).
		Select("COALESCE(MAX(order_index), 0)").
		Scan(&maxOrderIndex).Error != nil {
		return false
	}

	row := model.ChapterWorkLog{
		ChapterID:  chapterid,
		OrderIndex: maxOrderIndex + 1,
		Role:       role,
		ActionType: actiontype,
		UserID:     userid,
		Note:       note,
	}

	if tx.Create(&row).Error != nil {
		return false
	}
	return true
}

func ChapterWorkInsert(tx *gorm.DB, chapterid int64, translationtext *string, proofreadtext *string, stage int16) bool {
	row := model.ChapterWork{
		ChapterID:       chapterid,
		TranslationText: translationtext,
		ProofreadText:   proofreadtext,
		Stage:           stage,
	}

	if tx.Create(&row).Error != nil {
		return false
	}
	return true
}

func MessageInsert(tx *gorm.DB, userid int64, senderid *int64, msgtype int16, title *string, content *string, reftype *string, refid *int64, isread bool) bool {
	row := model.Message{
		UserID:   userid,
		SenderID: senderid,
		Type:     msgtype,
		Title:    title,
		Content:  content,
		RefType:  reftype,
		RefID:    refid,
		IsRead:   isread,
	}

	if tx.Create(&row).Error != nil {
		return false
	}
	return true
}

func ProjectAssignmentInsert(tx *gorm.DB, projectid int64, userid int64, role int, assignedby *int64) bool {
	row := model.ProjectAssignment{
		ProjectID:  projectid,
		UserID:     userid,
		Role:       role,
		AssignedBy: assignedby,
	}

	if tx.Create(&row).Error != nil {
		return false
	}
	return true
}

func ProjectChapterInsert(tx *gorm.DB, projectid int64, chaptername string, isvisible bool) bool {
	var maxOrderIndex int
	if tx.Model(&model.ProjectChapter{}).
		Where("project_id = ?", projectid).
		Select("COALESCE(MAX(order_index), 0)").
		Scan(&maxOrderIndex).Error != nil {
		return false
	}

	row := model.ProjectChapter{
		ProjectID:   projectid,
		ChapterName: chaptername,
		IsVisible:   isvisible,
		OrderIndex:  maxOrderIndex + 1,
	}

	if tx.Create(&row).Error != nil {
		return false
	}
	return true
}

func ProjectTagInsert(tx *gorm.DB, projectid int64, tagid int64) bool {
	row := model.ProjectTag{
		ProjectID: projectid,
		TagID:     tagid,
	}

	if tx.Create(&row).Error != nil {
		return false
	}
	return true
}

func ProjectInsert(tx *gorm.DB, coverurl *string, name string, translatedname *string, author *string, sourcesite *string, description *string, translationdescription *string, status int, sourcemax int, translatormax int, proofreadermax int, typesettermax int, supervisormax int, adminmax int) bool {
	row := model.Project{
		CoverURL:               coverurl,
		Name:                   name,
		TranslatedName:         translatedname,
		Author:                 author,
		SourceSite:             sourcesite,
		Description:            description,
		TranslationDescription: translationdescription,
		Status:                 status,
		SourceMax:              sourcemax,
		TranslatorMax:          translatormax,
		ProofreaderMax:         proofreadermax,
		TypesetterMax:          typesettermax,
		SupervisorMax:          supervisormax,
		AdminMax:               adminmax,
	}

	if tx.Create(&row).Error != nil {
		return false
	}
	return true
}

func TagInsert(tx *gorm.DB, name string, color string) bool {
	row := model.Tag{
		Name:  name,
		Color: color,
	}

	if tx.Create(&row).Error != nil {
		return false
	}
	return true
}

func UserInsert(tx *gorm.DB, avatarurl *string, nickname string, role int, qqopenid *string, status int, username *string, passwordhash *string) bool {
	row := model.User{
		AvatarURL:    avatarurl,
		Nickname:     nickname,
		Role:         role,
		QQOpenID:     qqopenid,
		Status:       status,
		Username:     username,
		PasswordHash: passwordhash,
	}

	if tx.Create(&row).Error != nil {
		return false
	}
	return true
}
