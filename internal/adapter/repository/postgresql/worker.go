package postgresql

import (
	"context"
	"strings"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/gofiber/fiber/v2"
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

func (r *WorkerRepositoryClient) GetWorkers(ctx context.Context, filter dto.GetWorkersFilter) ([]domain.Worker, error) {
	var workersDB []WorkerDB

	queryBuilder := r.q.Debug().
		WithContext(ctx).
		Preload("WorkerServices").
		Preload("Tag").
		Preload("Reviews").
		Model(&domain.Worker{}).
		Joins("JOIN worker_services ON worker_services.worker_id = workers.id").
		Joins("JOIN tags ON tags.id = workers.tag_id").
		Select("DISTINCT workers.*, (SELECT MIN(price) FROM worker_services WHERE worker_services.worker_id = workers.id) AS lower_price, (SELECT AVG(rating) FROM reviews WHERE reviews.worker_id = workers.id ) AS rating, (SELECT COUNT(*) FROM reviews WHERE reviews.worker_id = workers.id) AS total_rating")

	if filter.Keyword != "" {
		lowerKeyword := "%" + strings.ToLower(filter.Keyword) + "%"
		queryBuilder = queryBuilder.Where(
			"LOWER(workers.description) LIKE ? OR LOWER(worker_services.service) LIKE ? OR LOWER(tags.tag) LIKE ? OR LOWER(tags.specialization) LIKE ?",
			lowerKeyword, lowerKeyword, lowerKeyword, lowerKeyword,
		)
	}

	if err := queryBuilder.Find(&workersDB).Error; err != nil {
		return nil, errx.New(fiber.StatusInternalServerError, "failed to get workers", err)
	}

	var workers []domain.Worker
	for _, worker := range workersDB {
		workers = append(workers, worker.format())
	}

	return workers, nil
}

func (r *WorkerRepositoryClient) GetWorkersByWorkerIDs(ctx context.Context, workerIDs []string) ([]domain.Worker, error) {
	var workersDB []WorkerDB

	if err := r.q.Debug().
		WithContext(ctx).
		Preload("WorkerServices").
		Preload("Tag").
		Preload("Reviews").
		Model(&domain.Worker{}).
		Where("id IN ?", workerIDs).
		Select("DISTINCT workers.*, (SELECT MIN(price) FROM worker_services WHERE worker_services.worker_id = workers.id) AS lower_price, (SELECT AVG(rating) FROM reviews WHERE reviews.worker_id = workers.id ) AS rating, (SELECT COUNT(*) FROM reviews WHERE reviews.worker_id = workers.id) AS total_rating").
		Find(&workersDB).Error; err != nil {
		return nil, errx.New(fiber.StatusInternalServerError, "failed to get workers", err)
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
		Select("DISTINCT workers.*, (SELECT MIN(price) FROM worker_services WHERE worker_services.worker_id = workers.id) AS lower_price, (SELECT AVG(rating) FROM reviews WHERE reviews.worker_id = workers.id ) AS rating, (SELECT COUNT(*) FROM reviews WHERE reviews.worker_id = workers.id) AS total_rating, (SELECT COUNT(*) FROM reviews WHERE reviews.worker_id = workers.id AND reviews.review != '') AS total_review").
		Where("id = ?", workerID).
		First(&worker).Error; err != nil {
		return domain.Worker{}, errx.New(fiber.StatusInternalServerError, "failed to get worker", err)
	}

	return worker.format(), nil
}

func (r *WorkerRepositoryClient) GetWorkersForBotResponse(ctx context.Context, keywords []string) ([]domain.Worker, error) {
	var workersDB []WorkerDB

	query := r.q.Debug().
		WithContext(ctx).
		Preload("WorkerServices").
		Preload("Tag").
		Preload("Reviews").
		Model(&domain.Worker{}).
		Joins("JOIN worker_services ON worker_services.worker_id = workers.id").
		Joins("JOIN tags ON tags.id = workers.tag_id")

	for _, keyword := range keywords {
		lowerKeyword := "%" + strings.ToLower(keyword) + "%"
		query = query.Where(
			"LOWER(workers.description) LIKE ? OR LOWER(worker_services.service) LIKE ? OR LOWER(tags.tag) LIKE ? OR LOWER(tags.specialization) LIKE ?",
			lowerKeyword, lowerKeyword, lowerKeyword, lowerKeyword,
		)
	}

	if err := query.
		Select(`DISTINCT workers.*, 
			(SELECT MIN(price) FROM worker_services WHERE worker_services.worker_id = workers.id) AS lower_price, 
			(SELECT AVG(rating) FROM reviews WHERE reviews.worker_id = workers.id) AS rating, 
			(SELECT COUNT(*) FROM reviews WHERE reviews.worker_id = workers.id) AS total_rating`).
		Order("rating DESC").
		Find(&workersDB).Error; err != nil {
		return nil, errx.New(fiber.StatusInternalServerError, "failed to get workers", err)
	}

	var workers []domain.Worker
	for _, worker := range workersDB {
		workers = append(workers, worker.format())
	}

	return workers, nil
}
