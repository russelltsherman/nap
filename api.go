package nap

import "fmt"

// API structure
type API struct {
	BaseURL       string // https://httpbin.org
	Resources     map[string]*RestResource
	DefaultRouter *CBRouter
	Client        *Client
}

// NewAPI return new API structure
func NewAPI(baseURL string) *API {
	return &API{
		BaseURL:       baseURL,
		Resources:     make(map[string]*RestResource),
		DefaultRouter: NewRouter(),
		Client:        NewClient(),
	}
}

// SetAuth add authentication to api client
func (api *API) SetAuth(auth Authentication) {
	api.Client.SetAuth(auth)
}

// AddResource to list of api resources
func (api *API) AddResource(name string, res *RestResource) {
	api.Resources[name] = res
}

// Call a resource
func (api *API) Call(name string, params map[string]string) error {
	res, ok := api.Resources[name]
	if !ok {
		return fmt.Errorf("Resource does not exist: %s", name)
	}
	if err := api.Client.ProcessRequest(api.BaseURL, res, params); err != nil {
		return err
	}
	return nil
}

// ResourceNames show lit of resources
func (api *API) ResourceNames() []string {
	resources := []string{}
	for k := range api.Resources {
		resources = append(resources, k)
	}
	return resources
}
