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

export function isThrl6pValidateNameRequest(arg: any): arg is models.Thrl6pValidateNameRequest {
  return (
  arg != null &&
  typeof arg === 'object' &&
    // name?: string
    ( typeof arg.name === 'undefined' || typeof arg.name === 'string' ) &&

  true
  );
  }


