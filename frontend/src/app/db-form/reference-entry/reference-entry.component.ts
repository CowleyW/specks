import {Component, Input} from '@angular/core';
import {FormGroup, ReactiveFormsModule} from "@angular/forms";
import {TablesService} from "../../services/tables.service";
import {NgForOf, NgIf} from "@angular/common";

@Component({
  selector: '[app-reference-entry]',
  standalone: true,
  imports: [
    NgForOf,
    ReactiveFormsModule,
    NgIf
  ],
  templateUrl: './reference-entry.component.html',
  styleUrl: './reference-entry.component.css'
})
export class ReferenceEntryComponent {
  @Input() reference!: FormGroup;
  @Input() referenceIdx!: number;
  @Input() tableIdx!: number;

  constructor(protected tables: TablesService) {}
}
