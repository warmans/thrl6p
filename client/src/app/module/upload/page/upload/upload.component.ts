import { Component, OnInit } from '@angular/core';
import { ThrFile } from '../../component/file-loader/file-loader.component';
import { THRL6P } from '../../../../lib/thrl6p/thrl6p';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { PatchServiceAPIClient } from '../../../../lib/api-client/services/patch-service';
import { Router } from '@angular/router';
import { Thrl6pPatch } from '../../../../lib/api-client/models';

@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.scss']
})
export class UploadComponent implements OnInit {

  /**
   * Raw patch file
   */
  file: ThrFile = null;

  /**
   * Parsed patch file
   */
  patch: THRL6P = null;

  metadataForm = new FormGroup({
    name: new FormControl(null, [Validators.required, Validators.maxLength(64)]),
    description: new FormControl(null, [Validators.maxLength(1024)])
  });

  constructor(private apiClient: PatchServiceAPIClient, private router: Router) {
  }

  ngOnInit(): void {
  }

  f(field: string) {
    return this.metadataForm.get(field);
  }

  fileLoaded(f: ThrFile) {
    console.log(f);
    this.file = f;
    this.patch = JSON.parse(f.data) as THRL6P;

    this.f('name').setValue(this.patch?.data?.meta?.name);
    this.nameChanged();
  }

  submit() {
    this.apiClient
      .createPatch(
        {
          body: {
            description: this.metadataForm.get('description').value,
            patch: JSON.stringify(this.patch),
          }
        }
      ).subscribe((res: Thrl6pPatch) => {
        console.log(res);
        this.router.navigateByUrl(`patch/${res.id}`);
      },
    );
  }

  cancel() {
    this.file = null;
    this.patch = null;
    this.metadataForm.reset({ name: '', description: '' });
  }

  nameChanged() {
    const nameValue = this.f('name').value;
    this.apiClient.validateName({ body: { name: nameValue } }).subscribe((resp) => {
        this.patch.data.meta.name = nameValue;
      },
      (err) => {
        this.f('name').setErrors({
          ...this.f('name').errors,
          nameNotAvailableReason: err.error.details[0].description
        });
      });
  }
}
