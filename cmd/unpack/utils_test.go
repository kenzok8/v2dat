package unpack

import (
	"bytes"
	"testing"

	"github.com/kenzok8/v2dat/v2data"
)

func TestSplitAttrs(t *testing.T) {
	tag, attrs := splitAttrs("google@ads@tracking")
	if tag != "google" {
		t.Fatalf("tag = %q, want google", tag)
	}
	if len(attrs) != 2 {
		t.Fatalf("attrs = %v, want 2 entries", attrs)
	}
	for _, want := range []string{"ads", "tracking"} {
		if _, ok := attrs[want]; !ok {
			t.Fatalf("attrs missing %q: %v", want, attrs)
		}
	}

	tag, attrs = splitAttrs("google")
	if tag != "google" || attrs != nil {
		t.Fatalf("splitAttrs(google) = (%q, %v), want (google, nil)", tag, attrs)
	}
}

func TestFilterAttrs(t *testing.T) {
	in := []*v2data.Domain{
		{Value: "a.com", Attribute: []*v2data.Domain_Attribute{{Key: "ads"}}},
		{Value: "b.com", Attribute: []*v2data.Domain_Attribute{{Key: "tracking"}}},
		{Value: "c.com"},
	}
	out := filterAttrs(in, map[string]struct{}{"ads": {}})
	if len(out) != 1 || out[0].Value != "a.com" {
		t.Fatalf("filterAttrs(ads) = %v, want [a.com]", out)
	}
	if got := filterAttrs(in, nil); len(got) != 3 {
		t.Fatalf("filterAttrs(nil) = %d entries, want 3", len(got))
	}
}

func TestFileName(t *testing.T) {
	for in, want := range map[string]string{
		"/tmp/geosite.dat": "geosite",
		"geoip.dat":        "geoip",
		"nodot":            "nodot",
	} {
		if got := fileName(in); got != want {
			t.Fatalf("fileName(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestConvertV2DomainToText(t *testing.T) {
	domains := []*v2data.Domain{
		{Type: v2data.Domain_Domain, Value: "google.com"},
		{Type: v2data.Domain_Plain, Value: "google"},
		{Type: v2data.Domain_Regex, Value: `www\.google\.com$`},
		{Type: v2data.Domain_Full, Value: "www.google.com"},
	}
	var buf bytes.Buffer
	if err := convertV2DomainToText(domains, &buf); err != nil {
		t.Fatal(err)
	}
	want := "google.com\nkeyword:google\nregexp:www\\.google\\.com$\nfull:www.google.com\n"
	if buf.String() != want {
		t.Fatalf("got %q, want %q", buf.String(), want)
	}

	if err := convertV2DomainToText([]*v2data.Domain{{Type: 99, Value: "x"}}, &bytes.Buffer{}); err == nil {
		t.Fatal("invalid domain type should error")
	}
}

func TestConvertV2CidrToText(t *testing.T) {
	cidrs := []*v2data.CIDR{
		{Ip: []byte{192, 168, 0, 0}, Prefix: 16},
		{Ip: []byte{0x20, 0x01, 0x0d, 0xb8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Prefix: 32},
	}
	var buf bytes.Buffer
	if err := convertV2CidrToText(cidrs, &buf); err != nil {
		t.Fatal(err)
	}
	want := "192.168.0.0/16\n2001:db8::/32\n"
	if buf.String() != want {
		t.Fatalf("got %q, want %q", buf.String(), want)
	}

	if err := convertV2CidrToText([]*v2data.CIDR{{Ip: []byte{1, 2, 3}, Prefix: 24}}, &bytes.Buffer{}); err == nil {
		t.Fatal("invalid ip should error")
	}
}
