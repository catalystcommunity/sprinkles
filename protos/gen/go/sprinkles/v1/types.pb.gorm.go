// Code generated by protoc-gen-go-gorm. DO NOT EDIT.
// source: sprinkles/v1/types.proto

package sprinklesv1

import (
	context "context"
	uuid "github.com/google/uuid"
	lo "github.com/samber/lo"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	gorm "gorm.io/gorm"
	clause "gorm.io/gorm/clause"
	time "time"
)

type HelloGormModels []*HelloGormModel
type HelloProtos []*Hello
type HelloGormModel struct {

	// @gotags: fake:"skip"
	Id *string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"id" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;" json:"created_at" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;" json:"updated_at" fake:"skip"`

	// @gotags: fake:"{number:1,1}"
	HelloType int `json:"hello_type" fake:"{number:1,1}"`

	// @gotags: fake:"{beername}"
	PersonName *string `json:"person_name" fake:"{beername}"`
}

func (m *HelloGormModel) TableName() string {
	return "hellos"
}

func (m HelloGormModels) ToProtos() (protos HelloProtos, err error) {
	protos = HelloProtos{}
	for _, model := range m {
		var proto *Hello
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p HelloProtos) ToModels() (models HelloGormModels, err error) {
	models = HelloGormModels{}
	for _, proto := range p {
		var model *HelloGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

func (m *HelloGormModel) ToProto() (theProto *Hello, err error) {
	if m == nil {
		return
	}
	theProto = &Hello{}

	theProto.Id = m.Id

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	theProto.HelloType = HelloType(m.HelloType)

	theProto.PersonName = m.PersonName

	return
}

func (p *Hello) ToModel() (theModel *HelloGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &HelloGormModel{}

	theModel.Id = p.Id

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	theModel.HelloType = int(p.HelloType)

	theModel.PersonName = p.PersonName

	return
}

func (m HelloGormModels) GetByModelIds(ctx context.Context, tx *gorm.DB, preloads ...string) (err error) {
	ids := []string{}
	for _, model := range m {
		if model.Id != nil {
			ids = append(ids, *model.Id)
		}
	}
	if len(ids) > 0 {
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		err = statement.Where("id in ?", ids).Find(&m).Error
	}
	return
}

func (p *HelloProtos) Upsert(ctx context.Context, tx *gorm.DB, selects, omits []string, fullSaveAssociations bool, preloads ...string) (err error) {
	if p != nil {
		omitMap := map[string]bool{}
		for _, omit := range omits {
			omitMap[omit] = true
		}
		creates, updates := []*HelloGormModel{}, []*HelloGormModel{}
		nilUid := uuid.Nil.String()
		var model *HelloGormModel
		for _, proto := range *p {
			if model, err = proto.ToModel(); err != nil {
				return
			} else {
				if model.Id != nil && *model.Id != "" && *model.Id != nilUid {
					updates = append(updates, model)
				} else {
					creates = append(creates, model)
				}
			}
		}
		session := tx.Select(selects).Omit(omits...).Session(&gorm.Session{FullSaveAssociations: fullSaveAssociations})
		if len(creates) > 0 {
			if err = session.Create(&creates).Error; err != nil {
				return
			}
		}
		if len(updates) > 0 {
			toSave := []*HelloGormModel{}
			for _, update := range updates {
				thing := &HelloGormModel{}
				*thing = *update
				toSave = append(toSave, thing)
			}
			if err = session.Save(&toSave).Error; err != nil {
				return
			}
		}
		models := HelloGormModels{}
		models = append(creates, updates...)
		if err = models.GetByModelIds(ctx, tx, preloads...); err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = HelloProtos{}
		}
	}
	return
}

func (p *HelloProtos) List(ctx context.Context, tx *gorm.DB, limit, offset int, order interface{}, preloads ...string) (err error) {
	if p != nil {
		var models HelloGormModels
		statement := tx.Preload(clause.Associations).Limit(limit).Offset(offset)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if order != nil {
			statement = statement.Order(order)
		}
		if err = statement.Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = HelloProtos{}
		}
	}
	return
}

func (p *HelloProtos) GetByIds(ctx context.Context, tx *gorm.DB, ids []string, preloads ...string) (err error) {
	if p != nil {
		var models HelloGormModels
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if err = statement.Where("id in ?", ids).Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = HelloProtos{}
		}
	}
	return
}

func DeleteHelloGormModels(ctx context.Context, tx *gorm.DB, ids []string) error {
	statement := tx.Where("id in ?", ids)
	return statement.Delete(&HelloGormModel{}).Error
}