/* tslint:disable */

import { HttpClient } from '@angular/common/http';
import { Inject, Injectable, Optional } from '@angular/core';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';
import { DefaultHttpOptions, HttpOptions } from './';
import { USE_DOMAIN, USE_HTTP_OPTIONS, PatchServiceAPIClient } from './patch-service-api-client.service';

import * as models from '../../models';
import * as guards from '../../guards';

/**
 * Created with https://github.com/flowup/api-client-generator
 */
@Injectable()
export class GuardedPatchServiceAPIClient extends PatchServiceAPIClient {

  constructor(readonly httpClient: HttpClient,
              @Optional() @Inject(USE_DOMAIN) domain?: string,
              @Optional() @Inject(USE_HTTP_OPTIONS) options?: DefaultHttpOptions) {
    super(httpClient, domain, options);
  }

  listPatch(
    args: {
      filter?: string,
      pageSize?: number,
      page?: number,
    },
    requestHttpOptions?: HttpOptions
  ): Observable<models.Thrl6pPatchList> {
    return super.listPatch(args, requestHttpOptions)
      .pipe(tap((res: any) => guards.isThrl6pPatchList(res) || console.error(`TypeGuard for response 'Thrl6pPatchList' caught inconsistency.`, res)));
  }

  createPatch(
    args: {
      body: models.Thrl6pCreatePatchRequest,
    },
    requestHttpOptions?: HttpOptions
  ): Observable<models.Thrl6pPatch> {
    return super.createPatch(args, requestHttpOptions)
      .pipe(tap((res: any) => guards.isThrl6pPatch(res) || console.error(`TypeGuard for response 'Thrl6pPatch' caught inconsistency.`, res)));
  }

  metadata(
    requestHttpOptions?: HttpOptions
  ): Observable<models.Thrl6pMeta> {
    return super.metadata(requestHttpOptions)
      .pipe(tap((res: any) => guards.isThrl6pMeta(res) || console.error(`TypeGuard for response 'Thrl6pMeta' caught inconsistency.`, res)));
  }

  validateName(
    args: {
      body: models.Thrl6pValidateNameRequest,
    },
    requestHttpOptions?: HttpOptions
  ): Observable<models.Thrl6pNameValidation> {
    return super.validateName(args, requestHttpOptions)
      .pipe(tap((res: any) => guards.isThrl6pNameValidation(res) || console.error(`TypeGuard for response 'Thrl6pNameValidation' caught inconsistency.`, res)));
  }

  getPatch(
    args: {
      id: string,
    },
    requestHttpOptions?: HttpOptions
  ): Observable<models.Thrl6pPatch> {
    return super.getPatch(args, requestHttpOptions)
      .pipe(tap((res: any) => guards.isThrl6pPatch(res) || console.error(`TypeGuard for response 'Thrl6pPatch' caught inconsistency.`, res)));
  }

}
