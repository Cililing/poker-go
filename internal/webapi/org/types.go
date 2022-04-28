package org

import "time"

type newOrgRequest struct {
	Name string `json:"name"`
}

type newOrgResponse struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
}

type addToOrgRequest struct {
	OrgName string `json:"name"`
	Who     string `json:"who"`
}

type listUserOrgRequest struct {
	// empty
}

type listUserOrgResponse struct {
	Orgs []orgResponse `json:"orgs"`
}

type idWithName struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type orgResponse struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	Admin     string       `json:"admin"`
	Members   []idWithName `json:"members"`
	CreatedAt time.Time    `json:"created_at"`
}
