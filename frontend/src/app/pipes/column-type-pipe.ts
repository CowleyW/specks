import {Pipe, PipeTransform} from "@angular/core";
import {FormGroup} from "@angular/forms";

@Pipe({
  standalone: true,
  name: 'columnTypeMinDate'
  })
export class ColumnTypeMinDatePipe implements PipeTransform {
  transform(value: FormGroup): Date {
    console.log(value.get('min')!.value, new Date(Date.parse(value.get('min')!.value)))
    return new Date(Date.parse(value.get('min')!.value));
  }
}

@Pipe({
  standalone: true,
  name: 'columnTypeMaxDate'
})
export class ColumnTypeMaxDatePipe implements PipeTransform {
  transform(value: FormGroup): Date {
    return new Date(Date.parse(value.get('max')!.value));
  }
}
