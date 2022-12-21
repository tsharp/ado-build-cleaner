package ado

import (
	"encoding/json"
	"fmt"
	"log"
)

type BuildFolder struct {
	Path string `json:"path"`
}

type BuildFolderResponse struct {
	Count   int           `json:"count"`
	Folders []BuildFolder `json:"value"`
}

func DeleteEmptyBuildFolders(orgInfo OrganizationInfo) {
	for {
		fmt.Println("Checking for Empty Folders ...")
		folderDeleted := false
		allBuildFolders := GetBuildFolders(orgInfo, "")

		for _, folder := range allBuildFolders {
			if IsBuildFolderEmpty(orgInfo, folder.Path) {
				ExecutePreviewRequest(orgInfo, "DELETE", "build/folders", fmt.Sprintf("path=%s", folder.Path), nil)
				fmt.Println("Deleted Folder ==> " + folder.Path)
				folderDeleted = true
			}
		}

		if !folderDeleted {
			break
		}
	}
}

func GetBuildFolders(orgInfo OrganizationInfo, path string) []BuildFolder {
	resp := ExecutePreviewRequest(orgInfo, "GET", fmt.Sprintf("build/folders/%s", path), "", nil)

	decoder := json.NewDecoder(resp.Body)
	val := &BuildFolderResponse{}
	err := decoder.Decode(val)

	if err != nil {
		log.Fatal(err)
	}

	buildFolders := []BuildFolder{}

	// Don't include the current folder in the response set ...
	for _, folder := range val.Folders {
		if folder.Path != path {
			buildFolders = append(buildFolders, folder)
		} else {
			fmt.Println(folder.Path)
		}
	}

	return buildFolders
}

func IsBuildFolderEmpty(orgInfo OrganizationInfo, path string) bool {
	buildDefinitions := GetBuildsInPath(orgInfo, path)
	buildFolders := GetBuildFolders(orgInfo, path)

	fmt.Printf("[%s] Sub Folders: %d, Build Definitions: %d\n\r", path, len(buildFolders), len(buildDefinitions))

	return len(buildFolders) == 0 && len(buildDefinitions) == 0
}
