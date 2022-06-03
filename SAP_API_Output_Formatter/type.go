package sap_api_output_formatter

type EhsIncident struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	EhsIncident   string `json:"ehs_incident"`
	Deleted       bool   `json:"deleted"`
}

type Incident struct {
	IncidentUUID                string `json:"IncidentUUID"`
	IncidentCategory            string `json:"IncidentCategory"`
	IncidentStatus              string `json:"IncidentStatus"`
	IncidentTitle               string `json:"IncidentTitle"`
	IncidentUTCDateTime         string `json:"IncidentUTCDateTime"`
	IncidentLocationDescription string `json:"IncidentLocationDescription"`
	EHSLocationUUID             string `json:"EHSLocationUUID"`
	To_Attachments              string `json:"to_Attachments"`
	To_Persons                  string `json:"to_Persons"`
	To_Location                 string `json:"to_Location"`
}
