package contract

import "context"

type NotificationService interface {
	Publish(context.Context, string) error
}
