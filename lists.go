/*
Copyright (C) 2013 Brandon Philips <brandon@ifup.org>

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package mailgun

import (
	"encoding/json"
	"net/url"
)

type ListMemberResponse struct {
	Member  ListMember `json:member`
	Message string     `json:message`
}

type ListMember struct {
	Address     string            `json:address`
	Subscribed  bool              `json:subscribed`
	Vars        map[string]string `json:vars`
	Name        string            `json:name`
	Description string            `json:description`
}

func (m *ListMember) setURLValues(v *url.Values) {
	// Translate the subscribed field to a string
	v.Set("subscribed", "False")
	if m.Subscribed == true {
		v.Set("subscribed", "True")
	}
	v.Set("address", m.Address)
	v.Set("name", m.Name)
	v.Set("description", m.Description)
	vars, _ := json.Marshal(m.Vars)
	v.Set("vars", string(vars))
}

func (c *Client) AddListMember(list string, m ListMember) (message string, err error) {
	v := url.Values{}
	m.setURLValues(&v)

	rsp, err := c.api("POST", "/lists/"+list+"/members", v)
	if err != nil {
		return
	}

	response := ListMemberResponse{}
	err = json.Unmarshal(rsp, &response)
	message = response.Message
	return
}

func (c *Client) UpdateListMember(list string, m ListMember) (message string, err error) {
	v := url.Values{}
	m.setURLValues(&v)

	rsp, err := c.api("PUT", "/lists/"+list+"/members/"+m.Address, v)
	if err != nil {
		return
	}

	response := ListMemberResponse{}
	err = json.Unmarshal(rsp, &response)
	message = response.Message
	return
}

func (c *Client) GetListMember(list string, email string) (member ListMember, err error) {
	v := url.Values{}

	rsp, err := c.api("GET", "/lists/"+list+"/members/"+email, v)
	if err != nil {
		return
	}

	response := ListMemberResponse{}
	err = json.Unmarshal(rsp, &response)
	member = response.Member

	return
}
