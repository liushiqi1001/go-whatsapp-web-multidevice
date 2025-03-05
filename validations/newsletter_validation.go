package validations

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	domainNewsletter "github.com/liushiqi1001/go-whatsapp-web-multidevice/domains/newsletter"
	pkgError "github.com/liushiqi1001/go-whatsapp-web-multidevice/pkg/error"
)

func ValidateUnfollowNewsletter(ctx context.Context, request domainNewsletter.UnfollowRequest) error {
	err := validation.ValidateStructWithContext(ctx, &request,
		validation.Field(&request.NewsletterID, validation.Required),
	)

	if err != nil {
		return pkgError.ValidationError(err.Error())
	}

	return nil
}
