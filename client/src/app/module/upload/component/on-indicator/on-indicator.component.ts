import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-on-indicator',
  templateUrl: './on-indicator.component.html',
  styleUrls: ['./on-indicator.component.scss']
})
export class OnIndicatorComponent implements OnInit {

  @Input()
  on: boolean;

  constructor() { }

  ngOnInit(): void {
  }

}
