import {Component, Input, OnInit} from '@angular/core';
import {Form, FormControl, FormGroup, FormsModule, ReactiveFormsModule} from "@angular/forms";
import {JsonPipe, NgForOf, NgIf} from "@angular/common";
import {BasicColumnType} from "../../services/type.service";
import {TablesService} from "../../services/tables.service";

@Component({
  selector: 'app-type-modifier',
  standalone: true,
  imports: [
    FormsModule,
    NgForOf,
    ReactiveFormsModule,
    NgIf,
    JsonPipe
  ],
  templateUrl: './type-modifier.component.html',
  styleUrl: './type-modifier.component.css'
})
export class TypeModifierComponent implements OnInit {
  @Input() column!: FormGroup;
  @Input() columnIdx!: number;
  @Input() tableIdx!: number;

  columnType!: FormGroup;

  constructor(protected tables: TablesService) {}

  ngOnInit() {
    this.columnType = this.column.get('columnType')! as FormGroup;
  }

  initType(event: Event) {
    const type = (event.target as HTMLSelectElement).value;

    switch (type) {
      case 'Random Number':
        this.columnType.get('min')!.setValue(0);
        this.columnType.get('max')!.setValue(100);
        break;
      default:
        this.columnType.get('min')!.setValue('');
        this.columnType.get('max')!.setValue('');
        break;
    }
  }

  protected readonly JSON = JSON;
}
