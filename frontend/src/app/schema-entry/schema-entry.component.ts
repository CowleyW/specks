import {Component, Input} from '@angular/core';
import {
  FormArray,
  FormBuilder,
  FormControl,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators
} from "@angular/forms";
import {NgForOf} from "@angular/common";
import {ColumnEntryComponent} from "../db-form/column-entry/column-entry.component";
import {BasicColumnType} from "../services/type.service";
import {TablesService} from "../services/tables.service";
import {ReferenceEntryComponent} from "../db-form/reference-entry/reference-entry.component";

@Component({
  selector: 'app-schema-entry',
  standalone: true,
  imports: [
    NgForOf,
    ReactiveFormsModule,
    ColumnEntryComponent,
    ReferenceEntryComponent
  ],
  templateUrl: './schema-entry.component.html',
  styleUrl: './schema-entry.component.css'
})
export class SchemaEntryComponent {
  @Input() table!: FormGroup;
  @Input() tableIdx!: number;

  constructor(protected tables: TablesService) {
  }
}
