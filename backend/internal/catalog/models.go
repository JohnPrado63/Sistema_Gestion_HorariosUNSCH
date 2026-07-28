package catalog

type Facultad struct {
	ID     int    `json:"id_facultad"`
	Nombre string `json:"nombre"`
}

type Departamento struct {
	ID         int    `json:"id_departamento"`
	IDFacultad int    `json:"id_facultad"`
	Nombre     string `json:"nombre"`
}

type Escuela struct {
	ID             int    `json:"id_escuela"`
	IDFacultad     int    `json:"id_facultad"`
	IDDepartamento int    `json:"id_departamento"`
	Nombre         string `json:"nombre"`
}

type Aula struct {
	ID           int    `json:"id_aula"`
	IDPabellon   int    `json:"id_pabellon"`
	IDEscuela    *int   `json:"id_escuela"`
	Codigo       string `json:"codigo"`
	Tipo         string `json:"tipo"`
	Aforo        int    `json:"aforo"`
	EsCompartida bool   `json:"es_compartida"`
}

type Usuario struct {
	ID     int    `json:"id_usuario"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	Rol    string `json:"rol"`
}
