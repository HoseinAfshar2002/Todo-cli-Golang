package contract

import "Todo-Cli-With-Golang/entity"

type UserWriteStore interface {
	Save(u entity.User)
}
type UserReadStore interface {
	Load() []entity.User
}