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
