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
    name: new FormControl(null, [Validators.required]),
    description: new FormControl(null, [Validators.maxLength(1024)])
  });

  constructor(private apiClient: PatchServiceAPIClient, private router: Router) {
  }

  ngOnInit(): void {
  }

  fileLoaded(f: ThrFile) {
    console.log(f);
    this.file = f;
    this.patch = JSON.parse(f.data) as THRL6P;

    this.metadataForm.get('name').setValue(this.patch?.data?.meta?.name);
  }

  submit() {
    this.apiClient
      .createPatch(
        {
          body: {
            name: this.metadataForm.get('name').value,
            description: this.metadataForm.get('description').value,
            patch: this.file.data
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

}
