import { Component, OnInit } from '@angular/core';
import { ThrFile } from '../../component/file-loader/file-loader.component';
import { THRL6P } from '../../../../lib/thrl6p/thrl6p';

@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.scss']
})
export class UploadComponent implements OnInit {

  file: ThrFile = null;
  patch: THRL6P = null;

  showRaw: boolean;

  constructor() {
  }

  ngOnInit(): void {
  }

  fileLoaded(f: ThrFile) {
    console.log(f);
    this.file = f;
    this.patch = JSON.parse(f.data) as THRL6P;
  }

}
