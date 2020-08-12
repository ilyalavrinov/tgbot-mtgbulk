package bot

import (
	"regexp"

	"github.com/admirallarimda/tgbotbase"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type searchHandler struct {
	tgbotbase.BaseHandler
}

func NewSearchHandler() tgbotbase.IncomingMessageHandler {
	handler := searchHandler{}
	return &handler
}

func (h *searchHandler) Init(outMsgCh chan<- tgbotapi.Chattable, srvCh chan<- tgbotbase.ServiceMsg) tgbotbase.HandlerTrigger {
	h.OutMsgCh = outMsgCh
	return tgbotbase.NewHandlerTrigger(regexp.MustCompile(".*"), nil)
}

func (h *searchHandler) Name() string {
	return "bulk_search"
}

func (h *searchHandler) HandleOne(msg tgbotapi.Message) {
	_ = msg.Text

	replyMsg := "Out of Order, come again later"

	reply := tgbotapi.NewMessage(msg.Chat.ID, replyMsg)
	reply.BaseChat.ReplyToMessageID = msg.MessageID
	h.OutMsgCh <- reply
}
