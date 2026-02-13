package util

import (
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/config"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/model"
)

func ChapterAssignmentInsert(chapterid int64, userid int64, role int, assignedby int64) bool {
	row := model.ChapterAssignment{
		ChapterID:  chapterid,
		UserID:     userid,
		Role:       role,
		AssignedBy: &assignedby,
	}

	if config.DB.Create(&row).Error != nil {
		return false
	}
	return true
}

func ChapterWorkImageInsert(chapterid int64, role int16, imageurl string) bool {
	var maxOrderIndex int
	if config.DB.Model(&model.ChapterWorkImage{}).
		Where("chapter_id = ?", chapterid).
		Select("COALESCE(MAX(order_index), 0)").
		Scan(&maxOrderIndex).Error != nil {
		return false
	}

	row := model.ChapterWorkImage{
		ChapterID:  chapterid,
		Role:       role,
		ImageURL:   imageurl,
		OrderIndex: maxOrderIndex + 1,
	}

	if config.DB.Create(&row).Error != nil {
		return false
	}
	return true
}

func ChapterWorkLogInsert(chapterid int64, role *int16, actiontype string, userid *int64, note *string) bool {
	var maxOrderIndex int
	if config.DB.Model(&model.ChapterWorkLog{}).
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

	if config.DB.Create(&row).Error != nil {
		return false
	}
	return true
}

func ChapterWorkInsert(chapterid int64, translationtext *string, proofreadtext *string, stage int16) bool {
	row := model.ChapterWork{
		ChapterID:       chapterid,
		TranslationText: translationtext,
		ProofreadText:   proofreadtext,
		Stage:           stage,
	}

	if config.DB.Create(&row).Error != nil {
		return false
	}
	return true
}

func MessageInsert(userid int64, senderid *int64, msgtype int16, title *string, content *string, reftype *string, refid *int64, isread bool) bool {
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

	if config.DB.Create(&row).Error != nil {
		return false
	}
	return true
}

func ProjectAssignmentInsert(projectid int64, userid int64, role int, assignedby *int64) bool {
	row := model.ProjectAssignment{
		ProjectID:  projectid,
		UserID:     userid,
		Role:       role,
		AssignedBy: assignedby,
	}

	if config.DB.Create(&row).Error != nil {
		return false
	}
	return true
}

func ProjectChapterInsert(projectid int64, chaptername string, isvisible bool) bool {
	var maxOrderIndex int
	if config.DB.Model(&model.ProjectChapter{}).
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

	if config.DB.Create(&row).Error != nil {
		return false
	}
	return true
}

func ProjectTagInsert(projectid int64, tagid int64) bool {
	row := model.ProjectTag{
		ProjectID: projectid,
		TagID:     tagid,
	}

	if config.DB.Create(&row).Error != nil {
		return false
	}
	return true
}

func ProjectInsert(coverurl *string, name string, translatedname *string, author *string, sourcesite *string, description *string, translationdescription *string, status int, sourcemax int, translatormax int, proofreadermax int, typesettermax int, supervisormax int, adminmax int) bool {
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

	if config.DB.Create(&row).Error != nil {
		return false
	}
	return true
}

func TagInsert(name string, color string) bool {
	row := model.Tag{
		Name:  name,
		Color: color,
	}

	if config.DB.Create(&row).Error != nil {
		return false
	}
	return true
}

func UserInsert(avatarurl *string, nickname string, role int, qqopenid *string, status int, username *string, passwordhash *string) bool {
	row := model.User{
		AvatarURL:    avatarurl,
		Nickname:     nickname,
		Role:         role,
		QQOpenID:     qqopenid,
		Status:       status,
		Username:     username,
		PasswordHash: passwordhash,
	}

	if config.DB.Create(&row).Error != nil {
		return false
	}
	return true
}
