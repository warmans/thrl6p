import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FormControl, FormGroup, ValidatorFn, Validators } from '@angular/forms';

const MAX_FILE_SIZE = 10240; // bytes

export interface ThrFile {
  name: string;
  data: string;
}

@Component({
  selector: 'app-file-loader',
  templateUrl: './file-loader.component.html',
  styleUrls: ['./file-loader.component.scss']
})
export class FileLoaderComponent implements OnInit {

  @Output()
  loaded: EventEmitter<ThrFile> = new EventEmitter<ThrFile>();

  uploadForm = new FormGroup({
    patch: new FormControl(null, [Validators.required, requiredFileType('thrl6p')])
  });

  constructor() {
  }

  ngOnInit(): void {

  }

  get patch() {
    return this.uploadForm.get('patch');
  }

  fileChangeEvent(f: File[]) {

    if (!this.patch.valid) {
      return;
    }

    if (f.length !== 1) {
      this.patch.setErrors({ wrongFileCount: true });
      return;
    }
    const file = f[0];

    const reader = new FileReader();
    reader.onload = ((evt: File) => {
      return (e) => {
        const realSize = stringSizeInBytes(e?.target?.result || '');
        // double check the file size is actually the same as reported.
        if (realSize > MAX_FILE_SIZE) {
          this.patch.setErrors({ fileTooBig: true });
          return;
        }
        if (realSize <= 0) {
          this.patch.setErrors({ fileTooSmall: true });
          return;
        }
        if (this.patch.valid) {
          this.loaded.next({ name: file.name, data: e.target.result });
        }
      };
    })(file);
    reader.readAsText(file);
  }
}

export function requiredFileType(type: string): ValidatorFn {
  return (control: FormControl) => {
    const file = control.value;
    if (file) {
      const extension = (file.name || file).split('.').pop().toLowerCase();
      if (type.toLowerCase() !== extension) {
        return {
          requiredFileType: true
        };
      }
      return null;
    }
    return null;
  };
}

function stringSizeInBytes(s: string): number {
  return (new TextEncoder().encode(s)).length;
}
