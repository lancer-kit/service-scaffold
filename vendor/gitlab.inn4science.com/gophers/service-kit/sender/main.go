// sender is a client package for sender service.
// It allow to make and send emails.
// `Message` can be sent through the HTTP or NATS.
package sender

import "github.com/go-ozzo/ozzo-validation"

// LetterType is an enum of predefined letter templates.
type LetterType int

// Provider is an enum of predefined providers for OTP.
type Provider int

const (
	LetterUniversal LetterType = 1 + iota
	LetterAdminSignUp
	LetterUserEmailVerify
	LetterUserRecovery
	LetterUserNewDevice
)

const (
	ViberProvider Provider = 1 + iota
	WAProvider
	SMSProvider
	TelegramProvider
	TestProvider Provider = 100
)

// NATSTopic is a topic in NATS, through which the sender receives new messages.
const NATSTopic = "sender.letters"

// OTPTopic is a topic in NATS, through which the sender receives new otp messages.
const OTPTopic = "sender.otp"

// HTTPURL is URL path in which the sender receives new messages.
const HTTPURL = "/v1/email"

// Message is the data for some template.
// The type field indicates which template will be sent.
type Message struct {
	// Type indicates which template will be used.
	Type LetterType `json:"type"`
	// Data to fill in the template, depends on the `Type`.
	Data MsgData `json:"data"`
}

func (m Message) Validate() (err error) {
	switch m.Type {
	case LetterUniversal:
		err = m.Data.Universal.Validate()
	case LetterAdminSignUp:
		err = m.Data.Base.Validate()
	case LetterUserEmailVerify:
		err = m.Data.Base.Validate()
	case LetterUserRecovery:
		err = m.Data.Base.Validate()
	case LetterUserNewDevice:
		err = m.Data.Base.Validate()
		if err != nil {
			return
		}
		err = m.Data.Device.Validate()
	}
	return
}

// MsgData data for letter templates.
type MsgData struct {
	Base      Base      `json:"base,omitempty"`
	Device    Device    `json:"device,omitempty"`
	Universal Universal `json:"universal,omitempty"`
}

func (ms MsgData) Validate() error {
	return validation.ValidateStruct(&ms,
		validation.Field(&ms.Device),
		validation.Field(&ms.Universal),
		validation.Field(&ms.Base),
	)
}

// Base is a structure for the base letter template.
type Base struct {
	// Email is a addressee of the letter.
	Email    string `json:"email"`
	Username string `json:"username"`
	Link     string `json:"link"`
}

func (b Base) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Email, validation.Required),
		validation.Field(&b.Link, validation.Required),
	)
}

// Device is a data extension for the `new device` letter.
type Device struct {
	Device   string `json:"device"`
	Location string `json:"location"`
	Ip       string `json:"ip"`
}

func (dv Device) Validate() error {
	return validation.ValidateStruct(&dv,
		validation.Field(&dv.Device, validation.Required),
		validation.Field(&dv.Location, validation.Required),
		validation.Field(&dv.Ip, validation.Required),
	)
}

// Universal is a message that does not have a template
// in the sender and must be sent as is. Data from the another fields
// of the MsgData will be ignored.
type Universal struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
	HTML    string `json:"html"`
}

func (un Universal) Validate() error {
	return validation.ValidateStruct(&un,
		validation.Field(&un.Email, validation.Required),
		validation.Field(&un.Subject, validation.Required),
	)
}

// OTPMessage is a structure for message sent through NATS used to send an OTP.
type OTPMessage struct {
	Phone    string   `json:"phone,omitempty"`
	Code     string   `json:"code,omitempty"`
	Provider Provider `json:"provider,omitempty"`
}

// Validate() function returns an error if the data in an OTPMessage is invalid.
func (otpm OTPMessage) Validate() error {
	return validation.ValidateStruct(&otpm,
		validation.Field(&otpm.Phone, validation.Required),
		validation.Field(&otpm.Code, validation.Required),
		validation.Field(&otpm.Provider, validation.Required),
	)
}
