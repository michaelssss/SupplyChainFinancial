package credit_management

import "FinacialSupplyChain/basic_information"

type CreditManager interface {
	Add(partner basic_information.CooperationPartner)
	UseCredit(partner basic_information.CooperationPartner) bool
}
