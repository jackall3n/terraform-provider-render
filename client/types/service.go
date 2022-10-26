package types

type ServiceItem struct {
	Service Service `json:"service"`
	Cursor  string  `json:"cursor"`
}

type ServiceDeploy struct {
	DeployId string  `json:"deployId"`
	Service  Service `json:"service"`
}

type Service struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Type           string         `json:"type"`
	Repo           string         `json:"repo"`
	OwnerId        string         `json:"ownerId"`
	ServiceDetails ServiceDetails `json:"serviceDetails"`
}

type EnvVar struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type EnvVarItem struct {
	EnvVar EnvVar `json:"envVar"`
	Cursor string `json:"cursor"`
}

type ServiceDetails struct {
	Env                string             `json:"env"`
	Region             string             `json:"region"`
	BuildCommand       string             `json:"buildCommand"`
	EnvSpecificDetails EnvSpecificDetails `json:"envSpecificDetails"`
}

type EnvSpecificDetails struct {
	BuildCommand string `json:"buildCommand"`
	StartCommand string `json:"startCommand"`
}
