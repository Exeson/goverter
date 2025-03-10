// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package generated

import house "github.com/jmattheis/goverter/example/house"

type ConverterImpl struct{}

func (c *ConverterImpl) ConvertHouse(source house.DBHouse) house.APIHouse {
	var houseAPIHouse house.APIHouse
	houseAPIHouse.Address = source.Address
	mapHouseAPIRoomNRHouseAPIApartment := make(map[house.APIRoomNR]house.APIApartment, len(source.Apartments))
	for key, value := range source.Apartments {
		mapHouseAPIRoomNRHouseAPIApartment[house.APIRoomNR(key)] = c.houseDBApartmentToHouseAPIApartment(value)
	}
	houseAPIHouse.Apartments = mapHouseAPIRoomNRHouseAPIApartment
	return houseAPIHouse
}
func (c *ConverterImpl) ConvertPerson(source house.DBPerson) house.APIPerson {
	var houseAPIPerson house.APIPerson
	houseAPIPerson.ID = source.ID
	houseAPIPerson.MiddleName = house.SQLStringToPString(source.MiddleName)
	pString := source.Name
	houseAPIPerson.FirstName = &pString
	houseAPIPersonList := make([]house.APIPerson, len(source.Friends))
	for i := 0; i < len(source.Friends); i++ {
		houseAPIPersonList[i] = c.ConvertPerson(source.Friends[i])
	}
	houseAPIPerson.Friends = houseAPIPersonList
	return houseAPIPerson
}
func (c *ConverterImpl) houseDBApartmentToHouseAPIApartment(source house.DBApartment) house.APIApartment {
	var houseAPIApartment house.APIApartment
	houseAPIApartment.Position = source.Position
	houseAPIApartment.Owner = c.ConvertPerson(source.Owner)
	houseAPIPersonList := make([]house.APIPerson, len(source.CoResident))
	for i := 0; i < len(source.CoResident); i++ {
		houseAPIPersonList[i] = c.ConvertPerson(source.CoResident[i])
	}
	houseAPIApartment.CoResident = houseAPIPersonList
	return houseAPIApartment
}
