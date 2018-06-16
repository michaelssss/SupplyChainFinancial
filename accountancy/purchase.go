package accountancy

type Purchase interface {
	Add(serial TransactionSerial)
	Remove(serial TransactionSerial)
	GetShould() float64
	GetActual() float64
}
