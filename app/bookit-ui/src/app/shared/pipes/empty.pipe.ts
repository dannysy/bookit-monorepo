import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'appEmpty',
  standalone: true,
})
export class EmptyPipe implements PipeTransform {
  public transform(value: string | undefined): string {
    return value?.length ? value : 'Не указано';
  }
}