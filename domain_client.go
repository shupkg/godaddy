package godaddy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	headerAccept        = "Accept"
	headerAuthorization = "Authorization"
	headerMarketID      = "X-Market-Id"
	headerShopperID     = "X-Shopper-Id"
	headerContent       = "Content-Type"
	mediaTypeJSON       = "application/json"
)

const (
	urlDomains = "/v1/domains"
)

func New() *DomainClient {
	return &DomainClient{
		baseUrl: "https://api.godaddy.com" + urlDomains,
		httpClient: &http.Client{
			Timeout:   time.Second * 30,
			Transport: newTransport(time.Second),
		},
	}
}

type DomainClient struct {
	httpClient *http.Client
	baseUrl    string

	marketId  string
	shopperId string
	key       string
	secret    string
}

func (c *DomainClient) Ote() *DomainClient {
	c.baseUrl = "https://api.ote-godaddy.com" + urlDomains
	return c
}

func (c *DomainClient) IsValid() error {
	if c.key == "" {
		return errors.New("miss godaddy sso key")
	}

	if c.secret == "" {
		return errors.New("miss godaddy sso secret")
	}

	return nil
}

func (c *DomainClient) Secret(key, secret string) *DomainClient {
	c.key = key
	c.secret = secret
	return c
}

func (c *DomainClient) SetSecret(secret string) *DomainClient {
	c.secret = secret
	return c
}

func (c *DomainClient) SetKey(key string) *DomainClient {
	c.key = key
	return c
}

func (c *DomainClient) SetShopperId(shopperId string) *DomainClient {
	c.shopperId = shopperId
	return c
}

func (c *DomainClient) SetMarketId(marketId string) *DomainClient {
	c.marketId = marketId
	return c
}

func (c *DomainClient) newRequest(method, url string, query ...Query) *Client {
	req := Client{}.Init()
	req.HttpClient(c.httpClient)
	req.HandleError(c.handleError)
	req.HeaderSet(headerAccept, mediaTypeJSON)
	req.HeaderSet(headerContent, mediaTypeJSON)
	req.HeaderSet(headerAuthorization, fmt.Sprintf("sso-key %s:%s", c.key, c.secret))

	if c.shopperId != "" {
		req.HeaderSet(headerShopperID, c.shopperId)
	}
	if c.marketId != "" {
		req.HeaderSet(headerMarketID, c.marketId)
	}
	u := c.baseUrl + url
	if len(query) > 0 {
		for i := range query {
			if !strings.ContainsRune(u, '?') {
				u += "?"
			}
			u += query[i].Encode()
		}
	}
	req.Method(method, u)
	return req
}

func (c *DomainClient) get(url string, query ...Query) *Client {
	return c.newRequest(http.MethodGet, url, query...)
}

func (c *DomainClient) post(url string, query ...Query) *Client {
	return c.newRequest(http.MethodPost, url, query...)
}

func (c *DomainClient) put(url string, query ...Query) *Client {
	return c.newRequest(http.MethodPut, url, query...)
}

func (c *DomainClient) patch(url string, query ...Query) *Client {
	return c.newRequest(http.MethodPatch, url, query...)
}

func (c *DomainClient) delete(url string, query ...Query) *Client {
	return c.newRequest(http.MethodDelete, url, query...)
}

func (c *DomainClient) handleError(statusCode int, body []byte) error {
	var errResp = struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Fields  []struct {
			Code        string `json:"code"`
			Message     string `json:"message"`
			Path        string `json:"path"`
			PathRelated string `json:"pathRelated"`
		} `json:"fields"`
	}{}

	if err := json.Unmarshal(body, &errResp); err != nil {
		return err
	}

	if len(errResp.Fields) == 0 {
		return fmt.Errorf("[%d:%s] %s", statusCode, errResp.Code, errResp.Message)
	}

	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("[%d:%s] %s (", statusCode, errResp.Code, errResp.Message))
	for i, field := range errResp.Fields {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("%s [%s]: %s", field.Path, field.Code, field.Message))
	}
	b.WriteString(")")
	return fmt.Errorf("%s", b.String())
}
