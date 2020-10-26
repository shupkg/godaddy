package godaddy

import "time"

//Consent 承诺
type Consent struct {
	AgreedAt      time.Time `json:"agreedAt"`
	AgreedBy      string    `json:"agreedBy"`
	AgreementKeys []string  `json:"agreementKeys"`
}

//Contact 联系方式
type Contact struct {
	NameFirst      string         `json:"nameFirst"`
	NameMiddle     string         `json:"nameMiddle"`
	NameLast       string         `json:"nameLast"`
	JobTitle       string         `json:"jobTitle"`
	Organization   string         `json:"organization"`
	Phone          string         `json:"phone"`
	Fax            string         `json:"fax"`
	Email          string         `json:"email"`
	AddressMailing AddressMailing `json:"addressMailing"`
}

type DomainContacts struct {
	ContactAdmin      *Contact `json:"contactAdmin,omitempty"`
	ContactBilling    *Contact `json:"contactBilling,omitempty"`
	ContactRegistrant *Contact `json:"contactRegistrant,omitempty"`
	ContactTech       *Contact `json:"contactTech,omitempty"`
}

//AddressMailing 联系方式邮政地址
type AddressMailing struct {
	Country    string `json:"country"`
	State      string `json:"state"`
	City       string `json:"city"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
	PostalCode string `json:"postalCode"`
}

//OrderResult 订单结果
type OrderResult struct {
	OrderID   int    `json:"orderId"`   //Unique identifier of the order processed to purchase the domain
	Currency  string `json:"currency"`  //($iso-currency-code), default: USD, Currency in which the total is listed
	Total     int    `json:"total"`     //Total cost of the domain and any selected add-ons
	ItemCount int    `json:"itemCount"` //Number items included in the order
}

//DNSRecord dns 记录
type DNSRecord struct {
	Protocol string `json:"protocol,omitempty"`
	Data     string `json:"data,omitempty"`
	Port     int    `json:"port,omitempty"`
	Service  string `json:"service,omitempty"`
	Name     string `json:"name,omitempty"`
	Weight   int    `json:"weight,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Type     string `json:"type,omitempty"`
	Ttl      int    `json:"ttl,omitempty"`
}
