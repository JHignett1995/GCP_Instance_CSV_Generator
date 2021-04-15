package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"gcp_instances/constants"
	"gcp_instances/structs"
	"google.golang.org/api/compute/v1"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct{
	Projects []struct{
		ProjectId string `json:"project_id"`
	}`json:"projects"`
}

func main() {
	var c Config
	ctx := context.Background()
	configFile, err := ioutil.ReadFile("./config/config.json")
	constants.ErrorCheck("Reading Config", err)
	err = json.Unmarshal(configFile, &c)
	constants.ErrorCheck("UnMarshalling Config", err)

	headers := []string{"Name", "Disks", "Machine Type", "Project"}

	file := constants.HandleFile("Instances")
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(headers)
	constants.ErrorCheck("Write Headers To CSV: ", err)

	for _, v := range c.Projects {
		constants.ErrorCheck("Starting: "+v.ProjectId, nil)
		err = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./ServiceAccounts/"+v.ProjectId+".json")
		constants.ErrorCheck("Setting Credentials", err)

		service, err := compute.NewService(ctx)
		constants.ErrorCheck("Creating Service", err)

		for _, zone := range constants.GetZones() {
			var instanceData structs.InstanceData
			req := service.Instances.List(v.ProjectId, zone)
			err := req.Pages(ctx, func(page *compute.InstanceList) error {
				instances := page.Items
				for _, instance := range instances{// j := 0; j < len(instances); j++ {

					// Marshal From GCP Object To JSON
					pageJSON, err := instance.MarshalJSON()
					constants.ErrorCheck("Marshal Page Into JSON", err)

					// Pre-Filter JSON To Remove Superfluous Strings
					formattedJSON := strings.ReplaceAll(string(pageJSON), "https://www.googleapis.com/compute/v1/projects/", "")

					// Unmarshal From JSON To Custom Struct
					err = json.Unmarshal([]byte(formattedJSON), &instanceData)
					constants.ErrorCheck("UnMarshal JSON", err)

					// Struct Variables Mapped To Local For Filtering Before Use

					InstanceName := instanceData.Name
					Disks := strings.Join(structs.HandleDisks(instanceData.Disks), ", ")
					MachineType := strings.ReplaceAll(instanceData.MachineType, v.ProjectId+"/zones/"+zone+"/machineTypes/", "")
					//CPUPlatform := instanceData.CPUPlatform

					err = writer.Write([]string{InstanceName,Disks, MachineType, v.ProjectId})
					constants.ErrorCheck("Write To CSV: ", err)
				}
				return nil
			})
			constants.ErrorCheck("Getting Instance Data", err)
		}
		ListDisks(ctx, service, v.ProjectId)
	}
}

func ListDisks(ctx context.Context, computeService *compute.Service, projectID string) {
	zonesArray := constants.GetZones()

	// Append To File If It Exists, If Not Create It.
	file := constants.HandleFile("disks")
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Create Initial Line, Which Is The Headers
	err := writer.Write([]string{"Name", "SizeGb", "Project ID"})
	constants.ErrorCheck("Write String To CSV: ", err)

	for _, zone := range zonesArray {
		var diskData structs.DiskData
		req := computeService.Disks.List(projectID, zone)
		err2 := req.Pages(ctx, func(page *compute.DiskList) error {
			disks := page.Items

			// All Assignment Should Happen Here As It Iterates Over Each Item In The List
			for j := 0; j < len(disks); j++ {

				// Marshal From GCP Object To JSON
				pageJSON, err := disks[j].MarshalJSON()
				constants.ErrorCheck("Marshal Page Into JSON: ", err)

				// Pre-Filter JSON To Remove Superfluous Strings
				formattedJSON := strings.ReplaceAll(string(pageJSON), "https://www.googleapis.com/compute/v1/projects/", "")

				// Unmarshal From JSON To Custom Struct
				err2 := json.Unmarshal([]byte(formattedJSON), &diskData)
				constants.ErrorCheck("UnMarshal JSON: ", err2)

				// Struct Variables Mapped To Local For Filtering Before Use
				Name := diskData.Name
				SizeGb := diskData.SizeGb

				// Append Line To CSV
				err3 := writer.Write([]string{Name, SizeGb, projectID})
				constants.ErrorCheck("Write To CSV: ", err3)
			}
			return nil
		})
		constants.ErrorCheck("Getting Networks list: ", err2)
	}
}
