import {Component, Input} from '@angular/core';
import {FormGroup, FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgForOf} from "@angular/common";
import {ColumnType} from "../../services/type.service";

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

  types: ColumnType[] = [
    new ColumnType("First Name"),
    new ColumnType("Last Name"),
    new ColumnType("Character"),
    new ColumnType("Age"),
    new ColumnType("Color"),
    new ColumnType("Boolean"),
    new ColumnType("SSN"),
    new ColumnType("Row Number"),
    new ColumnType("Random Number"),
    new ColumnType("Date"),
    new ColumnType("Time"),
    new ColumnType("Datetime")
  ];
  protected readonly JSON = JSON;
}
