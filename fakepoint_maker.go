package fakepoint

import (
	"net/http"
)

func NewFakepointMaker() *FakepointMaker {
	agent := NewAgent()
	fakeClient := &FakepointMaker{
		agent: agent,
	}

	agent.client = fakeClient
	fakeClient.Transport = agent
	return fakeClient
}

type FakepointMaker struct {
	http.Client
	agent *Agent
}

func (f FakepointMaker) client() *http.Client {
	return &http.Client{ Transport: f.agent }
}

func (f FakepointMaker) AddTrip(method, url string, statusCode int, document string) *Fakepoint {
	fr := &Fakepoint{
		method:     method,
		url:        url,
		document:   document,
		statusCode: statusCode,
		header:     &http.Header{},
		agent:      f.agent,
	}

	f.agent.add(url, method, *fr)
	return fr
}

func (f FakepointMaker) PlanGet(url string, statusCode int, document string) *Fakepoint {
	return f.AddTrip("GET", url, statusCode, document)
}

func (f FakepointMaker) PlanPost(url string, statusCode int, document string) *Fakepoint {
	return f.AddTrip("POST", url, statusCode, document)
}

func (f FakepointMaker) PlanPut(url string, statusCode int, document string) *Fakepoint {
	return f.AddTrip("PUT", url, statusCode, document)
}

func (f FakepointMaker) PlanDelete(url string, statusCode int, document string) *Fakepoint {
	return f.AddTrip("DELETE", url, statusCode, document)
}
