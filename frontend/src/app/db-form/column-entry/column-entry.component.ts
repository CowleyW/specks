import {Component, Input, OnInit} from '@angular/core';
import {Form, FormControl, FormGroup, FormsModule, ReactiveFormsModule} from "@angular/forms";
import {DatePipe, JsonPipe, NgForOf, NgIf} from "@angular/common";
import {BasicColumnType} from "../../services/type.service";
import {TablesService} from "../../services/tables.service";
import {ColumnTypeMaxDatePipe, ColumnTypeMinDatePipe} from "../../pipes/column-type-pipe";

@Component({
  selector: '[app-column-entry]',
  standalone: true,
  imports: [
    FormsModule,
    NgForOf,
    ReactiveFormsModule,
    NgIf,
    JsonPipe,
    ColumnTypeMinDatePipe,
    DatePipe,
    ColumnTypeMaxDatePipe
  ],
  templateUrl: './column-entry.component.html',
  styleUrl: './column-entry.component.css'
})
export class ColumnEntryComponent implements OnInit {
  @Input() column!: FormGroup;
  @Input() columnIdx!: number;
  @Input() tableIdx!: number;

  columnType!: FormGroup;
  columnMin!: FormControl;

  isTypeDate: boolean;
  isTypeRange: boolean;
  isTypeTime: boolean;

  constructor(protected tables: TablesService) {
    this.isTypeTime = false;
    this.isTypeDate = false;
    this.isTypeRange = false;
  }

  ngOnInit() {
    this.columnType = this.column.get('columnType')! as FormGroup;
    this.columnMin = this.columnType.get('min')! as FormControl;
  }

  initType(event: Event) {
    const type = (event.target as HTMLSelectElement).value;

    this.columnType.get('min')!.setValue('');
    this.columnType.get('max')!.setValue('');
    this.columnType.get('format')!.setValue('');

    this.isTypeTime = false;
    this.isTypeDate = false;
    this.isTypeRange = false;

    switch (type) {
      case 'Date':
      case 'Datetime':
        let today = new Date();
        let oneYearAgo = new Date(today);
        oneYearAgo.setFullYear(today.getFullYear() - 1);

        this.columnType.get('min')!.setValue(oneYearAgo.toISOString().slice(0, 10));
        this.columnType.get('max')!.setValue(today.toISOString().slice(0, 10));
        this.columnType.get('format')!.setValue("YYYY-MM-DD");
        this.isTypeDate = true;

        break;
      case 'Time':
        this.columnType.get('min')!.setValue("00:00");
        this.columnType.get('max')!.setValue("23:59");
        this.columnType.get('format')!.setValue("hh:mm:ss");
        this.isTypeTime = true;

        break;
      case 'Age':
        this.columnType.get('min')!.setValue(18);
        this.columnType.get('max')!.setValue(100);
        this.isTypeRange = true;

        break;
      case 'Random Number':
        this.columnType.get('min')!.setValue(0);
        this.columnType.get('max')!.setValue(100);
        this.isTypeRange = true;

        break;
      default:
        break;
    }
  }
}
