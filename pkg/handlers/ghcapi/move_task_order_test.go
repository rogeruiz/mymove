//RA Summary: gosec - errcheck - Unchecked return value
//RA: Linter flags errcheck error: Ignoring a method's return value can cause the program to overlook unexpected states and conditions.
//RA: Functions with unchecked return values in the file are used fetch data and assign data to a variable that is checked later on
//RA: Given the return value is being checked in a different line and the functions that are flagged by the linter are being used to assign variables
//RA: in a unit test, then there is no risk
//RA Developer Status: Mitigated
//RA Validator Status: Mitigated
//RA Modified Severity: N/A
// nolint:errcheck
package ghcapi

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"time"

	"github.com/go-openapi/runtime/middleware"

	"github.com/transcom/mymove/pkg/apperror"
	moverouter "github.com/transcom/mymove/pkg/services/move"
	"github.com/transcom/mymove/pkg/trace"

	"github.com/stretchr/testify/mock"

	"github.com/transcom/mymove/pkg/services/mocks"

	"github.com/transcom/mymove/pkg/gen/ghcmessages"

	mtoserviceitem "github.com/transcom/mymove/pkg/services/mto_service_item"

	"github.com/transcom/mymove/pkg/etag"
	"github.com/transcom/mymove/pkg/services/query"

	"github.com/go-openapi/strfmt"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/move_task_order"
	movetaskorderops "github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/move_task_order"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
	movetaskorder "github.com/transcom/mymove/pkg/services/move_task_order"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *HandlerSuite) TestGetMoveTaskOrderHandlerIntegration() {
	order := testdatagen.MakeDefaultOrder(suite.DB())
	moveTaskOrder := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Order: order,
	})
	testdatagen.MakeReService(suite.DB(), testdatagen.Assertions{
		ReService: models.ReService{
			ID:   uuid.FromStringOrNil("1130e612-94eb-49a7-973d-72f33685e551"),
			Code: models.ReServiceCodeMS,
		},
	})

	testdatagen.MakeReService(suite.DB(), testdatagen.Assertions{
		ReService: models.ReService{
			ID:   uuid.FromStringOrNil("9dc919da-9b66-407b-9f17-05c0f03fcb50"),
			Code: models.ReServiceCodeCS,
		},
	})
	request := httptest.NewRequest("GET", "/move-task-orders/{moveTaskOrderID}", nil)
	params := move_task_order.GetMoveTaskOrderParams{
		HTTPRequest:     request,
		MoveTaskOrderID: moveTaskOrder.ID.String(),
	}
	handlerConfig := handlers.NewHandlerConfig(suite.DB(), suite.Logger())
	handler := GetMoveTaskOrderHandler{
		handlerConfig,
		movetaskorder.NewMoveTaskOrderFetcher(),
	}

	response := handler.Handle(params)
	suite.IsNotErrResponse(response)
	moveTaskOrderResponse := response.(*movetaskorderops.GetMoveTaskOrderOK)
	moveTaskOrderPayload := moveTaskOrderResponse.Payload

	suite.Assertions.IsType(&move_task_order.GetMoveTaskOrderOK{}, response)
	suite.Equal(strfmt.UUID(moveTaskOrder.ID.String()), moveTaskOrderPayload.ID)
	suite.Nil(moveTaskOrderPayload.AvailableToPrimeAt)
	// TODO: Check that the *moveTaskOrderPayload.Status is not "canceled"
	// suite.False(*moveTaskOrderPayload.IsCanceled)
	suite.Equal(strfmt.UUID(moveTaskOrder.OrdersID.String()), moveTaskOrderPayload.OrderID)
	suite.NotNil(moveTaskOrderPayload.ReferenceID)
}

