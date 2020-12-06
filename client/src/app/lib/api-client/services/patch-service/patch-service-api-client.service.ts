/* tslint:disable */

import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Inject, Injectable, InjectionToken, Optional } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { DefaultHttpOptions, HttpOptions, PatchServiceAPIClientInterface } from './';

import * as models from '../../models';

export const USE_DOMAIN = new InjectionToken<string>('PatchServiceAPIClient_USE_DOMAIN');
export const USE_HTTP_OPTIONS = new InjectionToken<HttpOptions>('PatchServiceAPIClient_USE_HTTP_OPTIONS');

type APIHttpOptions = HttpOptions & {
  headers: HttpHeaders;
  params: HttpParams;
  responseType?: 'arraybuffer' | 'blob' | 'text' | 'json';
};

/**
 * Created with https://github.com/flowup/api-client-generator
 */
@Injectable()
export class PatchServiceAPIClient implements PatchServiceAPIClientInterface {

  readonly options: APIHttpOptions;

  readonly domain: string = `//${window.location.hostname}${window.location.port ? ':'+window.location.port : ''}`;

  constructor(private readonly http: HttpClient,
              @Optional() @Inject(USE_DOMAIN) domain?: string,
              @Optional() @Inject(USE_HTTP_OPTIONS) options?: DefaultHttpOptions) {

    if (domain != null) {
      this.domain = domain;
    }

    this.options = {
      headers: new HttpHeaders(options && options.headers ? options.headers : {}),
      params: new HttpParams(options && options.params ? options.params : {}),
      ...(options && options.reportProgress ? { reportProgress: options.reportProgress } : {}),
      ...(options && options.withCredentials ? { withCredentials: options.withCredentials } : {})
    };
  }

  /**
   * Lists patches.
   * Response generated for [ 200 ] HTTP response code.
   */
  listPatch(
    args: {
      filter?: string,
      pageSize?: number,
      page?: number,
    },
    requestHttpOptions?: HttpOptions
  ): Observable<models.Thrl6pPatchList> {
    const path = `/api/patch`;
    const options: APIHttpOptions = {
      ...this.options,
      ...requestHttpOptions,
    };

    if ('filter' in args) {
      options.params = options.params.set('filter', String(args.filter));
    }
    if ('pageSize' in args) {
      options.params = options.params.set('pageSize', String(args.pageSize));
    }
    if ('page' in args) {
      options.params = options.params.set('page', String(args.page));
    }
    return this.sendRequest<models.Thrl6pPatchList>('GET', path, options);
  }

  /**
   * Uploads a new patch.
   * Response generated for [ 200 ] HTTP response code.
   */
  createPatch(
    args: {
      body: models.Thrl6pCreatePatchRequest,
    },
    requestHttpOptions?: HttpOptions
  ): Observable<models.Thrl6pPatch> {
    const path = `/api/patch`;
    const options: APIHttpOptions = {
      ...this.options,
      ...requestHttpOptions,
    };

    return this.sendRequest<models.Thrl6pPatch>('POST', path, options, JSON.stringify(args.body));
  }

  /**
   * Metadata for the patch properties.
   * Response generated for [ 200 ] HTTP response code.
   */
  metadata(
    requestHttpOptions?: HttpOptions
  ): Observable<models.Thrl6pMeta> {
    const path = `/api/patch/metadata`;
    const options: APIHttpOptions = {
      ...this.options,
      ...requestHttpOptions,
    };

    return this.sendRequest<models.Thrl6pMeta>('GET', path, options);
  }

  /**
   * Validate a patch name.
   * Response generated for [ 200 ] HTTP response code.
   */
  validateName(
    args: {
      body: models.Thrl6pValidateNameRequest,
    },
    requestHttpOptions?: HttpOptions
  ): Observable<object> {
    const path = `/api/patch/validate/name`;
    const options: APIHttpOptions = {
      ...this.options,
      ...requestHttpOptions,
    };

    return this.sendRequest<object>('POST', path, options, JSON.stringify(args.body));
  }

  /**
   * Gets a patch with associated metadata.
   * Response generated for [ 200 ] HTTP response code.
   */
  getPatch(
    args: {
      id: string,
    },
    requestHttpOptions?: HttpOptions
  ): Observable<models.Thrl6pPatch> {
    const path = `/api/patch/${args.id}`;
    const options: APIHttpOptions = {
      ...this.options,
      ...requestHttpOptions,
    };

    return this.sendRequest<models.Thrl6pPatch>('GET', path, options);
  }

  private sendRequest<T>(method: string, path: string, options: HttpOptions, body?: any): Observable<T> {
    switch (method) {
      case 'DELETE':
        return this.http.delete<T>(`${this.domain}${path}`, options);
      case 'GET':
        return this.http.get<T>(`${this.domain}${path}`, options);
      case 'HEAD':
        return this.http.head<T>(`${this.domain}${path}`, options);
      case 'OPTIONS':
        return this.http.options<T>(`${this.domain}${path}`, options);
      case 'PATCH':
        return this.http.patch<T>(`${this.domain}${path}`, body, options);
      case 'POST':
        return this.http.post<T>(`${this.domain}${path}`, body, options);
      case 'PUT':
        return this.http.put<T>(`${this.domain}${path}`, body, options);
      default:
        console.error(`Unsupported request: ${method}`);
        return throwError(`Unsupported request: ${method}`);
    }
  }
}
