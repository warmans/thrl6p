import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { PatchServiceAPIClient } from '../../../../lib/api-client/services/patch-service';
import { Thrl6pPatch } from '../../../../lib/api-client/models';
import { THRL6P } from '../../../../lib/thrl6p/thrl6p';

@Component({
  selector: 'app-view',
  templateUrl: './view.component.html',
  styleUrls: ['./view.component.scss']
})
export class ViewComponent implements OnInit {

  patch: Thrl6pPatch = null;
  parsed: THRL6P = null;

  constructor(private route: ActivatedRoute, private apiClient: PatchServiceAPIClient) {
  }

  ngOnInit(): void {
    this.apiClient.getPatch({ id: this.route.snapshot.params.id }).subscribe(
      (res) => {
        this.patch = res;
        this.parsed = JSON.parse(res.patch) as THRL6P;
      }
    );
  }
}
