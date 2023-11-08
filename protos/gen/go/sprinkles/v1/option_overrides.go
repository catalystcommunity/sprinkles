package sprinklesv1

import (
	"context"

	gorm "gorm.io/gorm"
	clause "gorm.io/gorm/clause"
)

func (p *OptionOverrideProtos) GetOptionsByGroup(ctx context.Context, tx *gorm.DB, groups []string) (err error) {
	if p != nil {
		var models OptionOverrideGormModels
		statement := tx.Preload(clause.Associations)
		if err = statement.Where("group in ?", groups).Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = OptionOverrideProtos{}
		}
	}
	return
}

func (p *OptionOverrideProtos) GetOptionValue(ctx context.Context, tx *gorm.DB, ignoreGroups []string) (option_value string, err error) {
	var models OptionOverrideGormModels
	statement := tx.Preload(clause.Associations)
	if err = statement.Where("group not in ?", ignoreGroups).Find(&models).Error; err != nil {
		return
	}
	if len(models) > 0 {
		*p, err = models.ToProtos()
	} else {
		*p = OptionOverrideProtos{}
	}
	return
}
