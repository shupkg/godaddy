package godaddy

import (
    "context"
)

//AddRecords Add the specified DNS Records to the specified Domain
func (c *DomainClient) AddRecords(ctx context.Context, domain string, records ...DNSRecord) error {
    _, err := c.patch(RecordSearch{}.Path(domain)).Input(records).Run(ctx)
    return err
}

//GetRecords Retrieve DNS Records for the specified Domain, optionally with the specified Type and/or Name
func (c *DomainClient) GetRecords(ctx context.Context, domain string, search RecordSearch) (result []DNSRecord, err error) {
    _, err = c.get(search.Path(domain)).Output(&result).Run(ctx)
    return
}

//SetRecords add or replace all DNS Records for the specified Domain
func (c *DomainClient) SetRecords(ctx context.Context, domain string, search RecordSearch, records ...DNSRecord) error {
    _, err := c.put(search.Path(domain)).Input(records).Run(ctx)
    return err
}

//DeleteRecords delete DNS Records for the specified Domain
func (c *DomainClient) DeleteRecords(ctx context.Context, domain string, search RecordSearch) error {
    _, err := c.delete(search.Path(domain)).Run(ctx)
    return err
}

func Search(search ...string) (rs RecordSearch) {
    if len(search) > 0 {
        rs.Type = search[1]
    }
    if len(search) > 1 {
        rs.Name = search[2]
    }
    return
}

type RecordSearch struct {
    Type, Name string
}

//Path v1/domains/{domain}/records/{type}/{name}
func (rs RecordSearch) Path(domain string) string {
    action := "/" + domain + "/records"
    if rs.Type != "" {
        action += "/" + rs.Type
        if rs.Name != "" {
            action += "/" + rs.Name
        }
    }
    return action
}
