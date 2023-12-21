import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SchemaEntryComponent } from './schema-entry.component';

describe('SchemaEntryComponent', () => {
  let component: SchemaEntryComponent;
  let fixture: ComponentFixture<SchemaEntryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SchemaEntryComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(SchemaEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
