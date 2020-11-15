/* tslint:disable */

import { NgModule, ModuleWithProviders } from '@angular/core';
import { HttpHeaders, HttpParams } from '@angular/common/http';
import { PatchServiceAPIClient, USE_DOMAIN, USE_HTTP_OPTIONS } from './patch-service-api-client.service';
import { GuardedPatchServiceAPIClient } from './guarded-patch-service-api-client.service';

export { PatchServiceAPIClient } from './patch-service-api-client.service';
export { PatchServiceAPIClientInterface } from './patch-service-api-client.interface';
export { GuardedPatchServiceAPIClient } from './guarded-patch-service-api-client.service';

/**
 * provided options, headers and params will be used as default for each request
 */
export interface DefaultHttpOptions {
  headers?: {[key: string]: string};
  params?: {[key: string]: string};
  reportProgress?: boolean;
  withCredentials?: boolean;
}

export interface HttpOptions {
  headers?: HttpHeaders;
  params?: HttpParams;
  reportProgress?: boolean;
  withCredentials?: boolean;
}

export interface PatchServiceAPIClientModuleConfig {
  domain?: string;
  guardResponses?: boolean; // validate responses with type guards
  httpOptions?: DefaultHttpOptions;
}

@NgModule({})
export class PatchServiceAPIClientModule {
  /**
   * Use this method in your root module to provide the PatchServiceAPIClientModule
   *
   * If you are not providing
   * @param { PatchServiceAPIClientModuleConfig } config
   * @returns { ModuleWithProviders }
   */
  static forRoot(config: PatchServiceAPIClientModuleConfig = {}): ModuleWithProviders<PatchServiceAPIClientModule> {
    return {
      ngModule: PatchServiceAPIClientModule,
      providers: [
        ...(config.domain != null ? [{provide: USE_DOMAIN, useValue: config.domain}] : []),
        ...(config.httpOptions ? [{provide: USE_HTTP_OPTIONS, useValue: config.httpOptions}] : []),
        ...(config.guardResponses ? [{provide: PatchServiceAPIClient, useClass: GuardedPatchServiceAPIClient }] : [PatchServiceAPIClient]),
      ]
    };
  }
}
