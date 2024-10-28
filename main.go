package validatingextra

import (
	"fmt"
	"net"
	"net/mail"

	"github.com/RussellLuo/validating/v3"
	"github.com/th0th/is-email-disposable/pkg/isemaildisposable"
	"github.com/th0th/is-email-disposable/pkg/service/domain"
)

var disposableDomainService isemaildisposable.DomainService

func init() {
	disposableDomainService2, err := domain.New()
	if err != nil {
		panic(err)
	}

	disposableDomainService = disposableDomainService2
}

// Email is a validator that checks if the field is a valid e-mail address.
func Email() *validating.MessageValidator {
	messageValidator := validating.MessageValidator{
		Message: "is not a valid e-mail address",
	}

	messageValidator.Validator = validating.Func(func(field *validating.Field) validating.Errors {
		v, ok := field.Value.(string)
		if !ok {
			return validating.NewUnsupportedErrors("Email", field, "")
		}

		address, err := mail.ParseAddress(v)
		if err != nil || address.Address != v {
			return validating.NewInvalidErrors(field, messageValidator.Message)
		}

		return nil
	})

	return &messageValidator
}

// EmailNonDisposable is a validator that checks if the field is a valid e-mail address and is not a disposable e-mail address.
func EmailNonDisposable() *validating.MessageValidator {
	messageValidator := validating.MessageValidator{
		Message: "is disposable e-mail address",
	}

	messageValidator.Validator = validating.Func(func(field *validating.Field) validating.Errors {
		v, ok := field.Value.(string)
		if !ok {
			return validating.NewUnsupportedErrors("EmailNonDisposable", field, "")
		}

		checkResult := disposableDomainService.Check(v)
		if checkResult.IsDisposable {
			return validating.NewInvalidErrors(field, messageValidator.Message)
		}

		return nil
	})

	return &messageValidator
}

// IpAddress is a validator that checks if the field is a valid IP address.
func IpAddress() *validating.MessageValidator {
	messageValidator := validating.MessageValidator{
		Message: "is not a valid IP address",
	}

	messageValidator.Validator = validating.Func(func(field *validating.Field) validating.Errors {
		v, ok := field.Value.(string)
		if !ok {
			return validating.NewUnsupportedErrors("IpAddress", field, "")
		}

		if net.ParseIP(v) == nil {
			return validating.NewInvalidErrors(field, messageValidator.Message)
		}

		return nil
	})

	return &messageValidator
}

// PointerValue is a composite validator that checks if the field is a non-nil pointer and then validates the value it points to.
func PointerValue[T any](validator validating.Validator) validating.Validator {
	return validating.Func(func(field *validating.Field) validating.Errors {
		ptr, isOk := field.Value.(*T)
		if !isOk {
			return validating.NewErrors(field.Name, validating.ErrUnsupported, fmt.Sprintf("expected a %T but got %T", new(T), field.Value))
		}

		if ptr == nil {
			return validating.NewErrors(field.Name, validating.ErrUnsupported, "expected a non-nil pointer but got nil pointer")
		}

		field2 := validating.Field{
			Name:  field.Name,
			Value: *ptr,
		}

		return validator.Validate(&field2)
	})
}
