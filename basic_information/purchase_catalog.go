package basic_information

type PurchaseCatalog interface {
	Add(partner CooperationPartner)
	Remove(partner CooperationPartner)
	Get(id int64) CooperationPartner
}
