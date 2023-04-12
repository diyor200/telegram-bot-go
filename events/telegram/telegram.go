package telegram

import "gobot/gobot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}
