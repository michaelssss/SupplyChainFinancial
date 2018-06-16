/*
this package contain the partner information maintain operation
 */
package basic_information

type CooperationPartnerCatalog interface {
	Add(partner CooperationPartner)
	Remove(partner CooperationPartner)
	Permit(partner CooperationPartner)
}
