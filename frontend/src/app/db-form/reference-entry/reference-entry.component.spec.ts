import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReferenceEntryComponent } from './reference-entry.component';

describe('ReferenceEntryComponent', () => {
  let component: ReferenceEntryComponent;
  let fixture: ComponentFixture<ReferenceEntryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ReferenceEntryComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ReferenceEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
