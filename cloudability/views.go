package cloudability

import (
	"encoding/json"
)

const viewsEndpoint = "/v3/views/"

// ViewsEndpoint - Cloudability Views Endpoint
type ViewsEndpoint struct {
	*v3Endpoint
}

// Views endpoint
func (c *Client) Views() *ViewsEndpoint {
	return &ViewsEndpoint{newV3Endpoint(c, viewsEndpoint)}
}

// ViewFilter - Cloudability ViewFilter
type ViewFilter struct {
	Field string `json:"field"`
	Comparator string `json:"comparator"`
	Value string `json:"value"`
}

// View - Cloudabiity View
type View struct {
	ID string `json:"id"`
	Title string `json:"title"`
	SharedWithUsers []string `json:"sharedWithUsers"`
	SharedWithOrganization bool `json:"sharedWithOrganization"`
	OwnerID string `json:"ownerId"`
	Filters []*ViewFilter `json:"filters"`
}

// GetViews - returns all views
func (e ViewsEndpoint) GetViews() ([]View, error) {
	var views []View
	err := e.get(e, "", &views)
	return views, err
}

// GetView - return a single view
func (e ViewsEndpoint) GetView(id string) (*View, error) {
	var view View
	err := e.get(e, id, &view)
	return &view, err
}

type viewPayload struct {
	Title string `json:"title"`
	SharedWithUsers []string `json:"sharedWithUsers"`
	SharedWithOrganization bool `json:"sharedWithOrganization"`
	Filters []ViewFilter `json:"filters"`
}

// NewView - create a new view
func (e *ViewsEndpoint) NewView(view *View) (*View, error) {
	viewPayload := new(viewPayload)
	jsonView, _ := json.Marshal(view)
	json.Unmarshal(jsonView, viewPayload)
	var newView View
	err := e.post(e, "", viewPayload, &newView)
	return &newView, err
}

// UpdateView - update a view
func (e *ViewsEndpoint) UpdateView(view *View) error {
	viewPayload := new(viewPayload)
	jsonView, _ := json.Marshal(view)
    json.Unmarshal(jsonView, viewPayload)
	return e.put(e, view.ID, viewPayload)
}

// DeleteView - delete a view
func (e *ViewsEndpoint) DeleteView(id string) error {
	return e.delete(e, id)
}
