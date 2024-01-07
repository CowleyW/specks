import {Pipe, PipeTransform} from "@angular/core";

@Pipe({
  standalone: true,
  name: 'jsonFormat'
})
export class JSONFormatPipe implements PipeTransform {
  transform(value: any[]): any {
  }
}
