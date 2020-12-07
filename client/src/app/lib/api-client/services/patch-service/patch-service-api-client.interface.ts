/* tslint:disable */

import { Observable } from 'rxjs';
import { HttpOptions } from './';
import * as models from '../../models';

export interface PatchServiceAPIClientInterface {

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
  ): Observable<models.Thrl6pPatchList>;

  /**
   * Uploads a new patch.
   * Response generated for [ 200 ] HTTP response code.
   */
  createPatch(
    args: {
      body: models.Thrl6pCreatePatchRequest,
    },
    requestHttpOptions?: HttpOptions
  ): Observable<models.Thrl6pPatch>;

  /**
   * Metadata for the patch properties.
   * Response generated for [ 200 ] HTTP response code.
   */
  metadata(
    requestHttpOptions?: HttpOptions
  ): Observable<models.Thrl6pMeta>;

  /**
   * Validate a patch name.
   * Response generated for [ 200 ] HTTP response code.
   */
  validateName(
    args: {
      body: models.Thrl6pValidateNameRequest,
    },
    requestHttpOptions?: HttpOptions
  ): Observable<models.Thrl6pNameValidation>;

  /**
   * Gets a patch with associated metadata.
   * Response generated for [ 200 ] HTTP response code.
   */
  getPatch(
    args: {
      id: string,
    },
    requestHttpOptions?: HttpOptions
  ): Observable<models.Thrl6pPatch>;

}
