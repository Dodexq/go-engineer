package gometr

type GoMetrClient struct {
	url     string
	timeout int32
}

func (g *GoMetrClient) NewGometr(url string, timeout int32) GoMetrClient {
	g.url = url
	g.timeout = timeout
	return *g
}
