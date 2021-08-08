package store

import (
	"context"
	"log"
	"os"

	"github.com/noodles623/csp/errors"
	"github.com/noodles623/csp/objects"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type pg struct {
	db *gorm.DB
}

func NewPostgresAssetStore(conn string) AssetStore {
	db, err := gorm.Open(postgres.Open(conn),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "", log.LstdFlags),
				logger.Config{
					LogLevel: logger.Info,
					Colorful: true,
				},
			),
		},
	)
	if err != nil {
		panic("Unable to connect to database: " + err.Error())
	}
	if err := db.AutoMigrate(&objects.Asset{}); err != nil {
		panic("Unable to migrate database: " + err.Error())
	}
	return &pg{db: db}
}

func (p *pg) Get(ctx context.Context, in *objects.GetRequest) (*objects.Asset, error) {
	asst := &objects.Asset{}
	err := p.db.WithContext(ctx).Take(asst, "id = ?", in.Id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrAssetNotFound
	}
	return asst, err
}

func (p *pg) List(ctx context.Context, in *objects.ListRequest) ([]*objects.Asset, error) {
	if in.Limit == 0 || in.Limit > objects.MaxListLimit {
		in.Limit = objects.MaxListLimit
	}
	query := p.db.WithContext(ctx).Limit(in.Limit)
	list := make([]*objects.Asset, 0, in.Limit)
	err := query.Order("id").Find(&list).Error
	return list, err
}

func (p *pg) Create(ctx context.Context, in *objects.CreateRequest) error {
	if in.Asset == nil {
		return errors.ErrObjectIsRequired
	}
	in.Asset.ID = GenerateUniqueID()
	return p.db.WithContext(ctx).Create(in.Asset).Error
}

func (p *pg) Delete(ctx context.Context, in *objects.DeleteRequest) error {
	asst := &objects.Asset{ID: in.Id}
	return p.db.WithContext(ctx).Model(asst).Delete(asst).Error
}
