package sprinklesv1

import (
	"context"

	gorm "gorm.io/gorm"
	clause "gorm.io/gorm/clause"
)

func DeleteGroupsByNames(ctx context.Context, tx *gorm.DB, names []string, force_cascade bool) (err error) {
	if len(names) > 0 {
		var models GroupGormModels
		statement := tx.Preload(clause.Associations)
		if err = statement.Where("name in ?", names).Delete(&models).Error; err != nil {
			return
		}
	}
	return
}
