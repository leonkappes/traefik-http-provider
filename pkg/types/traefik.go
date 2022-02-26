package types

type Http struct {
	Routers     map[string]Router     `json:"routers"`
	Middlewares map[string]Middleware `json:"middlewares"`
	Services    map[string]Service    `json:"services"`
}

type Router struct {
	Rule       string   `json:"rule"`
	Service    string   `json:"service"`
	Midlewares []string `json:"middlewares,omitempty"`
}

type Middleware struct {
	Headers   map[string]interface{} `json:"headers,omitempty"`
	BasicAuth *BasicAuth             `json:"basicAuth,omitempty"`
}

type BasicAuth struct {
	Users []string `json:"users,omitempty"`
}

type Service struct {
	LoadBalancer LoadBalancer `json:"loadBalancer"`
}

type LoadBalancer struct {
	Servers []Server `json:"servers"`
}

type Server struct {
	Url string `json:"url"`
}
