import { ChangeDetectionStrategy, Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { IconButtonComponent } from '@shared/icon-button/icon-button.component';
import { UserMiniatureComponent } from './components/user-miniature/user-miniature.component';
import { CardDirective } from '@shared/directives/card.directive';
import { RouterModule } from '@angular/router';
import { NzIconModule } from 'ng-zorro-antd/icon';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [
    CommonModule,
    IconButtonComponent,
    UserMiniatureComponent,
    CardDirective,
    RouterModule,
    NzIconModule,
  ],
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.less'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class HeaderComponent {}
