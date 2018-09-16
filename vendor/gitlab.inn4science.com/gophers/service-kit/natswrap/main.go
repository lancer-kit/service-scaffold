package natswrap

import (
	"encoding/json"

	"github.com/nats-io/go-nats"
	"github.com/pkg/errors"
)

var natsConn *nats.Conn
var cfg *Config

func SetConfig(config *Config) {
	cfg = config
}

func GetConn() (*nats.Conn, error) {
	if natsConn != nil {
		return natsConn, nil
	}

	if cfg == nil {
		return nil, errors.New("Nats config didn't set")
	}
	var err error
	natsConn, err = nats.Connect(
		cfg.ToURL(),
		nats.UserInfo(
			cfg.User,
			cfg.Password),
	)
	if err != nil {
		return nil, err
	}
	return natsConn, nil
}

func PublishMessage(topic string, msg interface{}) error {
	rawMsg, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err,
			"unable to marshal message for the topic - "+topic)
	}

	_, err = GetConn()
	if err != nil {
		return errors.Wrap(err, "unable to establish connection")
	}

	err = natsConn.Publish(topic, rawMsg)
	if err != nil {
		return errors.Wrap(err, "unable to publish into the topic "+topic)
	}

	return nil
}

func Subscribe(topic string, msgs chan *nats.Msg) (*nats.Subscription, error) {
	_, err := GetConn()
	if err != nil {
		return nil, errors.Wrap(err, "unable to establish connection")
	}

	subscription, err := natsConn.ChanSubscribe(topic, msgs)
	if err != nil {
		return nil, errors.Wrap(err, "unable to subscribe into the topic "+topic)
	}
	return subscription, nil
}
