package basic_information

type SupplierCatalog interface {
	Add(partner CooperationPartner)
	Remove(partner CooperationPartner)
	Get(id int64) CooperationPartner
}
