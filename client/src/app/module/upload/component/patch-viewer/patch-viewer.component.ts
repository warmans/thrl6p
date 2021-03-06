import { Component, Input, OnInit } from '@angular/core';
import { THRL6P } from '../../../../lib/thrl6p/thrl6p';

@Component({
  selector: 'app-patch-viewer',
  templateUrl: './patch-viewer.component.html',
  styleUrls: ['./patch-viewer.component.scss']
})
export class PatchViewerComponent implements OnInit {

  @Input()
  patch: THRL6P = null;

  showRaw: boolean;

  constructor() { }

  ngOnInit(): void {
  }

  toggleRaw() {
    this.showRaw = !this.showRaw;
  }

  rawPatch(): string {
    return JSON.stringify(this.patch, null, '  ');
  }

}
