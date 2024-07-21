// Code generated by ogen, DO NOT EDIT.

package ogencl

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *Payment) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Confirmation.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "confirmation",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *PaymentConfirmation) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *PaymentConfirmationEmbedded) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s PaymentConfirmationEmbeddedType) Validate() error {
	switch s {
	case "embedded":
		return nil
	case "external":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s PaymentConfirmationType) Validate() error {
	switch s {
	case "embedded":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s PaymentStatus) Validate() error {
	switch s {
	case "pending":
		return nil
	case "waiting_for_capture":
		return nil
	case "succeeded":
		return nil
	case "canceled":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *ReqPayment) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Amount.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "amount",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Confirmation.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "confirmation",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ReqPaymentAmount) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Currency.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "currency",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s ReqPaymentAmountCurrency) Validate() error {
	switch s {
	case "RUB":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s ReqPaymentConfirmation) Validate() error {
	switch s.Type {
	case PaymentConfirmationEmbeddedReqPaymentConfirmation:
		if err := s.PaymentConfirmationEmbedded.Validate(); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s *V3PaymentsGetOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Items {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "items",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *YookassaHookPostReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Event.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "event",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Object.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "object",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s YookassaHookPostReqEvent) Validate() error {
	switch s {
	case "payment.waiting_for_capture":
		return nil
	case "payment.succeeded":
		return nil
	case "payment.canceled":
		return nil
	case "refund.succeeded":
		return nil
	case "payout.succeeded":
		return nil
	case "payout.canceled":
		return nil
	case "deal.closed":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s YookassaHookPostReqObject) Validate() error {
	switch s.Type {
	case ReqPaymentYookassaHookPostReqObject:
		if err := s.ReqPayment.Validate(); err != nil {
			return err
		}
		return nil
	case PaymentYookassaHookPostReqObject:
		if err := s.Payment.Validate(); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s YookassaHookPostReqType) Validate() error {
	switch s {
	case "notification":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
