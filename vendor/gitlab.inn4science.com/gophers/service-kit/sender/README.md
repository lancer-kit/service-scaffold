# sender

---
    import "gitlab.inn4science.com/gophers/service-kit/sender"

*sender* is a client package for sender service. It allow to make and send emails.
`Message` can be sent through the HTTP or NATS.

## Usage

``` go
const HTTPURL = "/v1/email"
```
HTTPURL is URL path in which the sender receives new messages.

``` go
const NATSTopic = "sender.letters"
```
NATSTopic is a topic in NATS, through which the sender receives new messages.

#### type Base

``` go
type Base struct {
	// Email is a addressee of the letter.
	Email    string `json:"email"`
	Username string `json:"username"`
	Link     string `json:"link"`
}
```

Base is a structure for the base letter template.

#### type Device

``` go
type Device struct {
	Device   string `json:"device"`
	Location string `json:"location"`
	Ip       string `json:"ip"`
}
```

Device is a data extension for the `new device` letter.

#### type LetterType

``` go
type LetterType int
```

LetterType is an enum of predefined letter templates.

``` go
const (
	LetterUniversal LetterType = 1 + iota
	LetterAdminSignUp
	LetterUserEmailVerify
	LetterUserRecovery
	LetterUserNewDevice
)
```

#### type Message

``` go
type Message struct {
	// Type indicates which template will be used.
	Type LetterType `json:"type"`
	// Data to fill in the template, depends on the `Type`.
	Data MsgData `json:"data"`
}
```

Message is the data for some template. The type field indicates which template
will be sent.

#### type MsgData

``` go
type MsgData struct {
	Base      Base      `json:"base,omitempty"`
	Device    Device    `json:"device,omitempty"`
	Universal Universal `json:"universal,omitempty"`
}
```

MsgData data for letter templates.

#### type Universal

``` go
type Universal struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
	HTML    string `json:"html"`
}
```

Universal is a message that does not have a template in the sender and must be
sent as is. Data from the another fields of the MsgData will be ignored.



#### type OTPMessage

```go
type OTPMessage struct {
	Phone    string   `json:"phone,omitempty"`
	Code     string   `json:"code,omitempty"`
	Provider Provider `json:"provider,omitempty"`
}
```

OTPMessage is a structure for message sent through NATS used to send an OTP.



#### func (OTPMessage) Validate() error

Validate() function returns an error if the data in an OTPMessage is invalid.