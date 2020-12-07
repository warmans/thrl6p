import { HttpErrorResponse } from '@angular/common/http';

export function fieldErrors(err: HttpErrorResponse, fieldName: string): string[] {
  if (err === null || err === undefined) {
    return [];
  }
  const errors: string[] = [];
  if (err?.error?.details?.length > 0) {
    for (const v of err.error.details) {
      switch (v['@type']) {
        case 'type.googleapis.com/google.rpc.BadRequest':
          if (v?.field_violations?.length > 0) {
            for (const f of v.field_violations) {
              if (f.field === fieldName) {
                errors.push(f.description);
              }
            }
          }
          break;
      }
    }
  }
  return errors;
}
