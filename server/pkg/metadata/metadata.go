package metadata

import (
	"encoding/json"
	v1 "github.com/warmans/thrl6p/server/gen/api/v1"
)

//go:generate go-bindata -modtime 1 -mode 420 -pkg metadata -o ./metadata_gen.go ../../etc/thr-metadata.json

var metadata *Meta

func init() {
	metadata = &Meta{}
	if err := json.Unmarshal(MustAsset("../../etc/thr-metadata.json"), metadata); err != nil {
		panic(err)
	}
}

func All() *Meta {
	return metadata
}

type Amp string

func (a Amp) Proto() v1.AmpMeta_Amp {
	switch a {
	case AmpClassic:
		return v1.AmpMeta_CLASSIC
	case AmpBoutique:
		return v1.AmpMeta_BOUTIQUE
	case AmpModern:
		return v1.AmpMeta_MODERN
	}
	return v1.AmpMeta_UNKNOWN_AMP_GROUP
}

const (
	AmpUnknown  Amp = ""
	AmpClassic  Amp = "classic"
	AmpBoutique Amp = "boutique"
	AmpModern   Amp = "modern"
)

type Channel string

func (c Channel) Proto() v1.ChannelMeta_Channel {
	switch c {
	case ChannelClean:
		return v1.ChannelMeta_CLEAN
	case ChannelCrunch:
		return v1.ChannelMeta_CRUNCH
	case ChannelLead:
		return v1.ChannelMeta_LEAD
	case ChannelHiGain:
		return v1.ChannelMeta_HI_GAIN
	case ChannelSpecial:
		return v1.ChannelMeta_SPECIAL
	case ChannelBass:
		return v1.ChannelMeta_BASS
	case ChannelAcoustic:
		return v1.ChannelMeta_ACOUSTIC
	case ChannelFlat:
		return v1.ChannelMeta_FLAT
	}
	return v1.ChannelMeta_UNKNOWN_CHANNEL
}

const (
	ChannelUnknown  Channel = ""
	ChannelClean    Channel = "clean"
	ChannelCrunch   Channel = "crunch"
	ChannelLead     Channel = "lead"
	ChannelHiGain   Channel = "hi_gain"
	ChannelSpecial  Channel = "special"
	ChannelBass     Channel = "bass"
	ChannelAcoustic Channel = "acoustic"
	ChannelFlat     Channel = "flat"
)

type Meta struct {
	Amps []*AmpMeta `json:"amps"`
}

func (m *Meta) Proto() *v1.Meta {
	p := &v1.Meta{
		Amps: make([]*v1.AmpMeta, len(m.Amps)),
	}
	for k, v := range m.Amps {
		p.Amps[k] = v.Proto()
	}
	return p
}

type AmpMeta struct {
	Amp      Amp            `json:"amp"`
	Channels []*ChannelMeta `json:"channels"`
}

func (m *AmpMeta) Proto() *v1.AmpMeta {
	p := &v1.AmpMeta{
		Amp: m.Amp.Proto(),
		Channels: make([]*v1.ChannelMeta, len(m.Channels)),
	}
	for k, v := range m.Channels {
		p.Channels[k] = v.Proto()
	}
	return p
}

type ChannelMeta struct {
	Channel     Channel `json:"channel"`
	Name        string  `json:"name"`
	InspiredBy  string  `json:"inspired_by"`
	Description string  `json:"description"`
}

func (m *ChannelMeta) Proto() *v1.ChannelMeta {
	return &v1.ChannelMeta{
		Channel:     m.Channel.Proto(),
		Name:        m.Name,
		InspiredBy:  m.InspiredBy,
		Description: m.Description,
	}
}
