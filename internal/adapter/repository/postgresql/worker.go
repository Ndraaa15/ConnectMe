package postgresql

import (
	"context"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// this is needed because, there was lower price that i got from query so i need to specify the column name
// and also i don't want that table also migrating
type WorkerDB struct {
	ID             string                 `gorm:"type:varchar(36);primaryKey"`
	Name           string                 `gorm:"type:varchar(255)"`
	TagID          uint64                 `gorm:"type:integer"`
	Tag            domain.Tag             `gorm:"references:ID;foreignKey:TagID"`
	Description    string                 `gorm:"type:text"`
	WorkExperience uint64                 `gorm:"type:integer"`
	LowerPrice     float64                `gorm:"column:lower_price"`
	WorkerServices []domain.WorkerService `gorm:"references:ID;foreignKey:WorkerID;constraint:OnDelete:CASCADE"`
	Image          string                 `gorm:"type:text"`
	WorkHour       pq.StringArray         `gorm:"type:varchar(5)[]"`
	Rating         float64                `gorm:"column:rating"`
	TotalRating    uint64                 `gorm:"column:total_rating"`
	TotalReview    uint64                 `gorm:"column:total_review"`
	Reviews        []domain.Review        `gorm:"references:ID;foreignKey:WorkerID;constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time              `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt      time.Time              `gorm:"type:timestamp;autoUpdateTime"`
}

func (dbData *WorkerDB) format() domain.Worker {
	return domain.Worker{
		ID:             dbData.ID,
		Name:           dbData.Name,
		TagID:          dbData.TagID,
		Tag:            dbData.Tag,
		Description:    dbData.Description,
		WorkExperience: dbData.WorkExperience,
		LowerPrice:     dbData.LowerPrice,
		WorkerServices: dbData.WorkerServices,
		Image:          dbData.Image,
		WorkHour:       dbData.WorkHour,
		Rating:         dbData.Rating,
		TotalRating:    dbData.TotalRating,
		TotalReview:    dbData.TotalReview,
		Reviews:        dbData.Reviews,
		CreatedAt:      dbData.CreatedAt,
		UpdatedAt:      dbData.UpdatedAt,
	}
}

func NewWorkerRepository(db *gorm.DB) *WorkerRepository {
	return &WorkerRepository{
		db: db,
	}
}

type WorkerRepository struct {
	db *gorm.DB
}

func (r *WorkerRepository) NewWorkerRepositoryClient(tx bool) port.WorkerRepositoryClientItf {
	if tx {
		return &WorkerRepositoryClient{
			q: r.db.Begin(),
		}
	} else {
		return &WorkerRepositoryClient{
			q: r.db,
		}
	}
}

type WorkerRepositoryClient struct {
	q *gorm.DB
}

func (r *WorkerRepositoryClient) Commit() error {
	return r.q.Commit().Error
}

func (r *WorkerRepositoryClient) Rollback() error {
	return r.q.Rollback().Error
}

func (r *WorkerRepositoryClient) GetWorkers(ctx context.Context) ([]domain.Worker, error) {
	var workersDB []WorkerDB

	if err := r.q.Debug().
		WithContext(ctx).
		Preload("WorkerServices").
		Preload("Tag").
		Preload("Reviews").
		Model(&domain.Worker{}).
		Select("workers.*, (SELECT MIN(price) FROM worker_services WHERE worker_services.worker_id = workers.id) AS lower_price, (SELECT AVG(rating) FROM reviews WHERE reviews.worker_id = workers.id ) AS rating, (SELECT COUNT(*) FROM reviews WHERE reviews.worker_id = workers.id) AS total_rating").
		Find(&workersDB).Error; err != nil {
		return nil, err
	}

	var workers []domain.Worker
	for _, worker := range workersDB {
		workers = append(workers, worker.format())
	}

	return workers, nil
}

func (r *WorkerRepositoryClient) GetWorker(ctx context.Context, workerID string) (domain.Worker, error) {
	var worker WorkerDB

	if err := r.q.Debug().
		WithContext(ctx).
		Preload("WorkerServices").
		Preload("Reviews.User").
		Preload("Tag").
		Model(&domain.Worker{}).
		Select("workers.*, (SELECT MIN(price) FROM worker_services WHERE worker_services.worker_id = workers.id) AS lower_price, (SELECT AVG(rating) FROM reviews WHERE reviews.worker_id = workers.id ) AS rating, (SELECT COUNT(*) FROM reviews WHERE reviews.worker_id = workers.id) AS total_rating, (SELECT COUNT(*) FROM reviews WHERE reviews.worker_id = workers.id AND reviews.description != '') AS total_review").
		Where("id = ?", workerID).
		First(&worker).Error; err != nil {
		return domain.Worker{}, err
	}

	return worker.format(), nil
}
