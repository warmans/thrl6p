/* tslint:disable */

import * as models from '../models';

/* pre-prepared guards for build in complex types */

function _isBlob(arg: any): arg is Blob {
  return arg != null && typeof arg.size === 'number' && typeof arg.type === 'string' && typeof arg.slice === 'function';
}

export function isFile(arg: any): arg is File {
return arg != null && typeof arg.lastModified === 'number' && typeof arg.name === 'string' && _isBlob(arg);
}

/* generated type guards */

export function isAmpMetaAmp(arg: any): arg is models.AmpMetaAmp {
  return false
   || arg === models.AmpMetaAmp.UNKNOWN_AMP_GROUP
   || arg === models.AmpMetaAmp.CLASSIC
   || arg === models.AmpMetaAmp.BOUTIQUE
   || arg === models.AmpMetaAmp.MODERN
  ;
  }

export function isChannelMetaChannel(arg: any): arg is models.ChannelMetaChannel {
  return false
   || arg === models.ChannelMetaChannel.UNKNOWN_CHANNEL
   || arg === models.ChannelMetaChannel.CLEAN
   || arg === models.ChannelMetaChannel.CRUNCH
   || arg === models.ChannelMetaChannel.LEAD
   || arg === models.ChannelMetaChannel.HI_GAIN
   || arg === models.ChannelMetaChannel.SPECIAL
   || arg === models.ChannelMetaChannel.BASS
   || arg === models.ChannelMetaChannel.ACOUSTIC
   || arg === models.ChannelMetaChannel.FLAT
  ;
  }

export function isThrl6pAmpMeta(arg: any): arg is models.Thrl6pAmpMeta {
  return (
  arg != null &&
  typeof arg === 'object' &&
    // amp?: AmpMetaAmp
    ( typeof arg.amp === 'undefined' || isAmpMetaAmp(arg.amp) ) &&
    // channels?: Thrl6pChannelMeta[]
    ( typeof arg.channels === 'undefined' || (Array.isArray(arg.channels) && arg.channels.every((item: unknown) => isThrl6pChannelMeta(item))) ) &&

  true
  );
  }

export function isThrl6pChannelMeta(arg: any): arg is models.Thrl6pChannelMeta {
  return (
  arg != null &&
  typeof arg === 'object' &&
    // channel?: ChannelMetaChannel
    ( typeof arg.channel === 'undefined' || isChannelMetaChannel(arg.channel) ) &&
    // description?: string
    ( typeof arg.description === 'undefined' || typeof arg.description === 'string' ) &&
    // inspired_by?: string
    ( typeof arg.inspired_by === 'undefined' || typeof arg.inspired_by === 'string' ) &&
    // name?: string
    ( typeof arg.name === 'undefined' || typeof arg.name === 'string' ) &&

  true
  );
  }

export function isThrl6pCreatePatchRequest(arg: any): arg is models.Thrl6pCreatePatchRequest {
  return (
  arg != null &&
  typeof arg === 'object' &&
    // description?: string
    ( typeof arg.description === 'undefined' || typeof arg.description === 'string' ) &&
    // patch?: string
    ( typeof arg.patch === 'undefined' || typeof arg.patch === 'string' ) &&

  true
  );
  }

export function isThrl6pMeta(arg: any): arg is models.Thrl6pMeta {
  return (
  arg != null &&
  typeof arg === 'object' &&
    // amps?: Thrl6pAmpMeta[]
    ( typeof arg.amps === 'undefined' || (Array.isArray(arg.amps) && arg.amps.every((item: unknown) => isThrl6pAmpMeta(item))) ) &&

  true
  );
  }

export function isThrl6pPatch(arg: any): arg is models.Thrl6pPatch {
  return (
  arg != null &&
  typeof arg === 'object' &&
    // description?: string
    ( typeof arg.description === 'undefined' || typeof arg.description === 'string' ) &&
    // id?: string
    ( typeof arg.id === 'undefined' || typeof arg.id === 'string' ) &&
    // name?: string
    ( typeof arg.name === 'undefined' || typeof arg.name === 'string' ) &&
    // patch?: string
    ( typeof arg.patch === 'undefined' || typeof arg.patch === 'string' ) &&
    // permalink?: string
    ( typeof arg.permalink === 'undefined' || typeof arg.permalink === 'string' ) &&

  true
  );
  }

export function isThrl6pPatchList(arg: any): arg is models.Thrl6pPatchList {
  return (
  arg != null &&
  typeof arg === 'object' &&
    // patches?: Thrl6pPatch[]
    ( typeof arg.patches === 'undefined' || (Array.isArray(arg.patches) && arg.patches.every((item: unknown) => isThrl6pPatch(item))) ) &&

  true
  );
  }

export function isThrl6pValidateNameRequest(arg: any): arg is models.Thrl6pValidateNameRequest {
  return (
  arg != null &&
  typeof arg === 'object' &&
    // name?: string
    ( typeof arg.name === 'undefined' || typeof arg.name === 'string' ) &&

  true
  );
  }


