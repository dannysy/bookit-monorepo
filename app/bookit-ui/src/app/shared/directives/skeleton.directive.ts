import { Directive, ElementRef, Input } from '@angular/core';

const DEFAULT_HEIGHT = '20px';

const DEFAULT_WIDTH = '100px';

@Directive({
  standalone: true,
  selector: '[appSkeleton]',
})
export class SkeletonDirective {
  @Input('appSkeleton') public set isLoading(value: boolean) {
    value ? this.start() : this.stop();
  }

  @Input() public height = DEFAULT_HEIGHT;

  @Input() public width = DEFAULT_WIDTH;

  private styles: SafeAny;

  public constructor(private el: ElementRef) {
    this.el.nativeElement.style.transition = '1s';
  }

  public start(): void {
    this.el.nativeElement.classList.add('skeleton');
    this.styles = this.el.nativeElement.styles;
    this.el.nativeElement.style.minWidth = this.width;
    this.el.nativeElement.style.minHeight = this.height;
    this.el.nativeElement.style.display = 'flex';
    this.el.nativeElement.style.color = 'transparent';
  }

  public stop(): void {
    this.el.nativeElement.classList.remove('skeleton');
    this.el.nativeElement.style = this.styles;
  }
}