func (suite *HandlerSuite) TestUpdateMoveTaskOrderHandlerIntegrationSuccess() {
	testdatagen.MakeReService(suite.DB(), testdatagen.Assertions{
		ReService: models.ReService{
			Code: models.ReServiceCodeMS,
		},
	})

	testdatagen.MakeReService(suite.DB(), testdatagen.Assertions{
		ReService: models.ReService{
			Code: models.ReServiceCodeCS,
		},
	})

	validStatuses := []struct {
		desc   string
		status models.MoveStatus
	}{
		{"Submitted", models.MoveStatusSUBMITTED},
		{"Service Counseling Completed", models.MoveStatusServiceCounselingCompleted},
	}
	for _, validStatus := range validStatuses {
		move := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{Move: models.Move{Status: validStatus.status}})

		request := httptest.NewRequest("PATCH", "/move-task-orders/{moveID}/status", nil)
		requestUser := testdatagen.MakeStubbedUser(suite.DB())
		request = suite.AuthenticateUserRequest(request, requestUser)

		traceID, err := uuid.NewV4()
		suite.FatalNoError(err, "Error creating a new trace ID.")
		request = request.WithContext(trace.NewContext(request.Context(), traceID))

		serviceItemCodes := ghcmessages.MTOApprovalServiceItemCodes{
			ServiceCodeMS: true,
			ServiceCodeCS: true,
		}
		params := move_task_order.UpdateMoveTaskOrderStatusParams{
			HTTPRequest:      request,
			MoveTaskOrderID:  move.ID.String(),
			IfMatch:          etag.GenerateEtag(move.UpdatedAt),
			ServiceItemCodes: &serviceItemCodes,
		}
		handlerConfig := handlers.NewHandlerConfig(suite.DB(), suite.Logger())
		queryBuilder := query.NewQueryBuilder()
		moveRouter := moverouter.NewMoveRouter()
		siCreator := mtoserviceitem.NewMTOServiceItemCreator(queryBuilder, moveRouter)

		// setup the handler
		handler := UpdateMoveTaskOrderStatusHandlerFunc{handlerConfig,
			movetaskorder.NewMoveTaskOrderUpdater(queryBuilder, siCreator, moveRouter),
		}

		// make the request
		response := handler.Handle(params)

		suite.IsNotErrResponse(response)
		moveResponse := response.(*movetaskorderops.UpdateMoveTaskOrderStatusOK)
		movePayload := moveResponse.Payload

		updatedMove := models.Move{}
		suite.DB().Find(&updatedMove, movePayload.ID)
		suite.Equal(models.MoveStatusAPPROVED, updatedMove.Status)

		suite.Assertions.IsType(&move_task_order.UpdateMoveTaskOrderStatusOK{}, response)
		suite.Equal(strfmt.UUID(move.ID.String()), movePayload.ID)
		suite.NotNil(movePayload.AvailableToPrimeAt)
		suite.HasWebhookNotification(move.ID, traceID) // this action always creates a notification for the Prime

		// also check MTO level service items are properly created
		var serviceItems models.MTOServiceItems
		suite.DB().Eager("ReService").Where("move_id = ?", move.ID).All(&serviceItems)
		suite.Len(serviceItems, 2, "Expected to find at most 2 service items")

		containsServiceCode := func(items models.MTOServiceItems, target models.ReServiceCode) bool {
			for _, si := range items {
				if si.ReService.Code == target {
					return true
				}
			}

			return false
		}

		suite.True(containsServiceCode(serviceItems, models.ReServiceCodeMS), fmt.Sprintf("Expected to find reServiceCode, %s, in array.", models.ReServiceCodeMS))
		suite.True(containsServiceCode(serviceItems, models.ReServiceCodeCS), fmt.Sprintf("Expected to find reServiceCode, %s, in array.", models.ReServiceCodeCS))
	}
}

func (suite *HandlerSuite) TestUpdateMoveTaskOrderHandlerIntegrationWithStaleEtag() {
	move := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Stub: true,
		Move: models.Move{
			Status: models.MoveStatusSUBMITTED,
		},
	})

	request := httptest.NewRequest("PATCH", "/move-task-orders/{moveTaskOrderID}/status", nil)
	requestUser := testdatagen.MakeStubbedUser(suite.DB())
	request = suite.AuthenticateUserRequest(request, requestUser)
	params := move_task_order.UpdateMoveTaskOrderStatusParams{
		HTTPRequest:     request,
		MoveTaskOrderID: move.ID.String(),
		IfMatch:         etag.GenerateEtag(time.Now()),
	}
	handlerConfig := handlers.NewHandlerConfig(suite.DB(), suite.Logger())

	// Stale ETags are already unit tested in the move_task_order_updater_test,
	// so we can mock this here to speed up the test and avoid hitting the DB
	moveUpdater := &mocks.MoveTaskOrderUpdater{}
	moveUpdater.On("MakeAvailableToPrime",
		mock.AnythingOfType("*appcontext.appContext"),
		mock.Anything,
		mock.Anything,
		false,
		false,
	).Return(nil, apperror.PreconditionFailedError{})

	// make the request
	handler := UpdateMoveTaskOrderStatusHandlerFunc{handlerConfig, moveUpdater}
	response := handler.Handle(params)
	suite.Assertions.IsType(&move_task_order.UpdateMoveTaskOrderStatusPreconditionFailed{}, response)
}

