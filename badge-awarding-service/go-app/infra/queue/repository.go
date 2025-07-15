package queue

type MailQueuePublisher interface {
	PublishMailMessage(message string) error
}

type MailQueue struct {
	config Config
}
