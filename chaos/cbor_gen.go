// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package chaos

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

var lengthBufState = []byte{129}

func (t *State) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufState); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Unmarshallable ([]*chaos.UnmarshallableCBOR) (slice)
	if len(t.Unmarshallable) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Unmarshallable was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Unmarshallable))); err != nil {
		return err
	}
	for _, v := range t.Unmarshallable {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *State) UnmarshalCBOR(r io.Reader) error {
	*t = State{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Unmarshallable ([]*chaos.UnmarshallableCBOR) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Unmarshallable: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Unmarshallable = make([]*UnmarshallableCBOR, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v UnmarshallableCBOR
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Unmarshallable[i] = &v
	}

	return nil
}

var lengthBufCreateActorArgs = []byte{132}

func (t *CreateActorArgs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufCreateActorArgs); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.UndefActorCID (bool) (bool)
	if err := cbg.WriteBool(w, t.UndefActorCID); err != nil {
		return err
	}

	// t.ActorCID (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.ActorCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.ActorCID: %w", err)
	}

	// t.UndefAddress (bool) (bool)
	if err := cbg.WriteBool(w, t.UndefAddress); err != nil {
		return err
	}

	// t.Address (address.Address) (struct)
	if err := t.Address.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *CreateActorArgs) UnmarshalCBOR(r io.Reader) error {
	*t = CreateActorArgs{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.UndefActorCID (bool) (bool)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.UndefActorCID = false
	case 21:
		t.UndefActorCID = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	// t.ActorCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.ActorCID: %w", err)
		}

		t.ActorCID = c

	}
	// t.UndefAddress (bool) (bool)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.UndefAddress = false
	case 21:
		t.UndefAddress = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	// t.Address (address.Address) (struct)

	{

		if err := t.Address.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Address: %w", err)
		}

	}
	return nil
}