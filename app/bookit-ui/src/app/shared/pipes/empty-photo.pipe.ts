import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'appEmptyPhoto',
  standalone: true,
})
export class EmptyPhotoPipe implements PipeTransform {
  public transform(value: string | undefined): string {
    return value?.length ? value : '/assets/images/user.svg';
  }
}