export interface THRL6P {
  data: Data;
  schema: string;
  version: number;
}

export interface Data {
  device: number;
  device_version: number;
  tone: Tone;
  meta: Meta;
}

export interface Meta {
  name: string;
  tnid: number;
}

export interface Tone {
  THRGroupAmp: THRGroupAmp;
  THRGroupCab: THRGroupCab;
  THRGroupFX1Compressor: THRGroupFX1Compressor;
  THRGroupFX3EffectEcho: THRGroupFX3EffectEcho;
  THRGroupFX4EffectReverb: THRGroupFX4EffectReverb;
  global: Global;
}

export interface THRGroupAmp {
  '@asset': string;
  Bass: number;
  Drive: number;
  Master: number;
  Mid: number;
  Treble: number;
}

export interface THRGroupCab {
  '@asset': string;
  '@enabled': boolean;
  Level: string;
  Sustain: string;
}

export interface THRGroupFX1Compressor {
  '@asset': string;
  '@enabled': boolean;
}

export interface THRGroupFX3EffectEcho {
  '@asset': string;
  '@enabled': boolean;
}

export interface THRGroupFX4EffectReverb {
  '@asset': string;
  '@enabled': boolean;
}

export interface THRGroupGate {
  '@asset': string;
  '@enabled': boolean;
}

export interface Global {
  THRPresetParamTempo: number;
}
