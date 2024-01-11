import { Component } from '@angular/core';
import {FormBuilder, FormGroup, ReactiveFormsModule} from "@angular/forms";
import {CommonModule} from "@angular/common";
import {SchemaEntryComponent} from "../schema-entry/schema-entry.component";
import {ApiService} from "../services/api.service";
import {TablesService} from "../services/tables.service";
import {ConverterFactory, Format} from "../services/converter";
import {JSONFormatPipe} from "../pipes/format";

@Component({
  selector: 'app-db-form',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, SchemaEntryComponent, JSONFormatPipe],
  templateUrl: './db-form.component.html',
  styleUrl: './db-form.component.css'
})
export class DbFormComponent {
  dbForm: FormGroup;
  api: ApiService;

  generateClicked: boolean;

  previewData?: string;

  constructor(protected tables: TablesService, api: ApiService) {
    this.api = api;
    this.dbForm = tables.tablesForm;
    tables.addNewTable();

    this.generateClicked = false;
  }

  onSubmit(): void {
    if (this.generateClicked) {
      this.generateData();
    } else {
      this.preview();
    }
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

  preview() {
    this.api.generatePreview(this.tables.toJSON()).subscribe({
      next: (response) => {
        console.log("Success\n", JSON.stringify(response));

        const format: Format = this.dbForm.get('outputFormat')!.value;
        const converter = ConverterFactory.createConverter(format);

        if (converter != null) {
          this.previewData = response;
        }
      },
      error: (error) => console.error("Error generating preview data\n", error)
    })
  }

  onGenerateClick() {
    this.generateClicked = true;
  }

  onPreviewClick() {
    this.generateClicked = false;
  }

  clearPreview() {
    this.previewData = "";
  }

  protected readonly Format = Format;
}
