package main

import (
	"github.com/sajosam/rabbitmq-exporter/exporter"
)

func main() {
	exp := exporter.New(map[string]interface{}{
		"export": map[string]interface{}{
			"kind":     "rabbitmq",
			"key":      "rabbit",
			"topic":    "op-org-14",
			"host":     "139.59.58.203",
			"port":     "5672",
			"enabled":  true,
			"password": "onepane",
			"username": "onepane",
		},
	})

	message := map[string]interface{}{
		"event_id":      "27879",
		"event_message": "demo-2 Triggered from refs/heads/main",
		"event_time":    1717397708000,
		"tags": map[string]interface{}{
			"event.user":       "Sajo Sam",
			"event.category":   "deployment_events",
			"event.finishtime": 1717397708000.000000,
			"event.queuetime":  1717397708000.000000,
			"event.starttime":  1717397708000.000000,
			"event.state":      "completed",
			"event.type":       "azuredevops.build",
			"event.url":        "https://dev.azure.com/onepane/5c595676-b99c-47c6-b06e-1899293219c7/_apis/build/Builds/2850",
			"event.from":       "refs/heads/main",
			"event.repo":       "62885e76-3ead-40e1-aaca-101ae6174e4b",
		},
		"source":      7,
		"resource_id": "33",
	}
	exp.Export([]interface{}{message}, map[string]string{
		"event": "open",
	})
}
