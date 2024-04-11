package user

import "github.com/krivenkov/pkg/busapi/topics"

const UpdateUserTopic topics.Topic = "user.update.user.1"

type User struct {
	ID      string
	Disable bool
}
