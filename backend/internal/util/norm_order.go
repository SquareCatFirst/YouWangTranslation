package util

import (
	"fmt"
	"sort"

	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/config"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/model"
	"gorm.io/gorm"
)

func NormChapterWorkImages(chapterID int64) int {
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		rows := []model.ChapterWorkImage{}
		if err := tx.
			Where("chapter_id = ?", chapterID).
			Order("order_index asc, id asc").
			Find(&rows).Error; err != nil {
			return err
		}

		if len(rows) == 0 {
			return nil
		}

		uniqueOrder := make([]int, 0, len(rows))
		seen := make(map[int]struct{}, len(rows))
		for _, row := range rows {
			if _, ok := seen[row.OrderIndex]; ok {
				continue
			}
			seen[row.OrderIndex] = struct{}{}
			uniqueOrder = append(uniqueOrder, row.OrderIndex)
		}

		sort.Ints(uniqueOrder)
		orderMap := make(map[int]int, len(uniqueOrder))
		for i, oldVal := range uniqueOrder {
			orderMap[oldVal] = i + 1
		}

		// 两阶段更新避免 (chapter_id, role, order_index) 唯一约束在交换/重排时中途冲突
		for i, row := range rows {
			tmp := -(i + 1)
			if err := tx.Model(&model.ChapterWorkImage{}).
				Where("id = ?", row.ID).
				Update("order_index", tmp).Error; err != nil {
				return err
			}
		}

		for _, row := range rows {
			newVal, ok := orderMap[row.OrderIndex]
			if !ok {
				return fmt.Errorf("order_index map missing: %d", row.OrderIndex)
			}
			if err := tx.Model(&model.ChapterWorkImage{}).
				Where("id = ?", row.ID).
				Update("order_index", newVal).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return 0
	}
	return 1
}

func NormProjChapter(projID int64) int {
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		rows := []model.ProjectChapter{}
		if err := tx.
			Where("project_id = ?", projID).
			Order("order_index asc, id asc").
			Find(&rows).Error; err != nil {
			return err
		}

		if len(rows) == 0 {
			return nil
		}

		uniqueOrder := make([]int, 0, len(rows))
		seen := make(map[int]struct{}, len(rows))
		for _, row := range rows {
			if _, ok := seen[row.OrderIndex]; ok {
				continue
			}
			seen[row.OrderIndex] = struct{}{}
			uniqueOrder = append(uniqueOrder, row.OrderIndex)
		}

		sort.Ints(uniqueOrder)
		orderMap := make(map[int]int, len(uniqueOrder))
		for i, oldVal := range uniqueOrder {
			orderMap[oldVal] = i + 1
		}

		for i, row := range rows {
			tmp := -(i + 1)
			if err := tx.Model(&model.ProjectChapter{}).
				Where("id = ?", row.ID).
				Update("order_index", tmp).Error; err != nil {
				return err
			}
		}

		for _, row := range rows {
			newVal, ok := orderMap[row.OrderIndex]
			if !ok {
				return fmt.Errorf("order_index map missing: %d", row.OrderIndex)
			}
			if err := tx.Model(&model.ProjectChapter{}).
				Where("id = ?", row.ID).
				Update("order_index", newVal).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return 0
	}
	return 1
}
