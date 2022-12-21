package ado

import "time"

type BuildLease struct {
	LeaseId         int       `json:"leaseId"`
	RunId           int       `json:"runId"`
	DefinitionId    int       `json:"definitionId"`
	OwnerId         string    `json:"ownerId"`
	CreatedOn       time.Time `json:"createdOn"`
	ValidUntil      time.Time `json:"validUntil"`
	ProtectPipeline bool      `json:"protectPipeline"`
}

type BuildLeaseResponse struct {
	Count  int          `json:"count"`
	Leases []BuildLease `json:"value"`
}

type Repository struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type Build struct {
	Id          int        `json:"id"`
	Status      string     `json:"status"`
	BuildNumber string     `json:"buildNumber"`
	Repository  Repository `json:"repository"`
	QueueTime   time.Time  `json:"queueTime"`
}

type BuildResponse struct {
	Count  int     `json:"count"`
	Builds []Build `json:"value"`
}

type BuildDefinition struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Uri         string `json:"uri"`
	Path        string `json:"path"`
	Type        string `json:"type"`
	QueueStatus string `json:"queueStatus"`
}

type BuildDefinitionResponse struct {
	Count       int               `json:"count"`
	Definitions []BuildDefinition `json:"value"`
}
