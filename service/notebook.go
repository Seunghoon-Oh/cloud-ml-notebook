package service

import "github.com/Seunghoon-Oh/cloud-ml-notebook-manager/data"

func GetNotebooks() []string {
	return data.GetRedisData()
}

func CreateNotebook() string {
	// TODO: Create Kubernetes Object
	return data.CreateRedisData()
}
