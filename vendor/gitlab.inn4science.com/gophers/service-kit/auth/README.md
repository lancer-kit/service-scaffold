# auth
--
    import "gitlab.inn4science.com/gophers/service-kit/auth"


## Usage

```
const (
	Header    = "Authorization"
	JWTHeader = "jwt"

	KeyUID = iota
)
```
Header name of the `Authorization` header.

```
const (
	HeaderHash        = "X-Auth-Hash"
	HeaderSignature   = "X-Auth-Signature"
	HeaderSigner      = "X-Auth-Signer"
	HeaderService     = "X-Auth-Service"
	HeaderContentType = "Content-Type"
)
```

#### func  CheckToken

```
func CheckToken(authtoken string) (int, []byte, error)
```
CheckToken checks `Authorization` token if it valid return nil.

#### func  ExtractUserID

```
func ExtractUserID() func(http.Handler) http.Handler
```

#### func  Init

```
func Init(usrApiLink string)
```

#### func  NewSignedGetRequest

```
func NewSignedGetRequest(privateKey, path, service string) (*http.Request, error)
```
NewSignedGetRequest creates a new GET request, sings the request details using
the `privateKey` and adds the auth headers.

#### func  NewSignedPostRequest

```
func NewSignedPostRequest(privateKey, path string, body []byte, mimeType, service string) (*http.Request, error)
```
NewSignedPostRequest creates a new POST request, hashes the body, sings the
request details using the `privateKey` and adds the auth headers.

#### func  ValidateAuthHeader

```
func ValidateAuthHeader(required bool) func(http.Handler) http.Handler
```
ValidateAuthHeader checks the request Authorization token. If token valid -
continue request handling flow, else redirect `userapi` response to the
requester.

#### func  VerifyRequestSignature

```
func VerifyRequestSignature(r *http.Request, publicKey string) (bool, error)
```
VerifyRequestSignature checks the request auth headers.
