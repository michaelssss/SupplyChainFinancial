package order

type EntrustedOrderCatalog interface {
	Add(order EntrustedOrder)
	Remove(order EntrustedOrder)
	Get(id int64) EntrustedOrder
}
