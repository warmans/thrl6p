import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OnIndicatorComponent } from './on-indicator.component';

describe('OnIndicatorComponent', () => {
  let component: OnIndicatorComponent;
  let fixture: ComponentFixture<OnIndicatorComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OnIndicatorComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OnIndicatorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
