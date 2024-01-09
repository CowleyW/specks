import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ColumnEntryComponent } from './column-entry.component';

describe('TypeModifierComponent', () => {
  let component: ColumnEntryComponent;
  let fixture: ComponentFixture<ColumnEntryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ColumnEntryComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ColumnEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
