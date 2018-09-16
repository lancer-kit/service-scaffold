package notifier

import (
	"encoding/json"

	"github.com/nats-io/go-nats"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type (
	// Interface for Sender
	SenderI interface {
		SetLogger(l *logrus.Entry) *Sender
		SetConfig(c *Config) *Sender
		Disconnect() *Sender
		Send(m *Message) error
	}

	// Working with the Notification Service
	Sender struct {
		cfg    *Config
		conn   *nats.Conn
		logger *logrus.Entry
	}

	//Sender configuration
	Config struct {
		Chanel   string //chanel (topic) name. default = notifier
		Url      string //NATS connection url
		User     string //NATS user
		Password string //Nats password
		Token    string //Nats auth token
	}

	// Main message structure
	Message struct {
		UserId      int         `json:"userId"`         //User ID for notifications
		Command     string      `json:"command"`        // "Command" for front-end
		IsBroadcast bool        `json:"isBroadcast"`    //Private or broadcast message
		Data        interface{} `json:"Data,omitempty"` //Optional data
	}

	// Optional data structure
	Data struct {
		IsError      bool        `json:"isError"`             //is error
		ErrorMessage string      `json:"errorMessage"`        //message when error
		ErrorKind    int         `json:"errorKind,omitempty"` //code / kind of error
		Data         interface{} `json:"data,omitempty"`      // Message data
	}
)

const (
	DefaultChanel = "notifier" //Default chanel (subj)
)

var Default *Sender

func emptyOption() nats.Option {
	return func(o *nats.Options) error {
		return nil
	}
}

//NewSender - initialize new sender structure and try to connect with NATS sever
func NewSender(cfg *Config) (*Sender, error) {
	if cfg == nil {
		cfg = &Config{
			Chanel: DefaultChanel,
			Url:    nats.DefaultURL,
		}
	}

	var op nats.Option
	op = emptyOption()
	if cfg.User != "" {
		op = nats.UserInfo(cfg.User, cfg.Password)
	}
	if cfg.Token != "" {
		op = nats.Token(cfg.Token)
	}
	nc, err := nats.Connect(cfg.Url, op)

	s := &Sender{
		cfg:  cfg,
		conn: nc,
	}

	if s.cfg.Url == "" {

	}

	return s, err
}

//SetLogger - set logger
func (s *Sender) SetLogger(l *logrus.Entry) *Sender {
	s.logger = l
	return s
}

//Disconnect - disconnect
func (s *Sender) Disconnect() *Sender {
	if s.conn == nil {
		return s
	}
	s.conn.Close()
	return s
}

func (s *Sender) checkConn() (err error) {
	if s.conn == nil {
		s.conn, err = nats.Connect(s.cfg.Url)
		return
	}
	if s.conn.Status() != nats.CONNECTED {
		s.conn.Close()
		s.conn, err = nats.Connect(s.cfg.Url)
		if err != nil {
			return
		}
	}

	return
}

//Send - main functional.Publish message to service
func (s *Sender) Send(msg *Message) (err error) {
	b, xe := json.Marshal(msg)
	if xe != nil {
		s.ErrorLog(xe, "unable to marshal message in Sender.Publish")
		err = errors.Wrap(xe, "unable to marshal message")
		return
	}
	if err = s.checkConn(); err != nil {
		s.ErrorLog(err, "unable to connect")
		err = errors.Wrap(xe, "unable to connect to NATS server")
		return
	}

	err = s.conn.Publish(s.cfg.Chanel, b)
	if err != nil {
		s.ErrorLog(err, "unable to publish message in Sender.Publish")
		err = errors.Wrap(err, "unable to publish message")
		return
	}

	err = s.conn.Flush()
	if err != nil {
		s.ErrorLog(err, "unable to flush connection in Sender.Publish")
		err = errors.Wrap(err, "unable to flush connection")
	}
	return
}

//IsConnected - check is connected to NATS
func (s *Sender) IsConnected() (err error) {
	err = s.checkConn()
	if err != nil {
		s.ErrorLog(err, "NATS connection error")
	}
	return err
}

//SetConfig - set new configuration
func (s *Sender) SetConfig(c *Config) *Sender {
	old := s.cfg
	err := s.IsConnected()
	if err != nil {
		old = nil
	}
	s.cfg = c

	err = s.IsConnected()
	if err == nil {
		return s
	}

	s.ErrorLog(err, "Invalid new configuration for Sender")
	if old != nil {
		s.cfg = old
		s.ErrorLog(nil, "Old configuration for Sender restored")
	}

	return s
}

func (s Sender) ErrorLog(err error, msg string) {
	if s.logger == nil {
		return
	}
	s.logger.WithError(err).Error(msg)
}

func init() {
	Default, _ = NewSender(nil)
}
