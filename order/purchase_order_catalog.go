package order

type PurchaseOrderCatalog interface {
	Add(order PurchaseOrder)
	Remove(order PurchaseOrder)
	Get(id int64) PurchaseOrder
}
