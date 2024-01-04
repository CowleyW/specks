export interface ColumnType {
  name: string;
}

export class BasicColumnType implements ColumnType {
  constructor(name: string) {
    this.name = name;
  }

  name: string;
}

export class BoundedColumnType implements ColumnType {
  name: string;
  min: number;
  max: number;

  constructor(name: string, min: number, max: number) {
    this.name = name;
    this.min = min;
    this.max = max;
  }
}
