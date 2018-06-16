package accountancy

type ReceivedCatalog interface {
	Add(received Received)
	Confirm(received Received)
}
