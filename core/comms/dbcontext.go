package comms

type context struct {
	Messages messagesTable
}

var ctx context

func init() {
	ctx = context{
		Messages: NewMessagesTable(),
	}
}
