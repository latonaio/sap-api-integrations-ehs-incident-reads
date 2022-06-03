package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-ehs-incident-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToIncident(raw []byte, l *logger.Logger) ([]Incident, error) {
	pm := &responses.Incident{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Incident. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	incident := make([]Incident, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		incident = append(incident, Incident{
			IncidentUUID:                data.IncidentUUID,
			IncidentCategory:            data.IncidentCategory,
			IncidentStatus:              data.IncidentStatus,
			IncidentTitle:               data.IncidentTitle,
			IncidentUTCDateTime:         data.IncidentUTCDateTime,
			IncidentLocationDescription: data.IncidentLocationDescription,
			EHSLocationUUID:             data.EHSLocationUUID,
			To_Attachments:              data.To_Attachments,
			To_Persons:                  data.To_Persons,
			To_Location:                 data.To_Location,
		})
	}

	return incident, nil
}
