package examples

import (
	v "github.com/RussellLuo/validating/v3"

	"github.com/th0th/validatingextra"
)

type Data struct {
	Email              string
	EmailNonDisposable string
	IpAddress          string
	PointerValue       *string
}

func (d Data) Schema() v.Schema {
	return v.Schema{
		// Email is a validator that checks if the field is a valid e-mail address.
		v.F("email", d.Email): validatingextra.Email(),

		// EmailNonDisposable is a validator that checks if the field is a valid e-mail address and is not a disposable e-mail address.
		v.F("emailNonDisposable", d.EmailNonDisposable): validatingextra.EmailNonDisposable(),

		// IpAddress is a validator that checks if the field is a valid IP address.
		v.F("ipAddress", d.IpAddress): validatingextra.IpAddress(),

		// PointerValue is a composite validator that checks if the field is a non-nil pointer and then validates the value it points to.
		v.F("pointerValue", d.PointerValue): validatingextra.PointerValue[string](validatingextra.Email()),
	}
}
