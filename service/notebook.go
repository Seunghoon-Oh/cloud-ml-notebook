package service

import "github.com/Seunghoon-Oh/cloud-ml-notebook-manager/data"

func GetNotebooks() []string {
	// notebooks := "notebook-1"
	return data.GetNotebooksData()
}

func CreateNotebook() string {

	// TODO: Create Kubernetes Object
	return data.CreateNotebookData()
}
