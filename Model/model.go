package model

type (
	// TUser -- t_users
	TUser struct {
		ID            uint `json:"ID,omitempty"`
		Fullname      string
		Gender        uint8
		Dob           string
		Email         string `json:"Email,omitempty"`
		Phone         string `json:"Phone,omitempty"`
		Propinsi      string
		Kota          string
		Kecamatan     string
		Kelurahan     string
		IDWilayah     string
		Address       string
		AlamatDefault string
		ImageURL      string
		Ektp          string
		Status        uint8
		Points        uint32
		Balance       uint64
		Timezone      uint8
		Created       string `json:"Created,omitempty"`
		Updated       string `json:"Updated,omitempty"`
		Npwp          string
	}
)
