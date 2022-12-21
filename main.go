package main

import (
	"fmt"

	ado "github.com/tsharp/ado-build-cleaner/ado"
)

func main() {
	orgInfo := ado.OrganizationInfo{
		PersonalAccessToken: "<<a pat token with sufficient privileges>>",
		Organization:        "<<your org name here>>",
		Project:             "<<your project name here>>",
	}

	definitionIds := ado.GetBuildDefinitionIds(orgInfo)

	for _, definitionId := range definitionIds {
		hasRecentLeases, leaseIds, _ := ado.GetAllBuildLeases(orgInfo, definitionId)

		// Delete Leases
		ado.DeleteBuildLeases(orgInfo, definitionId, leaseIds)

		if hasRecentLeases {
			fmt.Printf("[%d] This definition can't be deleted as it has recent leases.\n\r", definitionId)
			continue
		}

		// Delete Build Definition
		ado.DeleteBuildDefinition(orgInfo, definitionId)
	}

	ado.DeleteEmptyBuildFolders(orgInfo)
}
