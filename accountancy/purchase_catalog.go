package accountancy

type PurchaseCatalog interface {
	Add(purchase Purchase)
	Confirm(purchase Purchase)
}