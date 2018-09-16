package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"gitlab.inn4science.com/gophers/service-kit/api/render"
	"gitlab.inn4science.com/gophers/service-kit/log"
)

type ReturnAuthStruct struct {
	Jti          int64 `json:"jti,string"`
	IsAdmin      bool  `json:"isAdmin"`
	Is2PassValid bool  `json:"is2PassValid"`
}

func (a ReturnAuthStruct) SetContext(r *http.Request) *http.Request {
	rCtx := context.WithValue(r.Context(), KeyUID, a.Jti)
	rCtx = context.WithValue(rCtx, KeyIsAdmin, a.IsAdmin)
	rCtx = context.WithValue(rCtx, KeyIs2PassValid, a.Is2PassValid)
	return r.WithContext(rCtx)
}

// HeaderAuthorization name of the `Authorization` header.
const (
	// DEPRECATED
	Header = "Authorization"
	// DEPRECATED
	JWTHeader           = "jwt"
	HeaderAuthorization = "Authorization"
	HeaderJWT           = "jwt"
	Header2Password     = "X-2Pass"
)

type CtxKey string

const (
	KeyUID          CtxKey = "key_uid"
	KeyIsAdmin      CtxKey = "key_isAdmin"
	KeyIs2PassValid CtxKey = "key_is2PassValid"
)

var userApiLink string

func Init(usrApiLink string) {
	userApiLink = usrApiLink
}

// AuthtokenHeader extracts from the `http.Request` Authorization header.
func AuthtokenHeader(r *http.Request) string {
	return r.Header.Get(HeaderAuthorization)
}

// ExtractAuthData extracts `ReturnAuthStruct` from request.
func ExtractAuthData(r *http.Request) (res ReturnAuthStruct, ok bool) {
	jwtData := r.Header.Get(HeaderJWT)
	if jwtData == "" {
		return
	}

	err := json.Unmarshal([]byte(jwtData), &res)
	ok = err == nil
	return
}

// GetUID extracts User-ID  from the `http.Request` ctx.
func GetUID(r *http.Request) int64 {
	uid, _ := r.Context().Value(KeyUID).(int64)
	return uid
}

// IsAdmin return `true` if request sent by admin.
func IsAdmin(r *http.Request) bool {
	isAdmin, _ := r.Context().Value(KeyIsAdmin).(bool)
	return isAdmin
}

// Is2PassValid return `true` if user used to 2nd password for request auth.
func Is2PassValid(r *http.Request) bool {
	is2PassValid, _ := r.Context().Value(KeyIs2PassValid).(bool)
	return is2PassValid
}

// CheckToken checks `Authorization` token if it valid return nil.
func CheckToken(authtoken string) (int, []byte, error) {
	if userApiLink == "" {
		log.Default.Error("auth didn't init")
	}
	client := http.DefaultClient
	path := userApiLink + "/v1/auth"

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return http.StatusInternalServerError,
			nil, errors.Wrap(err, "failed to create auth check request")
	}

	req.Header.Set(HeaderAuthorization, authtoken)

	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError,
			nil, errors.Wrap(err, "failed to check token")
	}
	if resp.StatusCode == 200 {
		return 200, nil, nil
	}

	defer resp.Body.Close()
	respBody := make([]byte, 0)

	_, err = resp.Body.Read(respBody)
	if err != nil {
		log.Default.WithError(err).Error("unable to read response body")
		return http.StatusInternalServerError,
			nil, errors.Wrap(err, "failed read auth response body")
	}

	return resp.StatusCode, respBody, nil
}

// ValidateAuthHeader checks the request Authorization token.
// If token valid - continue request handling flow,
// else redirect `userapi` response to the requester.
func ValidateAuthHeader(required bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authtoken := r.Header.Get(HeaderAuthorization)
			if authtoken == "" || !required {
				next.ServeHTTP(w, r)
				return
			}

			statusCode, rawResp, err := CheckToken(authtoken)
			if statusCode == http.StatusOK {
				next.ServeHTTP(w, r)
				return
			}

			if err != nil {
				log.Default.WithError(err).Error("unable to check auth token")
			}

			w.WriteHeader(statusCode)
			w.Write(rawResp)
		})
	}
}

// Method reads JWT HeaderAuthorization and fill KeyUID and KeyIsAdmin in the context
// Use ExtractUserID() if jwt required
// Use ExtractUserID(false) if jwt not required
func ExtractUserID(required ...bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwt := ReturnAuthStruct{}
			rawJwt := r.Header.Get(HeaderJWT)
			if rawJwt == "" {
				if len(required) == 0 || (len(required) > 0 && required[0]) {
					render.ResultBadRequest.
						SetError("JWT HeaderAuthorization must not bee empty").
						Render(w)
					return
				}

				r = jwt.SetContext(r)
				next.ServeHTTP(w, r)
				return
			}

			err := json.Unmarshal([]byte(rawJwt), &jwt)
			if err != nil {
				render.ResultBadRequest.
					SetError("JWT HeaderAuthorization is invalid json").
					Render(w)
				return
			}

			r = jwt.SetContext(r)
			next.ServeHTTP(w, r)
			return

		})
	}
}
