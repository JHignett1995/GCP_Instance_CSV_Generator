# GCP Instance CSV Generator

## Requirements
1. Create 2 local folders: "config" & "ServiceAccounts".
2. Create a config file in the config folder:
`{   "projects": [        
   {"project_id": "My-First-Project"},  
   {"project_id": "My-Second-Project"}
  ]    
}`
3.Add all service account keys to the ServiecAccount folder.
   Ensure these are named the same as the Project ID. e.g. My-First-Project.json
## How to

Simply run using go run main.go
Build for intended OS systems using the appropriate go command lines.