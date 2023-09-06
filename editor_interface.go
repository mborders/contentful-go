package contentful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// EditorInterfacesService service
type EditorInterfacesService service

// EditorInterface model
type EditorInterface struct {
	Sys      *Sys       `json:"sys"`
	Controls []Controls `json:"controls"`
	SideBar  []Sidebar  `json:"sidebar"`
}

// Controls model
type Controls struct {
	FieldID         string            `json:"fieldId"`
	WidgetNameSpace string            `json:"widgetNamespace"`
	WidgetID        string            `json:"widgetId"`
	Settings        map[string]string `json:"settings,omitempty"`
}

// Sidebar model
type Sidebar struct {
	WidgetNameSpace string            `json:"widgetNamespace"`
	WidgetID        string            `json:"widgetId"`
	Settings        map[string]string `json:"settings,omitempty"`
	Disabled        bool              `json:"disabled"`
}

// List returns an EditorInterface collection
func (service *EditorInterfacesService) List(spaceID string) *Collection {
	path := fmt.Sprintf("/spaces/%s/environments/%s/editor_interface", spaceID, service.c.Environment)

	req, err := service.c.newRequest("GET", path, nil, nil)
	if err != nil {
		return &Collection{}
	}

	col := NewCollection(&CollectionOptions{})
	col.c = service.c
	col.req = req

	return col
}

// Get returns a single EditorInterface
func (service *EditorInterfacesService) Get(spaceID, contentTypeID string) (*EditorInterface, error) {
	path := fmt.Sprintf("/spaces/%s/environments/%s/content_types/%s/editor_interface", spaceID, service.c.Environment, contentTypeID)
	query := url.Values{}
	method := "GET"

	req, err := service.c.newRequest(method, path, query, nil)
	if err != nil {
		return &EditorInterface{}, err
	}

	var editorInterface EditorInterface
	if ok := service.c.do(req, &editorInterface); ok != nil {
		return nil, err
	}

	return &editorInterface, err
}

// Update updates an editor interface
func (service *EditorInterfacesService) Update(spaceID, contentTypeID string, e *EditorInterface) error {
	bytesArray, err := json.Marshal(e)
	if err != nil {
		return err
	}

	var path string
	var method string

	if contentTypeID != "" {
		path = fmt.Sprintf("/spaces/%s/environments/%s/content_types/%s/editor_interface", spaceID, service.c.Environment, contentTypeID)
		method = "PUT"
	}

	req, err := service.c.newRequest(method, path, nil, bytes.NewReader(bytesArray))
	if err != nil {
		return err
	}

	return service.c.do(req, e)
}
