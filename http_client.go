package toolkits

import (
	"bytes"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"net/url"
    "time"
    //"io/ioutil"
)

type HttpClient struct {
	Client *http.Client
}



func NewHttpClient(timeout,keepalive int) *HttpClient {
    if timeout > keepalive {
        return nil
    }
	return &HttpClient{Client: &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:  time.Duration(timeout) * time.Second ,
				KeepAlive: time.Duration(keepalive) * time.Second,
			}).DialContext,
		},
		Timeout: time.Duration(timeout) * time.Second ,
	},
	}
}

//func (c *HttpClient) DoReq(method, u string, body interface{},header map[string]string,queryparams map[string]string) (string, error) {
func (c *HttpClient) DoReq(method, u string, body interface{},header map[string]string,queryparams map[string]string) (*http.Response, error) {
	var (
		req *http.Request
		err error
	)

	req,err = c.NewReqByMethod(method, u,body,queryparams)
	if err != nil {
		return nil, err
	}

	if header != nil && len(header) > 0{
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil,err
	}

	//defer resp.Body.Close()
    if resp.StatusCode != 200 {
        return nil,errors.New("response status code is not 200")
    }
	//b, err := ioutil.ReadAll(resp.Body)
    //if err != nil {
    //    return "",err
    //}
	//return string(b), nil
	return resp, nil
}
func (c *HttpClient) NewReqByMethod(method, u string, body interface{},queryparams map[string]string) (*http.Request,error) {
	var (
		err error
		req *http.Request = &http.Request{}
		b []byte = make([]byte,0)
	)
	if body != nil {
		b, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
    url := CutLastestSlash(u)
	if method == "GET" || method == "DELETE" {
		req, err = http.NewRequest(method, url, nil)
        if err != nil {
            return nil,err
        }
		_,err = c.SetQuery(u,queryparams)
		if err != nil {
			return nil,err
		}
	} else if  method == "POST" || method == "PUT" {
		req, err = http.NewRequest(method, url, bytes.NewReader(b))
        if err != nil {
            return nil,err
        }
	} else {
		return nil,errors.New("request method invalid")
	}
	return req,nil
}
//set query
func (c *HttpClient) SetQuery(u string, params map[string]string) (string, error) {
	var q url.Values
	_u, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	if params != nil {
		q = _u.Query()
		for k, v := range params {
			q.Set(k, v)
		}
	}
	_u.RawQuery = q.Encode()

	return _u.String(), nil
}

//去掉url配置中最后一个‘/’
func CutLastestSlash(u string) string {
	if u[len(u)-1] == '/' {
		u = u[0:(len(u) - 1)]
	}

	return u
}
