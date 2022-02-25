package models

type RGroup struct {
	MessageId        string `validate:"required" json:"messageid"`
	ClientId         string `validate:"required" json:"clientid"`
	AuthorizationOam string `json:"authorizationoam"`
	Msisdn           string `validate:"required" json:"msisdn"`
	Status           string `validate:"required" json:"status"`
}
type Group struct {
	Membros []struct {
		Msisdn        string `validate:"required" json:"msisdn"`
		Status        string `validate:"required" json:"status"`
		Tipo          string `validate:"required" json:"tipo"`
		AtivoDesde    string `json:"ativo_desde"`
		ApelidoMembro string `validate:"required" json:"apelido_membro"`
		Docnumber     string `validate:"required" json:"docnumber"`
		ProfileName   string `validate:"required" json:"profile_name"`
		CodPlan       string `validate:"required" json:"cod_plan"`
	} `json:"membros"`
	Status  string `validate:"required" json:"status"`
	QtdDisp int    `validate:"required" json:"qtdDisp"`
}
