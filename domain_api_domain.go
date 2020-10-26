package godaddy

import (
	"context"
)

type GetDomainRequest struct {
	Statuses     []string `json:"statuses,omitempty"`
	StatusGroups []string `json:"statusGroups,omitempty"`
	Limit        int      `json:"limit,omitempty"`
	Marker       string   `json:"marker,omitempty"`
	Includes     []string `json:"includes,omitempty"`
	ModifiedDate string   `json:"modifiedDate,omitempty"`
}

//GetDomains Retrieve a list of Domains for the specified Shopper
func (c *DomainClient) GetDomains(ctx context.Context, query GetDomainRequest) (result []*DomainSummary, err error) {
	_, err = c.get("", ParseStruct(query)).Output(&result).Run(ctx)
	return
}

//GetDomain Retrieve details for the specified Domain
func (c *DomainClient) GetDomain(ctx context.Context, domain string) (*DomainSummary, error) {
	var result DomainSummary
	_, err := c.get("/" + domain).Output(&result).Run(ctx)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
