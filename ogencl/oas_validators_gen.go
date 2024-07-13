// Code generated by ogen, DO NOT EDIT.

package ogencl

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *Payment) Validate() error {
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

func (s *PaymentAmount) Validate() error {
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

func (s PaymentAmountCurrency) Validate() error {
	switch s {
	case "RUB":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s PaymentConfirmation) Validate() error {
	switch s.Type {
	case PaymentConfirmationEmbeddedPaymentConfirmation:
		if err := s.PaymentConfirmationEmbedded.Validate(); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
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