import { Component } from '@angular/core';
import {FormArray, FormBuilder, FormGroup, ReactiveFormsModule, Validators} from "@angular/forms";
import {CommonModule} from "@angular/common";
import {SchemaEntryComponent} from "../schema-entry/schema-entry.component";
import {ApiService} from "../services/api.service";

@Component({
  selector: 'app-db-form',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, SchemaEntryComponent],
  templateUrl: './db-form.component.html',
  styleUrl: './db-form.component.css'
})
export class DbFormComponent {
  dbForm: FormGroup = this.formBuilder.group({
    schemas: this.formBuilder.array([])
  });

  api: ApiService;

  constructor(private formBuilder: FormBuilder, api: ApiService) {
    this.api = api;
    this.addNewSchema();
  }

  addNewSchema(): void {
    let schemaForm = this.formBuilder.group({
      tableName: ['hello', Validators.required],
      columns: this.formBuilder.array([])
    });

    this.getSchemas().push(schemaForm);
  }

  getSchemas(): FormGroup[] {
    return (this.dbForm.get('schemas') as FormArray).controls as FormGroup[];
  }

  removeSchema(schemaIndex: number): void {
    (this.dbForm.get('schemas') as FormArray).removeAt(schemaIndex);
  }
}
