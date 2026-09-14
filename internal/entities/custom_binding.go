package entities

import "github.com/google/uuid"

type BindID struct {
	uuid.UUID
}

func (u *BindID) UnmarshalText(text []byte) error {
	parsed, err := uuid.Parse(string(text))
	if err != nil {
		return err
	}
	u.UUID = parsed
	return nil
}

func (u *BindID) UnmarshalParam(param string) error {
	id, err := uuid.Parse(param)
	if err != nil {
		return err
	}
	u.UUID = id
	return nil
}
