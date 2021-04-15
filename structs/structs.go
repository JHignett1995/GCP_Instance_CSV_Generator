package structs

import (
	"strconv"
	"strings"
)

type DiskData struct {
	CreationTimestamp      string   `json:"creationTimestamp"`
	ID                     string   `json:"id"`
	Kind                   string   `json:"kind"`
	LabelFingerprint       string   `json:"labelFingerprint"`
	LastAttachTimestamp    string   `json:"lastAttachTimestamp"`
	LicenseCodes           []string `json:"licenseCodes"`
	Licenses               []string `json:"licenses"`
	Name                   string   `json:"name"`
	PhysicalBlockSizeBytes string   `json:"physicalBlockSizeBytes"`
	SelfLink               string   `json:"selfLink"`
	SizeGb                 string   `json:"sizeGb"`
	SourceImage            string   `json:"sourceImage"`
	SourceImageID          string   `json:"sourceImageId"`
	Status                 string   `json:"status"`
	Type                   string   `json:"type"`
	Users                  []string `json:"users"`
	Zone                   string   `json:"zone"`
}


type InstanceData struct {
	CPUPlatform        string              `json:"cpuPlatform"`
	CreationTimestamp  string              `json:"creationTimestamp"`
	DeletionProtection bool                `json:"deletionProtection"`
	Disks              []Disks             `json:"disks"`
	DisplayDevice      DisplayDevice       `json:"displayDevice"`
	ID                 string              `json:"id"`
	Kind               string              `json:"kind"`
	LabelFingerprint   string              `json:"labelFingerprint"`
	Labels             Labels              `json:"labels"`
	MachineType        string              `json:"machineType"`
	Metadata           Metadata            `json:"metadata"`
	Name               string              `json:"name"`
	NetworkInterfaces  []NetworkInterfaces `json:"networkInterfaces"`
	Scheduling         Scheduling          `json:"scheduling"`
	SelfLink           string              `json:"selfLink"`
	ServiceAccounts    []ServiceAccounts   `json:"serviceAccounts"`
	Status             string              `json:"status"`
	Tags               Tags                `json:"tags"`
	Zone               string              `json:"zone"`
}


type Disks struct {
	AutoDelete bool     `json:"autoDelete"`
	Boot       bool     `json:"boot"`
	DeviceName string   `json:"deviceName"`
	Interface  string   `json:"interface"`
	Kind       string   `json:"kind"`
	Licenses   []string `json:"licenses"`
	Mode       string   `json:"mode"`
	Source     string   `json:"source"`
	Type       string   `json:"type"`
}

type Labels struct {
	HaGroup string `json:"ha_group"`
}

type Items struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DisplayDevice struct {
}

type Metadata struct {
	Fingerprint string  `json:"fingerprint"`
	Items       []Items `json:"items"`
	Kind        string  `json:"kind"`
}

type AccessConfigs struct {
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	NatIP       string `json:"natIP"`
	NetworkTier string `json:"networkTier"`
	Type        string `json:"type"`
}

type NetworkInterfaces struct {
	AccessConfigs []AccessConfigs `json:"accessConfigs"`
	Fingerprint   string          `json:"fingerprint"`
	Kind          string          `json:"kind"`
	Name          string          `json:"name"`
	Network       string          `json:"network"`
	NetworkIP     string          `json:"networkIP"`
	Subnetwork    string          `json:"subnetwork"`
}

type Scheduling struct {
	AutomaticRestart  bool   `json:"automaticRestart"`
	OnHostMaintenance string `json:"onHostMaintenance"`
}

type ServiceAccounts struct {
	Email  string   `json:"email"`
	Scopes []string `json:"scopes"`
}

type Tags struct {
	Fingerprint string   `json:"fingerprint"`
	Items       []string `json:"items"`
}

func HandleDisks(disks []Disks) []string{
	var diskArray []string
	for _ , disk := range disks{
		//AutoDelete := strconv.FormatBool(disk.AutoDelete)
		//Boot := strconv.FormatBool(disk.Boot)
		DeviceName := disk.DeviceName
		//Interface := disk.Interface
		//Kind := disk.Kind
		var Licenses string
		if len(disk.Licenses) >0{
			ops := strings.Split(disk.Licenses[len(disk.Licenses)-1], "/")
			Licenses = ops[len(ops)-1]
		}else{
			Licenses = "No Data"
		}
		//Mode := disk.Mode
		//Source := disk.Source
		//Type := disk.Type
		diskArray = append(diskArray, DeviceName, Licenses)

//		diskArray = append(diskArray, AutoDelete, Boot, DeviceName, Interface, Kind, Licenses, Mode, Source, Type)
	}
	return diskArray
}
func HandleAccessConfigs(accessConfigs []AccessConfigs) []string{
	var accessConfigArray []string
	for _, accessConfig := range accessConfigs {
		Kind := accessConfig.Kind
		Name := accessConfig.Name
		NatIP := accessConfig.NatIP
		NetworkTier := accessConfig.NetworkTier
		Type := accessConfig.Type
		accessConfigArray = append(accessConfigArray,Kind, Name, NatIP, NetworkTier, Type)
	}
	return accessConfigArray
}
func HandleNetworkInterfaces(networkInterfaces []NetworkInterfaces) []string{
	var networkInterfaceArray []string
	for _, networkInterface := range networkInterfaces {
		AccessConfigs := strings.Join(HandleAccessConfigs(networkInterface.AccessConfigs), ", ")
		NetworkInterfaceFingerprint := networkInterface.Fingerprint
		NetworkInterfaceKind := networkInterface.Kind
		NetworkInterfaceName := networkInterface.Name
		Network := networkInterface.Network
		NetworkIP := networkInterface.NetworkIP
		Subnetwork := networkInterface.Subnetwork
		networkInterfaceArray = append(networkInterfaceArray, AccessConfigs, NetworkInterfaceFingerprint,NetworkInterfaceKind,NetworkInterfaceName,Network,NetworkIP,Subnetwork)
	}
	return networkInterfaceArray
}
func HandleServiceAccounts(serviceAccounts []ServiceAccounts) []string {
	var serviceAccountArray []string
	for _, serviceAccount := range serviceAccounts {
		Email := serviceAccount.Email
		Scopes := strings.Join(serviceAccount.Scopes, ", ")
		serviceAccountArray = append(serviceAccountArray, Email, Scopes)
	}
	return serviceAccountArray
}
func HandleMetaData(metadata Metadata) []string {
	var metadataArray []string
	Fingerprint := metadata.Fingerprint
	MetaDataKind := metadata.Kind
	metadataArray = append(metadataArray, Fingerprint, MetaDataKind)
	return metadataArray
}
func HandleLabels(labels Labels) []string {
	var labelArray []string
	HaGroup := labels.HaGroup
	labelArray = append(labelArray,HaGroup)
	return labelArray
}
func HandleScheduling(schedules Scheduling) []string {
	var scheduleArray []string
	AutomaticRestart := strconv.FormatBool(schedules.AutomaticRestart)
	OnHostMaintenance := schedules.OnHostMaintenance
	scheduleArray = append(scheduleArray,AutomaticRestart,OnHostMaintenance)
	return scheduleArray
}
func HandleTags(tags Tags) []string {
	var tagArray []string
	TagFingerprint := tags.Fingerprint
	Items := strings.Join(tags.Items, ", ")
	tagArray = append(tagArray, TagFingerprint, Items)
	return tagArray
}