package accountancy

type Received interface {
	Add(serial TransactionSerial)
	Remove(serial TransactionSerial)
	GetShould() float64
	GetActual() float64
}