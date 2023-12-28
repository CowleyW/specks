import {Injectable, Injector} from "@angular/core";
import {FormArray, FormBuilder, FormGroup, Validators} from "@angular/forms";
import {ColumnType} from "./type.service";

@Injectable({
  providedIn: 'root'
})
export class TablesService {
  private formBuilder: FormBuilder;

  // TODO: Make private
  tablesForm: FormGroup;

  public types: ColumnType[] = [
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

  constructor(private injector: Injector) {
    this.formBuilder = this.injector.get(FormBuilder);

    this.tablesForm = this.newDbForm();
  }

  private newDbForm(): FormGroup {
    return this.formBuilder.group({
      tables: this.formBuilder.array([])
    });
  }

  newTableForm(): FormGroup {
    const tableNum = this.getTables().length + 1;
    return this.formBuilder.group({
      tableName: [`Table ${tableNum}`, Validators.required],
      columns: this.formBuilder.array([]),
      references: this.formBuilder.array([])
    });
  }

  newColumnForm(tableIndex: number): FormGroup {
    const columnNum = this.getColumns(tableIndex).length + 1;
    return this.formBuilder.group({
      columnName: [`Column ${columnNum}`, Validators.required],
      columnType: [null, Validators.required],
      columnPrimaryKey: [false, Validators.required],
      columnUnique: [false, Validators.required]
    });
  }

  newReferenceForm(tableIndex: number): FormGroup {
    const referenceNum = this.getReferences(tableIndex).length + 1;
    return this.formBuilder.group({
      referenceName: [`Reference ${referenceNum}`, Validators.required],
      referenceColumn: [null, Validators.required],
      tableIndex: ['', Validators.required],
      columnIndex: ['', Validators.required],
      referencePrimaryKey: [false, Validators.required],
      referenceUnique: [false, Validators.required]
    });
  }

  addNewTable(): void {
    let tableForm = this.newTableForm();
    this.getTables().push(tableForm);
  }

  getTables(): FormGroup[] {
    return (this.tablesForm.get('tables') as FormArray).controls as FormGroup[];
  }

  getTable(tableIndex: number) {
    return this.getTables()[tableIndex];
  }

  getPrimaryKeys(tableIndex: number): FormGroup[] {
    const cols: FormGroup[] = this.getColumns(tableIndex).filter((col) => col.get('columnPrimaryKey')!.value);
    console.log(cols.length);
    return cols;
  }

  removeTable(tableIndex: number): void {
    (this.tablesForm.get('tables') as FormArray).removeAt(tableIndex);
  }
  getColumns(tableIndex: number): FormGroup[] {
    const table = this.getTable(tableIndex);
    return (table.get('columns') as FormArray).controls as FormGroup[];
  }

  addNewColumn(tableIndex: number) {
    let columnForm = this.newColumnForm(tableIndex);

    const table = this.getTable(tableIndex);
    let cols = table.get('columns') as FormArray;
    cols.push(columnForm);
  }

  removeColumn(tableIndex: number, columnIndex: number): void {
    const table = this.getTable(tableIndex);
    (table.get('columns') as FormArray).removeAt(columnIndex);
  }

  getReferences(tableIndex: number): FormGroup[] {
    const table = this.getTable(tableIndex);
    return (table.get('references') as FormArray).controls as FormGroup[];
  }

  addNewReference(tableIndex: number) {
    let referenceForm = this.newReferenceForm(tableIndex);

    referenceForm.get('referenceColumn')!.valueChanges.subscribe((value: string) => {
      // Format: {tableIndex}-{columnIndex}
      const indices = value.split("-", 2).map((str: string) => parseInt(str));

      referenceForm.get('tableIndex')!.setValue(indices[0]);
      referenceForm.get('columnIndex')!.setValue(indices[1]);
    });

    const table = this.getTable(tableIndex);
    let refs = table.get('references') as FormArray;
    refs.push(referenceForm);
  }

  removeReference(tableIndex: number, referenceIndex: number) {
    const table = this.getTable(tableIndex);
    (table.get('references') as FormArray).removeAt(referenceIndex);
  }

  toJSON() {
    return this.getTables().map((table: FormGroup, idx: number) => {
      return {
        tableName: table.get('tableName')!.value,
        columns: this.getColumns(idx).map((c: FormGroup) => {
          return {
            columnName: c.get('columnName')!.value,
            columnType: JSON.parse(c.get('columnType')!.value),
            columnPrimaryKey: c.get('columnPrimaryKey')!.value,
            columnUnique: c.get('columnUnique')!.value
          };
        }),
        references: this.getReferences(idx).map((r: FormGroup) => {
          return {
            referenceName: r.get('referenceName')!.value,
            tableIndex: r.get('tableIndex')!.value,
            columnIndex: r.get('columnIndex')!.value,
            referencePrimaryKey: r.get('referencePrimaryKey')!.value,
            referenceUnique: r.get('referenceUnique')!.value
          };
        })
      };
    });
  }
}
