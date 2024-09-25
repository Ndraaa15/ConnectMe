package midtrans

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/env"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type Midtrans struct {
	midtrans coreapi.Client
}

func NewMidtrans(conf env.PaymentGateway) *Midtrans {
	coreapiClient := coreapi.Client{}
	coreapiClient.New(conf.ApiKey, midtrans.Sandbox)

	return &Midtrans{
		midtrans: coreapiClient,
	}
}

func (m *Midtrans) CreateTransaction(ctx context.Context, payment domain.Payment) (dto.TransactionResponse, error) {
	chargeReq := &coreapi.ChargeReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  payment.OrderID,
			GrossAmt: int64(payment.TotalPrice),
		},
	}

	if payment.PaymentType == domain.PaymentTypeBCAVa {
		chargeReq.PaymentType = coreapi.PaymentTypeBankTransfer
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		}
	} else if payment.PaymentType == domain.PaymentTypeMandiriVa {
		chargeReq.PaymentType = coreapi.PaymentTypeBankTransfer
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankMandiri,
		}
	} else if payment.PaymentType == domain.PaymentTypeBNIVa {
		chargeReq.PaymentType = coreapi.PaymentTypeBankTransfer
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBni,
		}
	} else if payment.PaymentType == domain.PaymentTypeBRIVa {
		chargeReq.PaymentType = coreapi.PaymentTypeBankTransfer
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBri,
		}
	} else if payment.PaymentType == domain.PaymentTypePermataVa {
		chargeReq.PaymentType = coreapi.PaymentTypeBankTransfer
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankPermata,
		}
	} else if payment.PaymentType == domain.PaymentTypeGopay {
		chargeReq.PaymentType = coreapi.PaymentTypeGopay
	} else if payment.PaymentType == domain.PaymentTypeShopeePay {
		chargeReq.PaymentType = coreapi.PaymentTypeShopeepay
	}

	resp, err := m.midtrans.ChargeTransaction(chargeReq)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	var paymentResp dto.TransactionResponse
	paymentResp.TotalPrice = payment.TotalPrice
	paymentResp.LimitTransactionDate = resp.ExpiryTime
	paymentResp.PaymentMethod = payment.PaymentType.String()

	if resp.PaymentType == "gopay" || resp.PaymentType == "shopeepay" {
		var eWalletAction []dto.EWalletAction
		for _, action := range resp.Actions {
			eWalletAction = append(eWalletAction, dto.EWalletAction{
				Name: action.Name,
				URL:  action.URL,
			})
		}

		paymentResp.EWalletResponse = dto.EWalletResponse{
			Actions: eWalletAction,
		}
	} else {
		paymentResp.VirtualAccountResponse = dto.VirtualAccountResponse{
			BankName:             resp.VaNumbers[0].Bank,
			VirtualAccountNumber: resp.VaNumbers[0].VANumber,
		}

	}

	return paymentResp, nil
}