func (suite *HandlerSuite) TestUpdateMoveTaskOrderHandlerIntegrationWithIncompleteOrder() {
	orderWithoutDefaults := testdatagen.MakeOrderWithoutDefaults(suite.DB(), testdatagen.Assertions{})
	move := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			Status: models.MoveStatusServiceCounselingCompleted,
		},
		Order: orderWithoutDefaults,
	})

	request := httptest.NewRequest("PATCH", "/move-task-orders/{moveTaskOrderID}/status", nil)
	requestUser := testdatagen.MakeStubbedUser(suite.DB())
	request = suite.AuthenticateUserRequest(request, requestUser)
	params := move_task_order.UpdateMoveTaskOrderStatusParams{
		HTTPRequest:     request,
		MoveTaskOrderID: move.ID.String(),
		IfMatch:         etag.GenerateEtag(move.UpdatedAt),
	}
	handlerConfig := handlers.NewHandlerConfig(suite.DB(), suite.Logger())
	queryBuilder := query.NewQueryBuilder()
	moveRouter := moverouter.NewMoveRouter()
	siCreator := mtoserviceitem.NewMTOServiceItemCreator(queryBuilder, moveRouter)

	// make the request
	handler := UpdateMoveTaskOrderStatusHandlerFunc{handlerConfig,
		movetaskorder.NewMoveTaskOrderUpdater(queryBuilder, siCreator, moveRouter),
	}
	response := handler.Handle(params)

	suite.Assertions.IsType(&move_task_order.UpdateMoveTaskOrderStatusUnprocessableEntity{}, response)
	invalidResponse := response.(*move_task_order.UpdateMoveTaskOrderStatusUnprocessableEntity).Payload
	errorDetail := invalidResponse.Detail

	suite.Contains(*errorDetail, "TransportationAccountingCode cannot be blank.")
	suite.Contains(*errorDetail, "OrdersNumber cannot be blank.")
	suite.Contains(*errorDetail, "DepartmentIndicator cannot be blank.")
	suite.Contains(*errorDetail, "OrdersTypeDetail cannot be blank.")
}

