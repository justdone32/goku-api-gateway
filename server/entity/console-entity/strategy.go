package entity

type Strategy struct {
	StrategyID     string `json:"strategyID"`
	StrategyName   string `json:"strategyName"`
	UpdateTime     string `json:"updateTime,omitempty"`
	CreateTime     string `json:"createTime,omitempty"`
	StrategyConfig string `json:"strategyConfig,omitempty"`
	GroupID        int    `json:"groupID"`
	GroupName      string `json:"groupName,omitempty"`
	EnableStatus   int    `json:"enableStatus"`
	StrategyType   int    `json:"strategyType"`
}
