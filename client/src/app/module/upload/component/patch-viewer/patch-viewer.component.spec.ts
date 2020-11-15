import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PatchViewerComponent } from './patch-viewer.component';

describe('PatchViewerComponent', () => {
  let component: PatchViewerComponent;
  let fixture: ComponentFixture<PatchViewerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PatchViewerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PatchViewerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
