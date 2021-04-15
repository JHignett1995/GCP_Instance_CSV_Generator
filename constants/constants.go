package constants

import "os"

func HandleFile(filename string) *os.File {
	file, err3 := os.OpenFile("./"+filename+".csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ErrorCheck("Could Not Create File: ", err3)
	return file
}

func GetZones() []string {
	zonesArray := []string{"us-central1-a",
		"australia-southeast1-c",
		"europe-north1-c",
		"us-west1-a",
		"europe-west3-a",
		"asia-northeast2-b",
		"asia-east1-a",
		"europe-west6-b",
		"us-west2-a",
		"us-west1-b",
		"europe-west4-b",
		"europe-west1-b",
		"europe-west3-b",
		"europe-west1-d",
		"europe-west6-c",
		"us-east4-a",
		"us-central1-f",
		"australia-southeast1-b",
		"southamerica-east1-b",
		"asia-northeast2-a",
		"northamerica-northeast1-b",
		"us-west2-c",
		"us-east4-b",
		"us-central1-b",
		"southamerica-east1-c",
		"southamerica-east1-a",
		"asia-east2-c",
		"europe-north1-a",
		"us-east1-c",
		"us-central1-c",
		"asia-south1-c",
		"asia-east2-a",
		"northamerica-northeast1-a",
		"us-east1-b",
		"us-west1-c",
		"europe-west2-c",
		"asia-northeast1-c",
		"asia-southeast1-c",
		"europe-west1-c",
		"europe-west3-c",
		"northamerica-northeast1-c",
		"us-west2-b",
		"europe-west2-a",
		"asia-southeast1-a",
		"australia-southeast1-a",
		"europe-west4-c",
		"europe-west2-b",
		"europe-west6-a",
		"us-east1-d",
		"asia-east1-b",
		"asia-south1-b",
		"asia-south1-a",
		"us-east4-c",
		"europe-west4-a",
		"asia-southeast1-b",
		"europe-north1-b",
		"asia-northeast2-c",
		"asia-east1-c",
		"asia-northeast1-b",
		"asia-northeast1-a",
		"asia-east2-b"}
	return zonesArray
}
