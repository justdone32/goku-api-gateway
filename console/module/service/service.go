package service

import "encoding/json"

type Simple struct {
	Name        string `json:"name"`
	Driver      string `json:"driver"`
	DriverTitle string `json:"driverTitle"`
	Type        string `json:"type"`
}

type Service struct {
	Simple

	Desc        string `json:"desc"`
	IsDefault   bool   `json:"isDefault"`
	HealthCheck bool   `json:"healthCheck"`
	UpdateTime  string `json:"updateTime"`
	CreateTime  string `json:"createTime"`
}

type Info struct {
	*Service
	Config             string            `json:"config"`
	ClusterConfig      string            `json:"-"`
	ClusterConfigObj   map[string]string `json:"clusterConfig"`
	HealthCheckPath    string            `json:"healthCheckPath"`
	HealthCheckPeriod  int               `json:"healthCheckPeriod"`
	HealthCheckCode    string            `json:"healthCheckCode"`
	HealthCheckTimeOut int               `json:"healthCheckTimeOut"`
}

func (i *Info) Decode() {
	json.Unmarshal([]byte(i.ClusterConfig), &i.ClusterConfigObj)
}

type AddParam struct {
	Name          string `opt:"name,require"`
	Driver        string `opt:"driver" default:"static"`
	Desc          string `opt:"desc"`
	Config        string `opt:"config"`
	ClusterConfig string `opt:"clusterConfig"`
	//ClusterConfigObj map[string]string `json:"clusterConfig"`
	HealthCheck        bool   `opt:"healthCheck" default:"false"`
	HealthCheckPath    string `opt:"healthCheckPath"`
	HealthCheckPeriod  int    `opt:"healthCheckPeriod" default:"5" min:"1" max:"60"`
	HealthCheckCode    string `opt:"healthCheckCode" default:"200"`
	HealthCheckTimeOut int    `opt:"healthCheckTimeOut" default:"300" max:"5000" min:"0"`
}
