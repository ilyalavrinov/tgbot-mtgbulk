package bot

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/admirallarimda/tgbotbase"
	"github.com/ilyalavrinov/tgbot-mtgbulkbuy/internal/log"
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
	resp, err := http.Post("http://127.0.0.1:8000/bulk", "text/plain", strings.NewReader(msg.Text))

	var replyMsg string
	if err != nil {
		replyMsg = err.Error()
	} else {
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			replyMsg = "Something went wrong, please contact the admin"
			log.Errorw("Error while reading response", "err", err)
		} else {
			replyMsg = string(b)
		}
	}

	reply := tgbotapi.NewMessage(msg.Chat.ID, replyMsg)
	reply.BaseChat.ReplyToMessageID = msg.MessageID
	h.OutMsgCh <- reply
}
