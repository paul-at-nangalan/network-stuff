package transport

import "net/http"

type WithCustomHeaders struct{
	rt http.RoundTripper
	headers http.Header
}

func NewWithCustomHeaders(rt http.RoundTripper)WithCustomHeaders{
	return WithCustomHeaders{
		rt: rt,
		headers: make(http.Header),
	}
}

func (p WithCustomHeaders) RoundTrip(request *http.Request) (*http.Response, error) {
	for key, vals := range p.headers{
		for _, val := range vals {
			request.Header.Add(key, val)
		}
	}
	return p.rt.RoundTrip(request)
}

func (p WithCustomHeaders) SetHeader(key string, value string){
	header, found := p.headers[key]
	if !found{
		p.headers[key] = make([]string, 0)
		header = p.headers[key]
	}
	header = append(header, value)
	p.headers[key] = header
}