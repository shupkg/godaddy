package godaddy

import "time"

// DomainSummary 域名信息封装
type DomainSummary struct {
	DomainContacts

	DomainID               int64     `json:"domainId"`
	Domain                 string    `json:"domain"`
	Status                 string    `json:"status"`
	Expires                time.Time `json:"expires"`                //2016-06-14T23:59:59.000Z
	CreatedAt              time.Time `json:"createdAt"`              //创建时间
	DeletedAt              time.Time `json:"deletedAt"`              //删除时间
	TransferAwayEligibleAt time.Time `json:"transferAwayEligibleAt"` //Date and time when this domain is eligible to transfer
	RenewDeadline          time.Time `json:"renewDeadline"`          //续费最终时间
	Renewable              bool      `json:"renewable"`              //可以续费
	RenewAuto              bool      `json:"renewAuto"`              //自动续费
	Privacy                bool      `json:"privacy"`                //whois 隐私保护
	Locked                 bool      `json:"locked"`                 //锁定
	HoldRegistrar          bool      `json:"holdRegistrar"`          //注册商转移锁定
	ExposeWhois            bool      `json:"exposeWhois"`            //是否暴露 whois
	ExpirationProtected    bool      `json:"expirationProtected"`    //到期保护

	AuthCode    string   `json:"authCode,omitempty"`    //转移密码
	NameServers []string `json:"nameServers,omitempty"` //解析服务器
}

//DomainPurchase 购买域名结构
type DomainPurchase struct {
	DomainContacts
	Domain      string   `json:"domain"`
	Consent     Consent  `json:"consent"`
	RenewAuto   bool     `json:"renewAuto"`
	Period      int64    `json:"period,omitempty"`
	Privacy     bool     `json:"privacy,omitempty"`
	NameServers []string `json:"nameServers,omitempty"`
}

//DomainAvailableBulk 域名检查是否可用结果
type DomainAvailableBulk struct {
	Domain     string `json:"domain"`
	Period     int    `json:"period"`
	Price      int    `json:"price"`
	Available  bool   `json:"available"`
	Currency   string `json:"currency"`
	Definitive bool   `json:"definitive"`
}

//LegalAgreement LegalAgreement
type LegalAgreement struct {
	AgreementKey string `json:"agreementKey"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Url          string `json:"url"`
}

//DomainsContactsBulk DomainsContactsBulk
type DomainsContactsBulk struct {
	DomainContacts
	Domains    []string `json:"domains"`              //An array of domain names to be validated against. Alternatively, you can specify the extracted tlds. However, full domain names are required if the tld is uk
	EntityType string   `json:"entityType,omitempty"` //Canadian Presence Requirement (CA)
}
