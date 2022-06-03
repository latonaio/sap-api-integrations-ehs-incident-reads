package responses

type Incident struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
