package exporter

import (
	"github.com/k0kubun/pp"
	log "github.com/sirupsen/logrus"
)

type Exporter struct {
	Kind   string
	Config map[string]interface{}
	Key    string
	Topic  string
}

func New(config map[string]interface{}) *Exporter {

	return &Exporter{
		Kind:   config["export"].(map[string]interface{})["kind"].(string),
		Config: config["export"].(map[string]interface{}),
		Key:    config["export"].(map[string]interface{})["key"].(string),
		Topic:  config["export"].(map[string]interface{})["topic"].(string),
	}
}

func (s *Exporter) Export(messages []interface{}, props map[string]string) {

	if s.Config["enabled"].(bool) {
		log.Info("Exporter enabled => " + s.Kind)
		switch s.Kind {
		case "rabbitmq":
			s.RabbitMQExport(messages, s.Topic, props)
		default:
			pp.Println("No information available for this exporter")
		}
	} else {
		log.Info("Exporter disabled")
	}

}
