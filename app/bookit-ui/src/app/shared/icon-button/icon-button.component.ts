import {
  ChangeDetectionStrategy,
  Component,
  Input,
} from '@angular/core';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { CommonModule } from '@angular/common';
import { ThemeType } from '@ant-design/icons-angular';

@Component({
  selector: 'app-icon-button',
  standalone: true,
  imports: [
    CommonModule,
    NzIconModule,
  ],
  templateUrl: './icon-button.component.html',
  styleUrls: ['./icon-button.component.less'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class IconButtonComponent {
  @Input() public icon!: string;

  @Input() public theme!: ThemeType;
}