func (suite *HandlerSuite) TestUpdateMTOStatusServiceCounselingCompletedHandler() {
	setupTestData := func() UpdateMTOStatusServiceCounselingCompletedHandlerFunc {
		handlerConfig := handlers.NewHandlerConfig(suite.DB(), suite.Logger())
		queryBuilder := query.NewQueryBuilder()
		moveRouter := moverouter.NewMoveRouter()
		siCreator := mtoserviceitem.NewMTOServiceItemCreator(queryBuilder, moveRouter)
		handler := UpdateMTOStatusServiceCounselingCompletedHandlerFunc{
			handlerConfig,
			movetaskorder.NewMoveTaskOrderUpdater(queryBuilder, siCreator, moveRouter),
		}
		return handler
	}

	suite.Run("Successful move status update to Service Counseling Completed - Integration", func() {
		handler := setupTestData()
		request := httptest.NewRequest("PATCH", "/move-task-orders/{moveTaskOrderID}/status/service-counseling-completed", nil)
		requestUser := testdatagen.MakeServicesCounselorOfficeUser(suite.DB(), testdatagen.Assertions{})
		request = suite.AuthenticateOfficeRequest(request, requestUser)
		move := testdatagen.MakeHHGMoveWithShipment(suite.DB(), testdatagen.Assertions{
			Move: models.Move{
				Status: models.MoveStatusNeedsServiceCounseling,
			},
		})

		params := move_task_order.UpdateMTOStatusServiceCounselingCompletedParams{
			HTTPRequest:     request,
			MoveTaskOrderID: move.ID.String(),
			IfMatch:         etag.GenerateEtag(move.UpdatedAt),
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)
		moveTaskOrderResponse := response.(*movetaskorderops.UpdateMTOStatusServiceCounselingCompletedOK)
		moveTaskOrderPayload := moveTaskOrderResponse.Payload
		suite.NoError(moveTaskOrderPayload.Validate(strfmt.Default))

		suite.IsType(&move_task_order.UpdateMTOStatusServiceCounselingCompletedOK{}, response)
		suite.Equal(strfmt.UUID(move.ID.String()), moveTaskOrderPayload.ID)
		suite.NotNil(moveTaskOrderPayload.ServiceCounselingCompletedAt)
		suite.EqualValues(models.MoveStatusServiceCounselingCompleted, moveTaskOrderPayload.Status)
	})

	suite.Run("Unsuccessful move status update to Service Counseling Completed, forbidden - Integration", func() {
		handler := setupTestData()
		request := httptest.NewRequest("PATCH", "/move-task-orders/{moveTaskOrderID}/status/service-counseling-completed", nil)
		forbiddenUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{})
		forbiddenRequest := suite.AuthenticateOfficeRequest(request, forbiddenUser)

		params := move_task_order.UpdateMTOStatusServiceCounselingCompletedParams{
			HTTPRequest: forbiddenRequest,
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)
		suite.IsType(&move_task_order.UpdateMTOStatusServiceCounselingCompletedForbidden{}, response)
	})

	suite.Run("Unsuccessful move status update to Service Counseling Completed, not found - Integration", func() {
		handler := setupTestData()
		request := httptest.NewRequest("PATCH", "/move-task-orders/{moveTaskOrderID}/status/service-counseling-completed", nil)
		requestUser := testdatagen.MakeServicesCounselorOfficeUser(suite.DB(), testdatagen.Assertions{})
		request = suite.AuthenticateOfficeRequest(request, requestUser)
		params := move_task_order.UpdateMTOStatusServiceCounselingCompletedParams{
			HTTPRequest:     request,
			MoveTaskOrderID: uuid.Must(uuid.NewV4()).String(),
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)
		suite.IsType(&move_task_order.UpdateMTOStatusServiceCounselingCompletedNotFound{}, response)
	})

	suite.Run("Unsuccessful move status update to Service Counseling Completed, eTag does not match - Integration", func() {
		handler := setupTestData()
		request := httptest.NewRequest("PATCH", "/move-task-orders/{moveTaskOrderID}/status/service-counseling-completed", nil)
		requestUser := testdatagen.MakeServicesCounselorOfficeUser(suite.DB(), testdatagen.Assertions{})
		request = suite.AuthenticateOfficeRequest(request, requestUser)
		move := testdatagen.MakeNeedsServiceCounselingMove(suite.DB())
		testdatagen.MakeMTOShipmentMinimal(suite.DB(), testdatagen.Assertions{
			Move: move,
		})

		params := move_task_order.UpdateMTOStatusServiceCounselingCompletedParams{
			HTTPRequest:     request,
			MoveTaskOrderID: move.ID.String(),
			IfMatch:         etag.GenerateEtag(time.Now()),
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)
		suite.IsType(&move_task_order.UpdateMTOStatusServiceCounselingCompletedPreconditionFailed{}, response)
	})

	suite.Run("Unsuccessful move status update to Service Counseling Completed, state conflict - Integration", func() {
		handler := setupTestData()
		request := httptest.NewRequest("PATCH", "/move-task-orders/{moveTaskOrderID}/status/service-counseling-completed", nil)
		requestUser := testdatagen.MakeServicesCounselorOfficeUser(suite.DB(), testdatagen.Assertions{})
		request = suite.AuthenticateOfficeRequest(request, requestUser)
		draftMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
			Move: models.Move{
				Status: models.MoveStatusDRAFT,
			},
		})
		testdatagen.MakeMTOShipmentMinimal(suite.DB(), testdatagen.Assertions{
			Move: draftMove,
		})

		params := move_task_order.UpdateMTOStatusServiceCounselingCompletedParams{
			HTTPRequest:     request,
			MoveTaskOrderID: draftMove.ID.String(),
			IfMatch:         etag.GenerateEtag(draftMove.UpdatedAt),
		}

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)
		suite.IsType(&move_task_order.UpdateMTOStatusServiceCounselingCompletedConflict{}, response)
	})

	suite.Run("Unsuccessful move status update to Service Counseling Completed, misc mocked errors - Integration", func() {
		handlerOrig := setupTestData()
		request := httptest.NewRequest("PATCH", "/move-task-orders/{moveTaskOrderID}/status/service-counseling-completed", nil)
		requestUser := testdatagen.MakeServicesCounselorOfficeUser(suite.DB(), testdatagen.Assertions{})
		request = suite.AuthenticateOfficeRequest(request, requestUser)
		testCases := []struct {
			mockError       error
			handlerResponse middleware.Responder
		}{
			{apperror.InvalidInputError{}, &move_task_order.UpdateMTOStatusServiceCounselingCompletedUnprocessableEntity{}},
			{apperror.QueryError{}, &move_task_order.UpdateMTOStatusServiceCounselingCompletedInternalServerError{}},
			{errors.New("generic error"), &move_task_order.UpdateMTOStatusServiceCounselingCompletedInternalServerError{}},
		}

		move := testdatagen.MakeStubbedMoveWithStatus(suite.DB(), models.MoveStatusNeedsServiceCounseling)
		eTag := etag.GenerateEtag(move.UpdatedAt)
		params := move_task_order.UpdateMTOStatusServiceCounselingCompletedParams{
			HTTPRequest:     request,
			MoveTaskOrderID: move.ID.String(),
			IfMatch:         eTag,
		}

		for _, testCase := range testCases {
			mockUpdater := mocks.MoveTaskOrderUpdater{}
			mockUpdater.On("UpdateStatusServiceCounselingCompleted",
				mock.AnythingOfType("*appcontext.appContext"),
				move.ID,
				eTag,
			).Return(&models.Move{}, testCase.mockError)

			handler := UpdateMTOStatusServiceCounselingCompletedHandlerFunc{
				handlerOrig.HandlerConfig,
				&mockUpdater,
			}

			response := handler.Handle(params)
			suite.IsNotErrResponse(response)
			suite.IsType(testCase.handlerResponse, response)
		}
	})
}

