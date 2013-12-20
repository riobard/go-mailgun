package mailgun

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Route struct {
    Id string           `json:"id"`
    Priority int        `json:"priority"`
    Description string  `json:"description"`
	Expression string   `json:"expression"`
	Actions []string    `json:"actions"`
}

func (mg *Mailgun) Routes(limit, skip int) (total int, res []Route, err error) {
    v := url.Values{}
	v.Set("limit", strconv.Itoa(limit))
	v.Set("skip", strconv.Itoa(skip))
	body, err := mg.api("GET", "/routes", v)
	if err != nil {
		return
	}

	var j struct {
		Total int     `json:"total_count"`
		Items []Route `json:"items"`
	}

	err = json.Unmarshal(body, &j)
	if err != nil {
		return
	}
	total, res = j.Total, j.Items
	return

}

func (mg *Mailgun) Get(routeId string) (r Route, err error) {
    rsp, err := mg.api("GET", "/routes/"+routeId, nil)
    if err != nil {
		return
	}
	var res struct {
		Message string `json:"message"`
	    R	    Route `json:"route"`
	}
	err = json.Unmarshal(rsp, &res)
	r = res.R

	return
}

func (mg *Mailgun) Create(r *Route) (routeId string, err error) {
	v := url.Values{}
	
    v.Set("priority", strconv.Itoa(r.Priority))
    v.Set("description", r.Description)
    v.Set("expression", r.Expression)
	
    for _, a := range r.Actions {
		v.Add("action", a)
	}

    rsp, err := mg.api("POST", "/routes", v)
    if err != nil {
		return
	}
	var res struct {
		Message string `json:"message"`
	    R	    Route `json:"route"`
	}
	err = json.Unmarshal(rsp, &res)
	routeId = res.R.Id
	return
}

func (mg *Mailgun) Update(r *Route) (routeId string, err error) {
	v := url.Values{}
	
    v.Set("priority", strconv.Itoa(r.Priority))
    v.Set("description", r.Description)
    v.Set("expression", r.Expression)
	
    for _, a := range r.Actions {
		v.Add("action", a)
	}

    rsp, err := mg.api("PUT", "/routes/"+r.Id, v)
    if err != nil {
		return
	}
	var res struct {
		Message string `json:"message"`
	    R	    Route `json:"route"`
	}
	err = json.Unmarshal(rsp, &res)
	routeId = res.R.Id

	return
}

func (mg *Mailgun) Delete(r *Route) (err error) {
    rsp, err := mg.api("DELETE", "/routes/"+r.Id, nil)
    if err != nil {
		return
	}
	var res struct {
		Message string `json:"message"`
	}
	err = json.Unmarshal(rsp, &res)
	return
}
