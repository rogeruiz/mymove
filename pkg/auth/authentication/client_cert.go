package authentication

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/transcom/mymove/pkg/appcontext"
	"github.com/transcom/mymove/pkg/auth"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
	"github.com/transcom/mymove/pkg/services/query"
	"github.com/transcom/mymove/pkg/services/user"
)

type authClientCertKey string

const clientCertContextKey authClientCertKey = "clientCert"

// SetClientCertInRequestContext returns a copy of the request's Context() with the client certificate data
func SetClientCertInRequestContext(r *http.Request, clientCert *models.ClientCert) context.Context {
	return context.WithValue(r.Context(), clientCertContextKey, clientCert)
}

// ClientCertFromRequestContext gets the reference to the ClientCert stored in the request.Context()
func ClientCertFromRequestContext(r *http.Request) *models.ClientCert {
	return ClientCertFromContext(r.Context())
}

// ClientCertFromContext gets the reference to the ClientCert stored in the request.Context()
func ClientCertFromContext(ctx context.Context) *models.ClientCert {
	if clientCert, ok := ctx.Value(clientCertContextKey).(*models.ClientCert); ok {
		return clientCert
	}
	return nil
}

// SetMockSessionInContext returns a copy of the request's Context() with a mock session containing the user id of the user
// linked to the client certificate data. This user id is used for auditing purposes.
func SetMockSessionInContext(ctx context.Context, user models.User) context.Context {
	mockSession := auth.Session{
		UserID: user.ID,
	}
	return auth.SetSessionInContext(ctx, &mockSession)
}

// ClientCertMiddleware enforces that the incoming request includes a known client certificate, and stores the fetched permissions in the session
func ClientCertMiddleware(appCtx appcontext.AppContext) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		mw := func(w http.ResponseWriter, r *http.Request) {
			newAppCtx := appcontext.NewAppContextFromContext(r.Context(), appCtx)

			if r.TLS == nil || len(r.TLS.PeerCertificates) == 0 {
				newAppCtx.Logger().Info("Unauthenticated")
				http.Error(w, http.StatusText(401), http.StatusUnauthorized)
				return
			}

			// get DER hash
			hash := sha256.Sum256(r.TLS.PeerCertificates[0].Raw)
			hashString := hex.EncodeToString(hash[:])

			clientCert, err := models.FetchClientCert(newAppCtx.DB(), hashString)
			if err != nil {
				// This is not a known client certificate at all
				newAppCtx.Logger().Info("Unknown / unregistered client certificate")
				http.Error(w, http.StatusText(401), http.StatusUnauthorized)
				return
			}
			ctx := SetClientCertInRequestContext(r, clientCert)
			user, err := user.NewUserFetcher(query.NewQueryBuilder()).FetchUser(newAppCtx, []services.QueryFilter{query.NewQueryFilter("id", "=", clientCert.UserID)})
			if err != nil {
				newAppCtx.Logger().Info("Client certificate not linked to user")
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}
			ctx = SetMockSessionInContext(ctx, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(mw)
	}
}

// DevlocalClientCertMiddleware fakes the client cert as always
// devlocal. This will only be used if devlocal auth is enabled
func DevlocalClientCertMiddleware(appCtx appcontext.AppContext) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		mw := func(w http.ResponseWriter, r *http.Request) {
			newAppCtx := appcontext.NewAppContextFromContext(r.Context(), appCtx)

			hashString := ""
			// if a TLS connection has a client cert, use that
			if r.TLS != nil && len(r.TLS.PeerCertificates) != 0 {
				// get DER hash
				hash := sha256.Sum256(r.TLS.PeerCertificates[0].Raw)
				hashString = hex.EncodeToString(hash[:])
				newAppCtx.Logger().Info("TLS connection has a client certificate")
			} else {
				// otherwise, for devlocal, default to the devlocal cert
				// This hash gets populated as part of migration
				// 20191212230438_add_devlocal-mtls_client_cert.up.sql
				hashString = "2c0c1fc67a294443292a9e71de0c71cc374fe310e8073f8cdc15510f6b0ef4db"
				newAppCtx.Logger().Info("TLS connection doesn't have a client certificate")
			}

			clientCert, err := models.FetchClientCert(newAppCtx.DB(), hashString)
			if err != nil {
				// This is not a known client certificate at all
				newAppCtx.Logger().Info("Unknown / unregistered client certificate")
				http.Error(w, http.StatusText(401), http.StatusUnauthorized)
				return
			}
			ctx := SetClientCertInRequestContext(r, clientCert)
			user, err := user.NewUserFetcher(query.NewQueryBuilder()).FetchUser(newAppCtx, []services.QueryFilter{query.NewQueryFilter("id", "=", clientCert.UserID)})
			if err != nil {
				newAppCtx.Logger().Info("Client certificate not linked to user")
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}
			ctx = SetMockSessionInContext(ctx, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(mw)
	}
}
