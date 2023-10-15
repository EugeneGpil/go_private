package go_private

type TheStruct struct {
	theProperty string
}

type NewTheStructDto struct {
	TheProperty string
}

func (theStruct *TheStruct) GetProperty() string {
	return theStruct.theProperty
}

func NewTheStruct(dto NewTheStructDto) TheStruct {
	return TheStruct{
		theProperty: dto.TheProperty,
	}
}
