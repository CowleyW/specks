import {Component} from '@angular/core';
import {FormGroup, ReactiveFormsModule} from "@angular/forms";
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

  generateClicked: boolean;

  previewData?: string;

  constructor(protected tables: TablesService, private api: ApiService) {
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
        console.log("Success\n", response);

        let blob = new Blob([response], { type: 'application/zip'});

        const link = document.createElement('a');
        link.href = window.URL.createObjectURL(blob);

        link.download = 'output.zip';

        document.body.appendChild(link);
        link.click();

        document.body.removeChild(link);
      },
      error: (error) => console.error("Error generating data\n", error)
    });
  }

  preview() {
    switch (this.getOutputFormat()) {
      case Format.JSON:
        this.api.generateJSONPreview(this.tables.toJSON()).subscribe({
          next: (response) => {
            this.previewData = response;
          },
          error: (error) => console.error("Error generating preview data\n", error)
        });
        break;
      case Format.CSV:
      case Format.SQL:
        this.api.generateTextPreview(this.tables.toJSON()).subscribe({
          next: (response) => {
            this.previewData = response;
          },
          error: (error) => console.error("Error generating preview data\n", error)
        });
        break;
      default:
        console.error("unknown output format");
        break;
    }
  }

  onGenerateClick() {
    this.generateClicked = true;
    this.tables.setGeneratePreview(false);
  }

  onPreviewClick() {
    this.generateClicked = false;
    this.tables.setGeneratePreview(true);
  }

  clearPreview() {
    this.previewData = "";
  }

  getOutputFormat(): Format | undefined {
    switch (this.dbForm.get('outputFormat')!.value) {
      case 'CSV':
        return Format.CSV;
      case 'JSON':
        return Format.JSON;
      case 'SQL':
        return Format.SQL;
      default:
        return undefined;
    }
  }

  protected readonly Format = Format;
}
