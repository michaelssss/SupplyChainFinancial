package subject_matter

import "FinacialSupplyChain/basic_information"

type Stock interface {
	NewOwner(partner basic_information.CooperationPartner)
	AddCargo(cargo Cargo)
	Remove(cargo Cargo)
}
