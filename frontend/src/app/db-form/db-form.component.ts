import { Component } from '@angular/core';
import {FormArray, FormBuilder, FormGroup, ReactiveFormsModule, Validators} from "@angular/forms";
import {CommonModule} from "@angular/common";
import {SchemaEntryComponent} from "../schema-entry/schema-entry.component";
import {ApiService} from "../services/api.service";
import {TablesService} from "../services/tables.service";

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

  constructor(protected tables: TablesService, api: ApiService) {
    this.api = api;
    this.dbForm = tables.tablesForm;
    tables.addNewTable();
  }
}
