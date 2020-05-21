package mock

import (
	"github.com/bxcodec/faker/v3"
)

type Mock struct {
}

func (m Mock) AmountWithCurrency() string {
	return faker.AmountWithCurrency()
}

func (m Mock) Currency() string {
	return faker.Currency()
}

func (m Mock) CreditCardNumber() string {
	return faker.CCNumber()
}

func (m Mock) CreditCardType() string {
	return faker.CCType()
}

func (m Mock) Century() string {
	return faker.Century()
}

func (m Mock) Date() string {
	return faker.Date()
}

func (m Mock) DayOfMonth() string {
	return faker.DayOfMonth()
}

func (m Mock) DayOfWeek() string {
	return faker.DayOfWeek()
}

func (m Mock) YearString() string {
	return faker.YearString()
}

func (m Mock) E164PhoneNumber() string {
	return faker.E164PhoneNumber()
}

func (m Mock) Name() string {
	return faker.Name()
}

func (m Mock) Username() string {
	return faker.Username()
}

func (m Mock) LastName() string {
	return faker.LastName()
}

func (m Mock) FirstName() string {
	return faker.FirstName()
}

func (m Mock) FirstNameFemale() string {
	return faker.FirstNameFemale()
}

func (m Mock) FirstNameMale() string {
	return faker.FirstNameMale()
}

func (m Mock) DomainName() string {
	return faker.DomainName()
}

func (m Mock) Email() string {
	return faker.Email()
}

func (m Mock) IPv4() string {
	return faker.IPv4()
}

func (m Mock) IPv6() string {
	return faker.IPv6()
}

func (m Mock) MacAddress() string {
	return faker.MacAddress()
}

func (m Mock) URL() string {
	return faker.URL()
}

func (m Mock) Latitude() float64 {
	return faker.Latitude()
}

func (m Mock) Longitude() float64 {
	return faker.Latitude()
}

func (m Mock) MonthName() string {
	return faker.MonthName()
}

func (m Mock) Paragraph() string {
	return faker.Paragraph()
}

func (m Mock) Password() string {
	return faker.Password()
}

func (m Mock) PhoneNumber() string {
	return faker.Phonenumber()
}

func (m Mock) RandomUnixTime() int64 {
	return faker.RandomUnixTime()
}

func (m Mock) TimeString() string {
	return faker.TimeString()
}

func (m Mock) Timeperiod() string {
	return faker.Timeperiod()
}

func (m Mock) Timestamp() string {
	return faker.Timestamp()
}

func (m Mock) Timezone() string {
	return faker.Timezone()
}

func (m Mock) TitleFemale() string {
	return faker.TitleFemale()
}

func (m Mock) TitleMale() string {
	return faker.TitleMale()
}

func (m Mock) TollFreePhoneNumber() string {
	return faker.TollFreePhoneNumber()
}

func (m Mock) UUIDDigit() string {
	return faker.UUIDDigit()
}

func (m Mock) UUIDHyphenated() string {
	return faker.UUIDHyphenated()
}

func (m Mock) UnixTime() int64 {
	return faker.UnixTime()
}

func (m Mock) Word() string {
	return faker.Word()
}

func (m Mock) Sentence() string {
	return faker.Sentence()
}