func (suite *HandlerSuite) TestUpdateMoveTIORemarksHandler() {
	setupTestData := func() (models.Move, UpdateMoveTIORemarksHandlerFunc, models.User) {

		order := testdatagen.MakeDefaultOrder(suite.DB())
		moveTaskOrder := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
			Move: models.Move{
				Status: models.MoveStatusNeedsServiceCounseling,
			},
			Order: order,
		})
		requestUser := testdatagen.MakeStubbedUser(suite.DB())
		handlerConfig := handlers.NewHandlerConfig(suite.DB(), suite.Logger())
		queryBuilder := query.NewQueryBuilder()
		moveRouter := moverouter.NewMoveRouter()
		siCreator := mtoserviceitem.NewMTOServiceItemCreator(queryBuilder, moveRouter)
		handler := UpdateMoveTIORemarksHandlerFunc{
			handlerConfig,
			movetaskorder.NewMoveTaskOrderUpdater(queryBuilder, siCreator, moveRouter),
		}
		return moveTaskOrder, handler, requestUser
	}

	remarks := "Reweigh requested"
	suite.Run("Successfully update the Move's TIORemarks field", func() {
		move, handler, requestUser := setupTestData()
		request := httptest.NewRequest("PATCH", "/move-task-orders/{moveTaskOrderID}/tio-remarks", nil)
		request = suite.AuthenticateUserRequest(request, requestUser)

		params := move_task_order.UpdateMoveTIORemarksParams{
			HTTPRequest:     request,
			MoveTaskOrderID: move.ID.String(),
			Body:            &ghcmessages.Move{TioRemarks: &remarks},
			IfMatch:         etag.GenerateEtag(move.UpdatedAt),
		}
		response := handler.Handle(params)
		suite.IsNotErrResponse(response)
		moveTaskOrderResponse := response.(*movetaskorderops.UpdateMoveTIORemarksOK)
		moveTaskOrderPayload := moveTaskOrderResponse.Payload

		suite.Assertions.IsType(&move_task_order.UpdateMoveTIORemarksOK{}, response)
		updatedMove := models.Move{}
		suite.DB().Find(&updatedMove, moveTaskOrderPayload.ID)
		suite.Equal(moveTaskOrderPayload.TioRemarks, updatedMove.TIORemarks)
	})

	suite.Run("Unsuccessful move TIO Remarks, eTag does not match", func() {
		move, handler, requestUser := setupTestData()
		request := httptest.NewRequest("PATCH", "/move-task-orders/{moveTaskOrderID}/tio-remarks", nil)
		request = suite.AuthenticateUserRequest(request, requestUser)

		params := move_task_order.UpdateMoveTIORemarksParams{
			HTTPRequest:     request,
			MoveTaskOrderID: move.ID.String(),
			Body:            &ghcmessages.Move{TioRemarks: &remarks},
		}
		response := handler.Handle(params)
		suite.IsNotErrResponse(response)

		suite.Assertions.IsType(&move_task_order.UpdateMoveTIORemarksPreconditionFailed{}, response)
	})

	suite.Run("Unsuccessful move TIO Remarks update, not found", func() {
		_, handler, requestUser := setupTestData()
		request := httptest.NewRequest("PATCH", "/move-task-orders/{moveTaskOrderID}/tio-remarks", nil)
		request = suite.AuthenticateUserRequest(request, requestUser)

		params := move_task_order.UpdateMoveTIORemarksParams{
			HTTPRequest:     request,
			MoveTaskOrderID: uuid.Must(uuid.NewV4()).String(),
			Body:            &ghcmessages.Move{TioRemarks: &remarks},
		}
		response := handler.Handle(params)
		suite.IsNotErrResponse(response)

		suite.Assertions.IsType(&move_task_order.UpdateMoveTIORemarksNotFound{}, response)
	})
}
