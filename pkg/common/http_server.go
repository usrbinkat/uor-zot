package common

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"

	"zotregistry.io/zot/pkg/api/config"
	"zotregistry.io/zot/pkg/api/constants"
	apiErr "zotregistry.io/zot/pkg/api/errors"
)

func AllowedMethods(methods ...string) []string {
	return append(methods, http.MethodOptions)
}

func AddExtensionSecurityHeaders() mux.MiddlewareFunc { //nolint:varnamelen
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			resp.Header().Set("X-Content-Type-Options", "nosniff")

			next.ServeHTTP(resp, req)
		})
	}
}

func ACHeadersHandler(config *config.Config, allowedMethods ...string) mux.MiddlewareFunc {
	allowedMethodsValue := strings.Join(allowedMethods, ",")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			resp.Header().Set("Access-Control-Allow-Methods", allowedMethodsValue)
			resp.Header().Set("Access-Control-Allow-Headers", "Authorization,content-type,"+constants.SessionClientHeaderName)

			if config.IsBasicAuthnEnabled() {
				resp.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			if req.Method == http.MethodOptions {
				return
			}

			next.ServeHTTP(resp, req)
		})
	}
}

func AuthzFail(w http.ResponseWriter, r *http.Request, realm string, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)

	// don't send auth headers if request is coming from UI
	if r.Header.Get(constants.SessionClientHeaderName) != constants.SessionClientHeaderValue {
		if realm == "" {
			realm = "Authorization Required"
		}

		realm = "Basic realm=" + strconv.Quote(realm)

		w.Header().Set("WWW-Authenticate", realm)
	}

	w.Header().Set("Content-Type", "application/json")
	WriteJSON(w, http.StatusForbidden, apiErr.NewErrorList(apiErr.NewError(apiErr.DENIED)))
}

func WriteJSON(response http.ResponseWriter, status int, data interface{}) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	WriteData(response, status, constants.DefaultMediaType, body)
}

func WriteData(w http.ResponseWriter, status int, mediaType string, data []byte) {
	w.Header().Set("Content-Type", mediaType)
	w.WriteHeader(status)
	_, _ = w.Write(data)
}

func QueryHasParams(values url.Values, params []string) bool {
	for _, param := range params {
		if !values.Has(param) {
			return false
		}
	}

	return true
}
