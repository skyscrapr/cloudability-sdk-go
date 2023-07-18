package cloudability

import (
	"encoding/json"
)

const viewsEndpoint = "/views/"

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
	Field      string `json:"field"`
	Comparator string `json:"comparator"`
	Value      string `json:"value"`
}

// View - Cloudabiity View
type View struct {
	ID                     string        `json:"id"`
	Title                  string        `json:"title"`
	SharedWithUsers        []string      `json:"sharedWithUsers"`
	SharedWithOrganization bool          `json:"sharedWithOrganization"`
	OwnerID                string        `json:"ownerId"`
	Filters                []*ViewFilter `json:"filters"`
}

// GetViews - returns all views
func (e ViewsEndpoint) GetViews() ([]View, error) {
	var result v3Result[[]View]
	err := e.get(e, "", &result)
	return result.Result, err
}

// GetView - return a single view
func (e ViewsEndpoint) GetView(id string) (*View, error) {
	var result v3Result[*View]
	err := e.get(e, id, &result)
	return result.Result, err
}

type viewPayload struct {
	Title                  string       `json:"title"`
	SharedWithUsers        []string     `json:"sharedWithUsers"`
	SharedWithOrganization bool         `json:"sharedWithOrganization"`
	Filters                []ViewFilter `json:"filters"`
}

// NewView - create a new view
func (e *ViewsEndpoint) NewView(view *View) (*View, error) {
	viewPayload := new(viewPayload)
	jsonView, _ := json.Marshal(view)
	json.Unmarshal(jsonView, viewPayload)
	var result v3Result[*View]
	err := e.post(e, "", viewPayload, &result)
	return result.Result, err
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
