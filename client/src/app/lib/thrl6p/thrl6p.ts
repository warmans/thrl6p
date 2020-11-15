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
  THRGroupFX2Effect: THRGroupFX2Effect;
  THRGroupFX3EffectEcho: THRGroupFX3EffectEcho;
  THRGroupFX4EffectReverb: THRGroupFX4EffectReverb;
  THRGroupGate: THRGroupGate;
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
  SpkSimType: string;
}

export interface THRGroupFX1Compressor {
  '@asset': string;
  '@enabled': boolean;
  Level: number;
  Sustain: number;
}

export interface THRGroupFX2Effect {
  '@asset': string;
  '@enabled': boolean;
  '@wetDry': number;
  Depth: number;
  Feedback: number;
  Freq: number;
  Pre: number;
}

export interface THRGroupFX3EffectEcho {
  '@asset': string;
  '@enabled': boolean;
  '@wetDry': number;
  'Bass': number;
  'Feedback': number;
  'Time': number;
  'Treble': number;
}

export interface THRGroupFX4EffectReverb {
  '@asset': string;
  '@enabled': boolean;
  '@wetDry': number;
  'Decay': number;
  'PreDelay': number;
  'Tone': number;
}

export interface THRGroupGate {
  '@asset': string;
  '@enabled': boolean;
  Decay: number;
  Thresh: number;
}

export interface Global {
  THRPresetParamTempo: number;
}
