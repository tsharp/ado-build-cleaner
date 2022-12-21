package ado

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func GetBuildsInPath(orgInfo OrganizationInfo, path string) []BuildDefinition {
	resp := ExecuteRequest(orgInfo, "GET", "build/definitions", fmt.Sprintf("path=%s", path), nil)

	decoder := json.NewDecoder(resp.Body)
	val := &BuildDefinitionResponse{}
	err := decoder.Decode(val)

	if err != nil {
		log.Fatal(err)
	}

	return val.Definitions
}

func GetAllBuildLeases(orgInfo OrganizationInfo, definitionId int) (bool, []int, error) {
	var cutoffDate = time.Date(2022, time.November, 1, 0, 0, 0, 0, time.UTC)

	resp := ExecuteRequest(orgInfo, "GET", "build/retention/leases", fmt.Sprintf("definitionId=%d", definitionId), nil)

	if resp.StatusCode >= 400 {
		log.Fatalf("Response Code: %d\n\r", resp.StatusCode)
	}

	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	decoder := json.NewDecoder(resp.Body)

	val := &BuildLeaseResponse{}
	jsonErr := decoder.Decode(val)

	if jsonErr != nil {
		log.Fatal("Failed to Parse Result")
	}

	hasRecentLeases := false
	leaseIds := []int{}

	for _, s := range val.Leases {
		// repoName := s.Repository.Name
		if s.CreatedOn.Before(cutoffDate) {
			leaseIds = append(leaseIds, s.LeaseId)
		} else {
			// Warn that lease cannot be removed!
			hasRecentLeases = true
		}
	}

	return hasRecentLeases, leaseIds, nil
}

func GetBuildDefinitionIds(orgInfo OrganizationInfo) []int {
	resp := ExecuteRequest(orgInfo, "GET", "build/definitions", "", nil)

	decoder := json.NewDecoder(resp.Body)
	val := &BuildDefinitionResponse{}
	jsonErr := decoder.Decode(val)

	if jsonErr != nil {
		log.Fatal("Failed to Parse Result", jsonErr)
	}

	definitionIds := []int{}
	// Subsets is a slice so you must loop over it
	for _, s := range val.Definitions {
		if strings.Contains(strings.ToLower(s.Name), "(deprecated)") ||
			strings.Contains(strings.ToLower(s.Name), "(draycarys)") {
			fmt.Printf("[%d] %s\n\r", s.Id, s.Name)
			definitionIds = append(definitionIds, s.Id)
		}
	}

	return definitionIds
}

func DeleteBuildLeases(orgInfo OrganizationInfo, definitionId int, leaseIds []int) {
	if len(leaseIds) == 0 {
		fmt.Println("No Leases to delete for this definition")
		return
	}

	var leaseIdStr []string
	for _, i := range leaseIds {
		leaseIdStr = append(leaseIdStr, strconv.Itoa(i))
	}

	ExecuteRequest(orgInfo, "DELETE", "build/retention/leases", "ids="+strings.Join(leaseIdStr, ","), nil)
}

func DeleteBuildDefinition(orgInfo OrganizationInfo, definitionId int) {
	// fmt.Printf("Deleting Build Definition! %d\n\r", definitionId)
	ExecuteRequest(orgInfo, "DELETE", fmt.Sprintf("build/definitions/%d", definitionId), "", nil)
}
