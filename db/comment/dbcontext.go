package comment

type context struct {
	Messages messagesTable
}

var ctx context

func NewContext() {
	ctx = context{
		Messages: NewMessagesTable(),
	}
}
