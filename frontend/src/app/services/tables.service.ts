import {Injectable, Injector} from "@angular/core";
import {FormArray, FormBuilder, FormGroup, Validators} from "@angular/forms";
import {BasicColumnType, BoundedColumnType, ColumnType} from "./type.service";
import {Format} from "./converter";

@Injectable({
  providedIn: 'root'
})
export class TablesService {
  private formBuilder: FormBuilder;

  tablesForm: FormGroup;

  public types: ColumnType[] = [
    new BasicColumnType("First Name"),
    new BasicColumnType("Last Name"),
    new BoundedColumnType("Age", 18, 100),
    new BasicColumnType("Color"),
    new BasicColumnType("Boolean"),
    new BasicColumnType("SSN"),
    new BasicColumnType("Row Number"),
    new BoundedColumnType("Random Number", 0, 1000),
    new BasicColumnType("Date"),
    new BasicColumnType("Time"),
    new BasicColumnType("Datetime")
  ];

  public dateFormats: string[] = [
    "YYYY-MM-DD"
  ]

  public timeFormats: string[] = [
    "hh:mm:ss"
  ]

  constructor(private injector: Injector) {
    this.formBuilder = this.injector.get(FormBuilder);

    this.tablesForm = this.newDbForm();
  }

  private newDbForm(): FormGroup {
    return this.formBuilder.group({
      tables: this.formBuilder.array([]),
      outputFormat: [Format.CSV, Validators.required],
      forPreview: [false, Validators.required]
    });
  }

  newTableForm(): FormGroup {
    const tableNum = this.getTables().length + 1;
    return this.formBuilder.group({
      tableName: [`Table ${tableNum}`, Validators.required],
      columns: this.formBuilder.array([]),
      references: this.formBuilder.array([]),
      numRows: [100, Validators.required]
    });
  }

  newColumnForm(tableIndex: number): FormGroup {
    const columnNum = this.getColumns(tableIndex).length + 1;
    return this.formBuilder.group({
      columnName: [`Column ${columnNum}`, Validators.required],
      columnType: this.formBuilder.group({
        name: ['', Validators.required],
        min: [''],
        max: [''],
        format: ['']
      }),
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

    this.addNewColumn(this.getTables().length - 1);
  }

  getTables(): FormGroup[] {
    return (this.tablesForm.get('tables') as FormArray).controls as FormGroup[];
  }

  getTablesAbove(tableIndex: number): FormGroup[] {
    return this.getTables().slice(0, tableIndex + 1);
  }

  getTable(tableIndex: number) {
    return this.getTables()[tableIndex];
  }

  getPrimaryKeys(tableIndex: number): FormGroup[] {
    return this.getColumns(tableIndex).filter((col) => col.get('columnPrimaryKey')!.value);
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
      const split = value.split("-", 2).map((str) => parseInt(str));

      referenceForm.get('tableIndex')!.setValue(split[0]);
      referenceForm.get('columnIndex')!.setValue(split[1]);
    });

    const table = this.getTable(tableIndex);
    let refs = table.get('references') as FormArray;
    refs.push(referenceForm);
  }

  removeReference(tableIndex: number, referenceIndex: number) {
    const table = this.getTable(tableIndex);
    (table.get('references') as FormArray).removeAt(referenceIndex);
  }

  hasReferencableColumn(tableIndex: number): boolean {
    const tablesAbove = this.getTablesAbove(tableIndex);
    for (let i = 0; i < tablesAbove.length; i += 1) {
      if (this.getPrimaryKeys(i).length != 0) {
        return true;
      }
    }

    return false;
  }

  setGeneratePreview(forPreview: boolean) {
    this.tablesForm.get('forPreview')!.setValue(forPreview);
  }

  toJSON() {
    return {
      tablesDescs: this.getTables().map((table: FormGroup, idx: number) => {
        return {
          tableName: table.get('tableName')!.value,
          columns: this.getColumns(idx).map((c: FormGroup) => {
            return {
              columnName: c.get('columnName')!.value,
              columnType: this.columnTypeToJSON(c.get('columnType')! as FormGroup),
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
              referenceUnique: r.get('referenceUnique')!.value,
              default: null
            };
          }),
          numRows: table.get('numRows')!.value,
        };
      }),
      outputFormat: this.tablesForm.get('outputFormat')!.value,
      forPreview: this.tablesForm.get('forPreview')!.value
    };
  }

  columnTypeToJSON(columnType: FormGroup): any {
    switch (columnType.get('name')!.value) {
      case 'Date':
      case 'Datetime':
        return {
          name: columnType.get('name')!.value,
          min: columnType.get('min')!.value,
          max: columnType.get('max')!.value,
          format: columnType.get('format')!.value
        }
      case 'Time':
        return {
          name: columnType.get('name')!.value,
          min: `${columnType.get('min')!.value}:00`,
          max: `${columnType.get('max')!.value}:59`,
          format: columnType.get('format')!.value
        }
      case 'Age':
      case 'Random Number':
        return {
          name: columnType.get('name')!.value,
          min: columnType.get('min')!.value,
          max: columnType.get('max')!.value
        }
      default:
        return {name: columnType.get('name')!.value}
    }
  }
}
