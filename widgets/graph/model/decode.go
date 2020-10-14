package model

import (
	"errors"
	"github.com/jackc/pgtype"
)

func (t *Widget) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
	r := pgtype.Record{
		Fields: []pgtype.Value{
			&pgtype.Int4{},
			&pgtype.Varchar{},
			&pgtype.Int4{},
		},
	}

	if err := r.DecodeBinary(ci, src); err != nil {
		return err
	}

	if r.Status != pgtype.Present {
		return errors.New("BUG: decoding should not be called on NULL value")
	}

	id := r.Fields[0].(*pgtype.Int4)
	color := r.Fields[1].(*pgtype.Varchar)
	size := r.Fields[2].(*pgtype.Int4)

	if err := id.AssignTo(&t.ID); err != nil {
		return err
	}

	if err := color.AssignTo(&t.Color); err != nil {
		return err
	}

	if err := size.AssignTo(&t.Size); err != nil {
		return err
	}

	return nil
}
