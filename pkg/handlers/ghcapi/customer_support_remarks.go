package ghcapi

import (
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/appcontext"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/handlers/ghcapi/internal/payloads"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"

	customersupportremarksop "github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/customer_support_remarks"
)

// ListCustomerSupportRemarksHandler is a struct that describes listing customer support remarks for a move
type ListCustomerSupportRemarksHandler struct {
	handlers.HandlerContext
	services.CustomerSupportRemarksFetcher
}

// Handle handles the handling for getting a list of customer support remarks for a move
func (h ListCustomerSupportRemarksHandler) Handle(params customersupportremarksop.GetCustomerSupportRemarksForMoveParams) middleware.Responder {
	return h.AuditableAppContextFromRequestWithErrors(params.HTTPRequest,
		func(appCtx appcontext.AppContext) (middleware.Responder, error) {

			customerSupportRemarks, err := h.ListCustomerSupportRemarks(appCtx, params.Locator)
			if err != nil {
				if err == models.ErrFetchNotFound {
					appCtx.Logger().Error("Error fetching customer support remarks: ", zap.Error(err))
					return customersupportremarksop.NewGetCustomerSupportRemarksForMoveNotFound(), err
				}
				appCtx.Logger().Error("Error fetching customer support remarks: ", zap.Error(err))
				return customersupportremarksop.NewGetCustomerSupportRemarksForMoveInternalServerError(), err
			}

			returnPayload := payloads.CustomerSupportRemarks(*customerSupportRemarks)
			return customersupportremarksop.NewGetCustomerSupportRemarksForMoveOK().WithPayload(returnPayload), nil
		})
}