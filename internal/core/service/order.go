package service

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/util"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrderService struct {
	repository           port.OrderRepositoryItf
	workerServiceService port.WorkerServiceServiceItf
	cache                port.CacheItf
	paymentGateway       port.PaymentGatewayItf
}

func NewOrderService(repository port.OrderRepositoryItf, cache port.CacheItf, workerServiceService port.WorkerServiceServiceItf, paymentGateway port.PaymentGatewayItf) *OrderService {
	orderService := &OrderService{
		repository:           repository,
		cache:                cache,
		workerServiceService: workerServiceService,
		paymentGateway:       paymentGateway,
	}

	return orderService
}

func (order *OrderService) CreateOrder(ctx context.Context, req dto.CreateOrderRequest, userID string) (dto.TransactionResponse, error) {
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

	workerServices, err := order.workerServiceService.GetWorkerServicesByWorkerServiceIDs(ctx, req.WorkerService)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	if len(workerServices) != len(req.WorkerService) {
		return dto.TransactionResponse{}, errx.New(fiber.StatusBadRequest, "worker service not found", errors.New("worker service not found"))
	}

	var totalWorkerServicePrice float64
	for _, workerService := range workerServices {
		totalWorkerServicePrice += workerService.Price
	}

	date, err := time.Parse("02 January 2006", req.Date)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	time, err := time.Parse("15:04", req.Time)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	orderData := domain.Order{
		OrderID:       util.GenerateOrderCode(),
		WorkerID:      req.WorkerID,
		UserID:        userID,
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
		return dto.TransactionResponse{}, err
	}

	orderData.Payment = domain.Payment{
		ID:                uuid.New().String(),
		OrderID:           orderData.OrderID,
		ServiceFee:        5000,
		TotalServicePrice: totalWorkerServicePrice,
		TotalPrice:        totalWorkerServicePrice + 5000,
		PaymentType:       paymentType,
		PromoCode:         req.Payment.PromoCode,
		Status:            domain.StatusPaymentOnGoing,
	}

	if err := repositoryClient.CreateOrder(ctx, &orderData); err != nil {
		return dto.TransactionResponse{}, err
	}

	if err := repositoryClient.Commit(); err != nil {
		return dto.TransactionResponse{}, err
	}

	res, err := order.paymentGateway.CreateTransaction(ctx, orderData.Payment)
	if err != nil {
		return dto.TransactionResponse{}, err
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

func (order *OrderService) GetOrders(ctx context.Context, userID string, filter dto.GetOrderFilter) ([]dto.OrderResponse, error) {
	repositoryClient := order.repository.NewOrderRepositoryClient(true)

	data, err := repositoryClient.GetOrdersByUserID(ctx, userID, filter)
	if err != nil {
		return nil, err
	}

	orders := make([]dto.OrderResponse, len(data))

	var wg sync.WaitGroup

	for i, order := range data {
		wg.Add(1)
		go func(i int, order domain.Order) {
			defer wg.Done()
			formatOrderResponse(&order, &orders[i])
		}(i, order)
	}

	wg.Wait()
	return orders, nil
}

func formatOrderResponse(order *domain.Order, orderResponse *dto.OrderResponse) {
	*orderResponse = dto.OrderResponse{
		OrderID:     order.OrderID,
		WorkerID:    order.WorkerID,
		StatusOrder: order.OrderStatus.String(),
		WorkerName:  order.Worker.Name,
		WorkerImage: order.Worker.Image,
		Tag: dto.TagResponse{
			ID:             order.Worker.Tag.ID,
			Tag:            order.Worker.Tag.Tag,
			Specialization: order.Worker.Tag.Specialization,
		},
		TotalPrice: order.Payment.TotalPrice,
		OrderDate:  order.Date.Format("02 January 2006"),
		OrderTime:  order.Time.Format("15:04"),
	}
}

func (s *OrderService) GetOrder(ctx context.Context, orderID string) (dto.OrderDetailResponse, error) {
	repositoryClient := s.repository.NewOrderRepositoryClient(false)

	orderData, err := repositoryClient.GetOrderByID(ctx, orderID)
	if err != nil {
		return dto.OrderDetailResponse{}, err
	}

	workerServiceData, err := s.workerServiceService.GetWorkerServicesByWorkerServiceIDs(ctx, orderData.WorkerService)
	if err != nil {
		return dto.OrderDetailResponse{}, err
	}

	var wg sync.WaitGroup
	workerServices := make([]dto.WorkerServiceResponse, len(orderData.WorkerService))
	for i, workerService := range workerServiceData {
		wg.Add(1)
		go func(i int, workerService domain.WorkerService) {
			defer wg.Done()
			formatWorkerServiceResponse(&workerService, &workerServices[i])
		}(i, *workerService)
	}

	wg.Wait()

	var orderDetailResponse dto.OrderDetailResponse
	formatOrderDetailResponse(&orderData, &orderDetailResponse)

	orderDetailResponse.WorkerService = workerServices

	return orderDetailResponse, nil
}

func formatOrderDetailResponse(order *domain.Order, orderDetailResponse *dto.OrderDetailResponse) {
	*orderDetailResponse = dto.OrderDetailResponse{
		OrderID:     order.OrderID,
		StatusOrder: order.OrderStatus.String(),
		WorkerID:    order.WorkerID,
		WorkerName:  order.Worker.Name,
		WorkerImage: order.Worker.Image,
		Tag: dto.TagResponse{
			ID:             order.Worker.Tag.ID,
			Tag:            order.Worker.Tag.Tag,
			Specialization: order.Worker.Tag.Specialization,
		},
		TransactionTime: order.Payment.CreatedAt.Format("02 January 2006 15:04"),
		PaymentMethod:   order.Payment.PaymentType.String(),
		PaymentStatus:   order.Payment.Status.String(),
		Location:        order.Address.Street,
		ServiceFee:      order.Payment.ServiceFee,
	}
}

func (s *OrderService) UpdateOrder(ctx context.Context, orderID string, req dto.UpdateOrderRequest) error {
	repositoryClient := s.repository.NewOrderRepositoryClient(true)

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

	orderData, err := repositoryClient.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	orderData, err = parseUpdateOrder(orderData, req)
	if err != nil {
		return err
	}

	if err := repositoryClient.UpdateOrder(ctx, &orderData); err != nil {
		return err
	}

	if err := repositoryClient.Commit(); err != nil {
		return err
	}

	return nil
}

func parseUpdateOrder(order domain.Order, req dto.UpdateOrderRequest) (domain.Order, error) {
	if req.Status != "" {
		status, err := domain.ParseStatusOrder(req.Status)
		if err != nil {
			return order, err
		}

		order.OrderStatus = status
	}

	return order, nil
}
