import { Component } from '@angular/core';
import {FormArray, FormBuilder, FormGroup, ReactiveFormsModule, Validators} from "@angular/forms";
import {CommonModule} from "@angular/common";
import {SchemaEntryComponent} from "../schema-entry/schema-entry.component";
import {ApiService} from "../services/api.service";
import {TablesService} from "../services/tables.service";
import {ConverterFactory, Format} from "../services/converter";

@Component({
  selector: 'app-db-form',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, SchemaEntryComponent],
  templateUrl: './db-form.component.html',
  styleUrl: './db-form.component.css'
})
export class DbFormComponent {
  dbForm: FormGroup;
  api: ApiService;

  constructor(protected tables: TablesService, api: ApiService, fb: FormBuilder) {
    this.api = api;
    this.dbForm = tables.tablesForm;
    tables.addNewTable();
  }

  onSubmit() {
    if (!this.dbForm.valid) {
      console.log("DB Form is invalid");
      return;
    }

    this.generateData();
  }

  generateData() {
    this.api.generateData(this.tables.toJSON()).subscribe({
      next: (response) => {
        console.log("Success\n", JSON.stringify(response));

        const format: Format = this.dbForm.get('outputFormat')!.value;
        const converter = ConverterFactory.createConverter(format);

        if (converter != null) {
          console.log(`Result: ${converter.convert(response)}`);
        } else {
          console.log(format);
        }
      },
      error: (error) => console.error("Error generating data\n", error)
    });
  }

  protected readonly Format = Format;
}
