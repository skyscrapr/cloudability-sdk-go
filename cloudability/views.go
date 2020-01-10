package cloudability

import (
	"strconv"
	"encoding/json"
)


type viewsEndpoint struct {
	*cloudabilityV3Endpoint
}

func newViewsEndpoint(apikey string) *viewsEndpoint {
	e := &viewsEndpoint{newCloudabilityV3Endpoint(apikey)}
	e.EndpointPath = "/v3/views/"
	return e
}

type ViewFilter struct {
	Field string `json:"field"`
	Comparator string `json:"comparator"`
	Value string `json:"value"`
}

type View struct {
	Id int `json:"id"`
	Title string `json:"title"`
	SharedWithUsers []string `json:"sharedWithUsers"`
	SharedWithOrganization bool `json:"sharedWithOrganization"`
	OwnerId string `json:"ownerId"`
	Filters []ViewFilter `json:"filters"`
}

func (e viewsEndpoint) GetViews() ([]View, error) {
	var views []View
	err := e.get(e, "", &views)
	return views, err
}

func (e viewsEndpoint) GetView(id int) (*View, error) {
	var view View
	err := e.get(e, strconv.Itoa(id), &view)
	return &view, err
}

type viewPayload struct {
	Title string `json:"title"`
	SharedWithUsers []string `json:"sharedWithUsers"`
	SharedWithOrganization bool `json:"sharedWithOrganization"`
	Filters []ViewFilter `json:"filters"`
}

func (e *viewsEndpoint) NewView(view *View) (*View, error) {
	viewPayload := new(viewPayload)
	jsonView, _ := json.Marshal(view)
	json.Unmarshal(jsonView, viewPayload)
	var newView View
	err := e.post(e, "", viewPayload, &newView)
	return &newView, err
}

func (e *viewsEndpoint) UpdateView(view *View) error {
	viewPayload := new(viewPayload)
	jsonView, _ := json.Marshal(view)
    json.Unmarshal(jsonView, viewPayload)
	return e.put(e, strconv.Itoa(view.Id), viewPayload)
}

func (e *viewsEndpoint) DeleteView(id int) error {
	return e.delete(e, strconv.Itoa(id))
}
