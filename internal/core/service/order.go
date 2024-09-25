package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrderService struct {
	repository                    port.OrderRepositoryItf
	workerServiceRepositoryClient port.WorkerServiceRepositoryClientItf
	paymentRepositoryClient       port.PaymentRepositoryClientItf
	cache                         port.CacheItf
	paymentGateway                port.PaymentGatewayItf
}

func NewOrderService(repository port.OrderRepositoryItf, cache port.CacheItf, workerServiceRepository port.WorkerServiceRepositoryItf, paymentRepository port.PaymentRepositoryItf, paymentGateway port.PaymentGatewayItf) *OrderService {
	workerServiceRepositoryClient := workerServiceRepository.NewWorkerServiceRepositoryClient(false)
	paymentRepositoryClient := paymentRepository.NewPaymentRepositoryClient(false)

	orderService := &OrderService{
		repository:                    repository,
		cache:                         cache,
		workerServiceRepositoryClient: workerServiceRepositoryClient,
		paymentRepositoryClient:       paymentRepositoryClient,
		paymentGateway:                paymentGateway,
	}

	return orderService
}

func (order *OrderService) CreateOrder(ctx context.Context, req dto.CreateOrderRequest, userID string) (dto.PaymentResponse, error) {
	repositoryClient := order.repository.NewOrderRepositoryClient(true)

	var (
		err error
	)

	defer func() {
		if err != nil {
			if errTx := repositoryClient.Rollback(); errTx != nil {
				err = errTx
			}
		}
	}()

	workerServices, err := order.workerServiceRepositoryClient.GetWorkerServicesByWorkerServiceIDs(ctx, req.WorkerService)
	if err != nil {
		return dto.PaymentResponse{}, err
	}

	if len(workerServices) != len(req.WorkerService) {
		return dto.PaymentResponse{}, errx.New(fiber.StatusBadRequest, "worker service not found", errors.New("worker service not found"))
	}

	var totalWorkerServicePrice float64
	for _, workerService := range workerServices {
		totalWorkerServicePrice += workerService.Price
	}

	date, err := time.Parse("02 January 2006", req.Date)
	if err != nil {
		return dto.PaymentResponse{}, err
	}

	time, err := time.Parse("15:04", req.Time)
	if err != nil {
		return dto.PaymentResponse{}, err
	}

	orderData := domain.Order{
		OrderID:       uuid.New(),
		WorkerID:      uuid.MustParse(req.WorkerID),
		UserID:        uuid.MustParse(userID),
		Date:          date,
		Time:          time,
		WorkerService: req.WorkerService,
		OrderStatus:   domain.StatusOrderOnGoing,
	}

	orderData.Address = domain.AddressOrder{
		OrderID:       orderData.OrderID,
		Street:        req.OrderAddress.Street,
		Latitude:      req.OrderAddress.Latitude,
		Longitude:     req.OrderAddress.Longitude,
		AddressType:   req.OrderAddress.AddressType,
		DetailAddress: req.OrderAddress.DetailAddress,
	}

	paymentType, err := parsePaymentType(req.Payment.PaymentType)
	if err != nil {
		return dto.PaymentResponse{}, err
	}

	paymentData := domain.Payment{
		PaymentID:         uuid.New(),
		OrderID:           orderData.OrderID,
		ServiceFee:        5000,
		TotalServicePrice: totalWorkerServicePrice,
		TotalPrice:        totalWorkerServicePrice + 5000,
		PaymentType:       paymentType,
		PromoCode:         req.Payment.PromoCode,
		Status:            domain.StatusPaymentOnGoing,
	}

	errorChan := make(chan error, 2)
	defer close(errorChan)

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := repositoryClient.CreateOrder(ctx, &orderData); err != nil {
			errorChan <- err
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := order.paymentRepositoryClient.CreatePayment(ctx, &paymentData); err != nil {
			fmt.Println("error create payment")
			errorChan <- err
			cancel()
		}
	}()

	wg.Wait()

	select {
	case err = <-errorChan:
		return dto.PaymentResponse{}, err
	default:
	}

	if err := repositoryClient.Commit(); err != nil {
		return dto.PaymentResponse{}, err
	}

	res, err := order.paymentGateway.CreateTransaction(ctx, paymentData)
	if err != nil {
		return dto.PaymentResponse{}, err
	}

	return res, nil
}

func parsePaymentType(paymentType string) (domain.PaymentType, error) {
	switch paymentType {
	case "BCA Virtual Account":
		return domain.PaymentTypeBCAVa, nil
	case "BRI Virtual Account":
		return domain.PaymentTypeBRIVa, nil
	case "Mandiri Virtual Account":
		return domain.PaymentTypeMandiriVa, nil
	case "BNI Virtual Account":
		return domain.PaymentTypeBNIVa, nil
	case "Gopay":
		return domain.PaymentTypeGopay, nil
	case "ShopeePay":
		return domain.PaymentTypeGopay, nil
	default:
		return domain.PaymentTypeUnknown, errx.New(fiber.StatusBadRequest, "input payment type is unknown", errors.New("invalid payment type"))
	}
}
