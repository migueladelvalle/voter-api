package process

type VoterDTO struct {
	id    int
	name  string
	email string
}

func NewVoterDTO(id int, name string, email string) VoterDTO {
	return VoterDTO{
		id:    id,
		name:  name,
		email: email,
	}
}

func (v *VoterDTO) GetId() int {
	return v.id
}

func (v *VoterDTO) GetName() string {
	return v.name
}

func (v *VoterDTO) GetEmail() string {
	return v.email
}
