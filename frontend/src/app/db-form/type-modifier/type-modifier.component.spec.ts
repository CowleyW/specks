import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TypeModifierComponent } from './type-modifier.component';

describe('TypeModifierComponent', () => {
  let component: TypeModifierComponent;
  let fixture: ComponentFixture<TypeModifierComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TypeModifierComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(TypeModifierComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
