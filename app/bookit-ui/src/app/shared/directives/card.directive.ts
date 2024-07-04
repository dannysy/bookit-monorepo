import { Directive, DoCheck, ElementRef, Input } from '@angular/core';

const DEFAULT_VALUE = 0.04;

const DEFAULT_BORDER_RADIUS = 20;

@Directive({
  standalone: true,
  selector: '[appCard]',
})
export class CardDirective implements DoCheck {
  @Input() public value = DEFAULT_VALUE;

  @Input() public boderRadius = DEFAULT_BORDER_RADIUS;

  public constructor(private el: ElementRef) {
    this.el.nativeElement.style.backgroundColor = 'white';
    this.el.nativeElement.style.borderRadius = `${this.boderRadius}px`;
    this.el.nativeElement.style.boxShadow = `0px 4px 6px 0px rgba(0, 0, 0, ${this.value})`;
  }

  public ngDoCheck(): void {
    this.el.nativeElement.style.backgroundColor = 'white';
    this.el.nativeElement.style.boxShadow = `0px 4px 6px 0px rgba(0, 0, 0, ${this.value})`;
    this.el.nativeElement.style.borderRadius = `${this.boderRadius}px`;
  }
}