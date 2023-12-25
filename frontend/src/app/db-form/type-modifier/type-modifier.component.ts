import {Component, Input} from '@angular/core';
import {FormGroup, FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgForOf} from "@angular/common";
import {ColumnType} from "../../services/type.service";
import {TablesService} from "../../services/tables.service";

@Component({
  selector: 'app-type-modifier',
  standalone: true,
  imports: [
    FormsModule,
    NgForOf,
    ReactiveFormsModule
  ],
  templateUrl: './type-modifier.component.html',
  styleUrl: './type-modifier.component.css'
})
export class TypeModifierComponent {
  @Input() column!: FormGroup;
  @Input() columnIdx!: number;
  @Input() tableIdx!: number;

  constructor(protected tables: TablesService) {
  }

  protected readonly JSON = JSON;
}
