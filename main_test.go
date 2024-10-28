package validatingextra_test

import (
	"reflect"
	"testing"

	"github.com/RussellLuo/validating/v3"

	"github.com/th0th/validatingextra"
)

func TestEmail(t *testing.T) {
	tests := []struct {
		errs   validating.Errors
		name   string
		schema validating.Schema
	}{
		{
			errs: nil,
			name: "valid e-mail",
			schema: validating.Schema{
				validating.F("value", "bojack@hollywoo.com"): validatingextra.Email(),
			},
		},
		{
			errs: validating.NewErrors("value", validating.ErrInvalid, "is not a valid e-mail address"),
			name: "invalid e-mail",
			schema: validating.Schema{
				validating.F("value", "bojack"): validatingextra.Email(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := validating.Validate(tt.schema)
			if !reflect.DeepEqual(errs, tt.errs) {
				t.Errorf("got = %v, want %v", errs, tt.errs)
			}
		})
	}
}

func TestEmailNonDisposable(t *testing.T) {
	tests := []struct {
		errs   validating.Errors
		name   string
		schema validating.Schema
	}{
		{
			errs: nil,
			name: "non-disposable e-mail",
			schema: validating.Schema{
				validating.F("value", "bojack@hollywoo.com"): validatingextra.EmailNonDisposable(),
			},
		},
		{
			errs: validating.NewErrors("value", validating.ErrInvalid, "is disposable e-mail address"),
			name: "disposable e-mail",
			schema: validating.Schema{
				validating.F("value", "dummy@getnada.com"): validatingextra.EmailNonDisposable(),
			},
		},
		{
			errs: validating.NewErrors("value", validating.ErrInvalid, "is not a valid e-mail address"),
			name: "invalid e-mail",
			schema: validating.Schema{
				validating.F("value", "bojack"): validatingextra.Email(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := validating.Validate(tt.schema)
			if !reflect.DeepEqual(errs, tt.errs) {
				t.Errorf("got = %v, want %v", errs, tt.errs)
			}
		})
	}
}

func TestIpAddress(t *testing.T) {
	tests := []struct {
		errs   validating.Errors
		name   string
		schema validating.Schema
	}{
		{
			errs: nil,
			name: "valid IPv4 address",
			schema: validating.Schema{
				validating.F("value", "95.54.80.130"): validatingextra.IpAddress(),
			},
		},
		{
			errs: nil,
			name: "valid IPv6 address",
			schema: validating.Schema{
				validating.F("value", "2001:0000:130F:0000:0000:09C0:876A:130B"): validatingextra.IpAddress(),
			},
		},
		{
			errs: validating.NewErrors("value", validating.ErrInvalid, "is not a valid IP address"),
			name: "some string",
			schema: validating.Schema{
				validating.F("value", "some string"): validatingextra.IpAddress(),
			},
		},
		{
			errs: validating.NewErrors("value", validating.ErrInvalid, "is not a valid IP address"),
			name: "a domain",
			schema: validating.Schema{
				validating.F("value", "Horsin' Around"): validatingextra.IpAddress(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := validating.Validate(tt.schema)
			if !reflect.DeepEqual(errs, tt.errs) {
				t.Errorf("got = %v, want %v", errs, tt.errs)
			}
		})
	}
}

func TestPointerValue(t *testing.T) {
	var nilStringPtr *string

	tests := []struct {
		errs   validating.Errors
		name   string
		schema validating.Schema
	}{
		{
			errs: nil,
			name: "valid e-mail address pointer",
			schema: validating.Schema{
				validating.F("value", pointer("bojack@hollywoo.com")): validatingextra.PointerValue[string](validatingextra.Email()),
			},
		},
		{
			errs: validating.NewErrors("value", validating.ErrUnsupported, "expected a *int but got *string"),
			name: "wrong type of pointer",
			schema: validating.Schema{
				validating.F("value", pointer("bojack@hollywoo.com")): validatingextra.PointerValue[int](validatingextra.Email()),
			},
		},
		{
			errs: validating.NewErrors("value", validating.ErrUnsupported, "expected a *string but got <nil>"),
			name: "nil",
			schema: validating.Schema{
				validating.F("value", nil): validatingextra.PointerValue[string](validatingextra.Email()),
			},
		},
		{
			errs: validating.NewErrors("value", validating.ErrUnsupported, "expected a non-nil pointer but got nil pointer"),
			name: "nil pointer",
			schema: validating.Schema{
				validating.F("value", nilStringPtr): validatingextra.PointerValue[string](validatingextra.Email()),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := validating.Validate(tt.schema)
			if !reflect.DeepEqual(errs, tt.errs) {
				t.Errorf("got = %v, want %v", errs, tt.errs)
			}
		})
	}
}

func pointer[T any](v T) *T {
	return &v
}
