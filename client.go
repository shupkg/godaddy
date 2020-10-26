package godaddy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client       *http.Client
	errorHandler func(statusCode int, body []byte) error
	headerSet    map[string]string
	headerDel    map[string]interface{}
	method       string
	url          string

	input  interface{}
	output interface{}
}

type Unmarshaler interface {
	Unmarshal([]byte) error
}

type Marshaler interface {
	Marshal() (io.Reader, error)
}

func (c Client) Init() *Client {
	c.client = http.DefaultClient
	c.headerSet = map[string]string{}
	c.headerDel = map[string]interface{}{}
	c.method = http.MethodGet
	return &c
}

//HttpClient 设置头部
func (c *Client) HttpClient(client *http.Client) *Client {
	c.client = client
	return c
}

//HandleError 错误处理
func (c *Client) HandleError(handler func(statusCode int, body []byte) error) *Client {
	c.errorHandler = handler
	return c
}

//HeaderSet 设置头部
func (c *Client) HeaderSet(name, value string) *Client {
	delete(c.headerDel, name)
	c.headerSet[name] = value
	return c
}

//HeaderDel 删除头部
func (c *Client) HeaderDel(name string) *Client {
	delete(c.headerSet, name)
	c.headerDel[name] = struct{}{}
	return c
}

//Method 请求方法
func (c *Client) Method(method string, url ...string) *Client {
	if len(url) > 0 {
		c.Url(url[0])
	}
	c.method = method
	return c
}

//Get 请求方法
func (c *Client) Get(url ...string) *Client {
	return c.Method(http.MethodGet, url...)
}

//Post 请求方法
func (c *Client) Post(url ...string) *Client {
	return c.Method(http.MethodPost, url...)
}

//Url 请求地址
func (c *Client) Url(url string) *Client {
	c.url = url
	return c
}

//Output 输出对象，默认JSON格式
func (c *Client) Output(output interface{}) *Client {
	c.output = output
	return c
}

//Input 输入对象，默认JSON格式
func (c *Client) Input(input interface{}) *Client {
	c.input = input
	return c
}

//JSON 设置 Content-Type 为 application/json
func (c *Client) JSON() *Client {
	c.HeaderSet("Content-Type", "application/json")
	return c
}

//Run 执行
func (c *Client) Run(ctx context.Context) ([]byte, error) {
	input, err := c.marshalInput()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, c.method, c.url, input)
	if err != nil {
		return nil, err
	}

	if len(c.headerDel) > 0 {
		for key := range c.headerDel {
			req.Header.Del(key)
		}
	}

	if len(c.headerSet) > 0 {
		for key, value := range c.headerSet {
			req.Header.Set(key, value)
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return buf, c.unmarshalOutput(buf)
	}
	return buf, c.handlerError(resp.StatusCode, buf)
}

func (c *Client) marshalInput() (io.Reader, error) {
	if c.input != nil {
		if in, ok := c.input.(Marshaler); ok {
			return in.Marshal()
		}
		v, err := json.Marshal(c.input)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(v), nil
	}
	return nil, nil
}

func (c *Client) unmarshalOutput(data []byte) error {
	if c.output != nil {
		if out, ok := c.output.(Unmarshaler); ok {
			return out.Unmarshal(data)
		}
		return json.Unmarshal(data, c.output)
	}
	return nil
}

func (c *Client) handlerError(statusCode int, body []byte) error {
	if c.errorHandler != nil {
		return c.errorHandler(statusCode, body)
	}
	return fmt.Errorf("[%d] %s", statusCode, http.StatusText(statusCode))
}
