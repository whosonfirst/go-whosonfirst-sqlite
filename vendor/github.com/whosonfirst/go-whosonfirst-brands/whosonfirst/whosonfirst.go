package whosonfirst

import (
	"github.com/whosonfirst/go-whosonfirst-brands"
	"github.com/whosonfirst/go-whosonfirst-flags"
	"github.com/whosonfirst/go-whosonfirst-flags/existential"
	"github.com/whosonfirst/go-whosonfirst-json"
	"github.com/whosonfirst/go-whosonfirst-json/properties"
	"io"
	_ "log"
	"os"
	"path/filepath"
)

func LoadWOFBrandFromFile(path string) (brands.Brand, error) {

	abs_path, err := filepath.Abs(path)

	if err != nil {
		return nil, err
	}

	fh, err := os.Open(abs_path)

	if err != nil {
		return nil, err
	}

	defer fh.Close()

	return LoadWOFBrandFromReader(fh)
}

func LoadWOFBrandFromReader(fh io.ReadCloser) (brands.Brand, error) {

	body, err := json.UnmarshalDocumentFromReader(fh)

	if err != nil {
		return nil, err
	}

	// check properties here...

	br := WOFBrand{
		body: body,
	}

	return &br, nil
}

type WOFBrand struct {
	json.Document
	brands.Brand
	body []byte
}

func (b *WOFBrand) Bytes() []byte {
	return b.body
}

func (b *WOFBrand) String() string {
	return string(b.Bytes())
}

func (b *WOFBrand) Id() int64 {
	return properties.Int64Property(b, []string{"wof:brand_id"}, -1)
}

func (b *WOFBrand) Name() string {
	return properties.StringProperty(b, []string{"wof:brand_name"}, "")
}

func (b *WOFBrand) Size() string {
	return properties.StringProperty(b, []string{"wof:brand_size"}, "")
}

func (b *WOFBrand) LastModified() int64 {
	return properties.Int64Property(b, []string{"wof:lastmodified"}, -1)
}

func (b *WOFBrand) IsCurrent() (flags.ExistentialFlag, error) {

	c := properties.Int64Property(b, []string{"mz:is_current"}, -1)

	if c == 0 || c == 1 {
		return existential.NewKnownUnknownFlag(c)
	}

	var fl flags.ExistentialFlag
	var err error

	fl, err = b.IsSuperseded()

	if err != nil {
		return nil, err
	}

	if fl.IsTrue() && fl.IsKnown() {
		return existential.NewKnownUnknownFlag(0)
	}

	fl, err = b.IsCeased()

	if err != nil {
		return nil, err
	}

	if fl.IsTrue() && fl.IsKnown() {
		return existential.NewKnownUnknownFlag(0)
	}

	fl, err = b.IsDeprecated()

	if err != nil {
		return nil, err
	}

	if fl.IsTrue() && fl.IsKnown() {
		return existential.NewKnownUnknownFlag(0)
	}

	// check something in order to return something with "1" with
	// some amount of confidence

	return existential.NewKnownUnknownFlag(-1)
}

func (b *WOFBrand) IsCeased() (flags.ExistentialFlag, error) {

	c := properties.StringProperty(b, []string{"edtf:cessation"}, "")

	var fl int64

	switch c {
	case "":
		fl = 0
	case "uuuu":
		fl = -1
	default:
		fl = 1
	}

	return existential.NewKnownUnknownFlag(fl)
}

func (b *WOFBrand) IsDeprecated() (flags.ExistentialFlag, error) {

	d := properties.StringProperty(b, []string{"edtf:deprecated"}, "")

	var fl int64

	switch d {
	case "":
		fl = 0
	case "uuuu":
		fl = -1
	default:
		fl = 1
	}

	return existential.NewKnownUnknownFlag(fl)
}

func (b *WOFBrand) IsSuperseding() (flags.ExistentialFlag, error) {

	supersedes := b.Supersedes()

	if len(supersedes) > 0 {
		return existential.NewKnownUnknownFlag(1)
	}

	return existential.NewKnownUnknownFlag(0)
}

func (b *WOFBrand) IsSuperseded() (flags.ExistentialFlag, error) {

	superseded_by := b.SupersededBy()

	if len(superseded_by) > 0 {
		return existential.NewKnownUnknownFlag(1)
	}

	return existential.NewKnownUnknownFlag(0)
}

func (b *WOFBrand) SupersededBy() []int64 {

	return properties.Int64PropertyArray(b, []string{"wof:superseded_by"})
}

func (b *WOFBrand) Supersedes() []int64 {

	return properties.Int64PropertyArray(b, []string{"wof:supersedes"})
}
