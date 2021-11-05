package core

type User struct {
	email, idep, nomComplet, password, ip string
	groups                                []string
}

type OnyxiaUser struct {
	user     User
	projects []Project
}
