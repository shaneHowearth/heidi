package profile

type Updated struct {
	PatientID  int    `json:"patientID"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	IsPregnant bool   `json:"isPregnant"`
	UpdatedAt  string `json:"updatedAt"` // RFC3339Nano
}
