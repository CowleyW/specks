import {Component, Input, OnInit} from '@angular/core';
import {Form, FormControl, FormGroup, FormsModule, ReactiveFormsModule} from "@angular/forms";
import {DatePipe, JsonPipe, NgForOf, NgIf} from "@angular/common";
import {BasicColumnType} from "../../services/type.service";
import {TablesService} from "../../services/tables.service";
import {ColumnTypeMaxDatePipe, ColumnTypeMinDatePipe} from "../../pipes/column-type-pipe";

@Component({
  selector: 'app-type-modifier',
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
  templateUrl: './type-modifier.component.html',
  styleUrl: './type-modifier.component.css'
})
export class TypeModifierComponent implements OnInit {
  @Input() column!: FormGroup;
  @Input() columnIdx!: number;
  @Input() tableIdx!: number;

  columnType!: FormGroup;
  columnMin!: FormControl;

  constructor(protected tables: TablesService) {}

  ngOnInit() {
    this.columnType = this.column.get('columnType')! as FormGroup;
    this.columnMin = this.columnType.get('min')! as FormControl;
  }

  initType(event: Event) {
    const type = (event.target as HTMLSelectElement).value;

    this.columnType.get('min')!.setValue('');
    this.columnType.get('max')!.setValue('');
    this.columnType.get('format')!.setValue('');

    switch (type) {
      case 'Date':
        let today = new Date();
        let oneYearAgo = new Date(today);
        oneYearAgo.setFullYear(today.getFullYear() - 1);

        this.columnType.get('min')!.setValue(oneYearAgo.toISOString().slice(0, 10));
        this.columnType.get('max')!.setValue(today.toISOString().slice(0, 10));
        this.columnType.get('format')!.setValue("YYYY-MM-DD");

        break;
      case 'Age':
        this.columnType.get('min')!.setValue(18);
        this.columnType.get('max')!.setValue(100);
        break;
      case 'Random Number':
        this.columnType.get('min')!.setValue(0);
        this.columnType.get('max')!.setValue(100);
        break;
      default:
        break;
    }
  }
}
