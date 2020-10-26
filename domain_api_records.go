package godaddy

import (
	"context"
	"log"
)

//v1/domains/{domain}/records/{type}/{name}
func (c *DomainClient) getRecordUrlPath(domain, recordType, recordName string) string {
	action := "/" + domain + "/records"
	if recordType != "" {
		action += "/" + recordType
		if recordName != "" {
			action += "/" + recordName
		}
	}
	return action
}

//ReplaceRecords Replace all DNS Records for the specified Domain
func (c *DomainClient) ReplaceRecords(ctx context.Context, domain, recordType, recordName string, records ...DNSRecord) error {
	_, err := c.put(c.getRecordUrlPath(domain, recordType, recordName)).Input(records).Run(ctx)
	//log.Println("ReplaceRecords:\n", string(v))
	return err
}

//DeleteRecords delete DNS Records for the specified Domain
func (c *DomainClient) DeleteRecords(ctx context.Context, domain, recordType, recordName string) error {
	_, err := c.delete(c.getRecordUrlPath(domain, recordType, recordName)).Run(ctx)
	return err
}

//AddRecords Add the specified DNS Records to the specified Domain
func (c *DomainClient) AddRecords(ctx context.Context, domain, recordType, recordName string, records ...DNSRecord) error {
	_, err := c.patch(c.getRecordUrlPath(domain, recordType, recordName)).Input(records).Run(ctx)
	return err
}

//GetRecords Retrieve DNS Records for the specified Domain, optionally with the specified Type and/or Name
func (c *DomainClient) GetRecords(ctx context.Context, domain, recordType, recordName string) (result []DNSRecord, err error) {
	var v []byte
	v, err = c.get(c.getRecordUrlPath(domain, recordType, recordName)).Output(&result).Run(ctx)
	log.Println(string(v))
	return
}

//SetRecords add or replace all DNS Records for the specified Domain
func (c *DomainClient) SetRecords(ctx context.Context, domain, recordType, recordName string, records ...DNSRecord) error {
	_, err := c.put(c.getRecordUrlPath(domain, recordType, recordName)).Input(records).Run(ctx)
	//log.Println("ReplaceRecords:\n", string(v))
	return err
}
