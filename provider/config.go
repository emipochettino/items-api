package provider

import "os"

const (
	queueUrl  = "QUEUE_URL"
	queueName = "QUEUE_NAME"
)

var QProvider IQueueProvider

func init() {
	configureProviders()
	initializeProviders()
}

//Set default environment variables
func configureProviders() {
	if url := os.Getenv(queueUrl); len(url) == 0 {
		os.Setenv(queueUrl, "amqp://guest:guest@localhost:5672/")
	}

	if name := os.Getenv(queueName); len(name) == 0 {
		os.Setenv(queueName, "ITEM_NOTIFICATION")
	}
}

//initialize all providers
func initializeProviders() {
	url := os.Getenv(queueUrl)
	name := os.Getenv(queueName)
	QProvider = QueueProvider{Url: url, QueueName: name}
}
