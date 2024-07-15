// Code generated by ogen, DO NOT EDIT.

package ogencl

// V3PaymentsGetParams is parameters of GET /v3/payments operation.
type V3PaymentsGetParams struct {
	Cursor OptString
	Limit  OptInt
}

// V3PaymentsPaymentIDGetParams is parameters of GET /v3/payments/{payment_id} operation.
type V3PaymentsPaymentIDGetParams struct {
	// Ключ идемпотентности.
	IdempotenceKey string
	PaymentID      string
}

// V3PaymentsPostParams is parameters of POST /v3/payments operation.
type V3PaymentsPostParams struct {
	// Ключ идемпотентности.
	IdempotenceKey string
}
