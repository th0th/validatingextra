package validatingextra

import (
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
