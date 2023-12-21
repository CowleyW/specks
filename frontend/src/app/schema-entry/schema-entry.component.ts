import {Component, Input} from '@angular/core';
import {FormArray, FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators} from "@angular/forms";
import {NgForOf} from "@angular/common";

@Component({
  selector: 'app-schema-entry',
  standalone: true,
  imports: [
    FormsModule,
    NgForOf,
    ReactiveFormsModule
  ],
  templateUrl: './schema-entry.component.html',
  styleUrl: './schema-entry.component.css'
})
export class SchemaEntryComponent {
  @Input() schema!: FormGroup;

  getColumns(): FormGroup[] {
    return (this.schema.get('columns') as FormArray).controls as FormGroup[];
  }

  addNewColumn() {
    let columnForm = this.formBuilder.group({
      columnName: ['', Validators.required],
      columnType: ['', Validators.required],
    });

    let cols = this.schema.get('columns') as FormArray;
    cols.push(columnForm);
  }

  removeColumn(columnIndex: number): void {
    (this.schema.get('columns') as FormArray).removeAt(columnIndex);
  }

  constructor(private formBuilder: FormBuilder) {
  }
}
