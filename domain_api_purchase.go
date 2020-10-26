package godaddy

import (
	"context"
	"strconv"
)

//GetTLDs Retrieves a list of TLDs supported and enabled for sale
func (c *DomainClient) GetTLDs(ctx context.Context) (map[string]string, error) {
	var result []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	_, err := c.get("/tlds").Output(&result).Run(ctx)
	if err != nil {
		return nil, err
	}
	var out = map[string]string{}
	for _, item := range result {
		out[item.Name] = item.Type
	}
	return out, nil
}

//Retrieve the legal agreement(s) required to purchase the specified TLD and add-ons
func (c *DomainClient) Agreements(ctx context.Context, tld []string, privacy bool, forTransfer bool) (result []LegalAgreement, err error) {
	q := Q("tlds", tld).Set("privacy", privacy).Set("forTransfer", forTransfer)
	_, err = c.get("/agreements", q).Output(&result).Run(ctx)
	return
}

//DomainAvailable Determine whether or not the specified domain is available for purchase
func (c *DomainClient) DomainAvailable(ctx context.Context, domain string, checkType string, forTransfer bool) (*DomainAvailableBulk, error) {
	if checkType == "" {
		checkType = "FAST"
	}

	var result DomainAvailableBulk
	q := Q("domain", domain).
		Set("checkType", checkType).
		Set("forTransfer", strconv.FormatBool(forTransfer))
	_, err := c.get("/available", q).Output(&result).Run(ctx)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//DomainsAvailable Determine whether or not the specified domains are available for purchase
func (c *DomainClient) DomainsAvailable(ctx context.Context, domains []string, checkType string) (result []*DomainAvailableBulk, err error) {
	if checkType == "" {
		checkType = "FAST"
	}
	_, err = c.post("/available", Q("checkType", checkType)).Input(domains).Output(&result).Run(ctx)
	return
}

//ContactsValidate Validate the request body using the Domain Contact Validation Schema for specified domains.
func (c *DomainClient) ContactsValidate(ctx context.Context, body DomainsContactsBulk) error {
	_, err := c.post("/contacts/validate").Input(body).Run(ctx)
	return err
}

//Purchase and register the specified Domain
func (c *DomainClient) Purchase(ctx context.Context, body DomainPurchase) (*OrderResult, error) {
	// /v1/domains/purchase
	var result OrderResult
	_, err := c.post("/purchase").Input(body).Output(&result).Run(ctx)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//PurchaseValidate Validate the request body using the Domain Purchase Schema for the specified TLD
func (c *DomainClient) PurchaseValidate(ctx context.Context, body DomainPurchase) error {
	// /v1/domains/purchase/validate
	_, err := c.post("/purchase/validate").Input(body).Run(ctx)
	return err
}

//CancelPurchased Cancel a purchased domain
func (c *DomainClient) CancelPurchased(ctx context.Context, domain string) (result []DNSRecord, err error) {
	_, err = c.delete("/" + domain).Output(&result).Run(ctx)
	return
}

//Renew the specified Domain
//prod: maximum: 10 minimum: 1, Number of years to extend the Domain. Must not exceed maximum for TLD.
// 		When omitted, defaults to period specified during original purchase
func (c *DomainClient) Renew(ctx context.Context, name string, prod int) (*OrderResult, error) {
	// /v1/domains/{domain}/renew
	var result OrderResult
	var input = map[string]int{"prod": prod}
	_, err := c.post("/" + name + "/renew").Input(input).Output(result).Run(ctx)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//VerifyRegistrantEmail Re-send Contact E-mail Verification for specified Domain
func (c *DomainClient) VerifyRegistrantEmail(ctx context.Context, name string) error {
	// /v1/domains/{domain}/verifyRegistrantEmail
	_, err := c.post("/" + name + "/verifyRegistrantEmail").Run(ctx)
	return err
}
