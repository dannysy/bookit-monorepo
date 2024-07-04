import { ChangeDetectionStrategy, Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { RouterModule } from '@angular/router';
import { CardDirective } from '@shared/directives/card.directive';

type NavItem = {
  icon: string;
  label: string;
  routerLink: string[];
  disabled?: boolean;
};

const NAV_ITEMS: NavItem[] = [
  {
    icon: 'home',
    label: 'Рабочий стол',
    routerLink: ['workspace'],
  },
  {
    icon: 'people',
    label: 'Сотрудники',
    routerLink: ['employees'],
  },
  {
    icon: 'cube',
    label: 'Проекты',
    routerLink: ['projects'],
  },
  {
    icon: 'opened-book',
    label: 'Учебный центр',
    routerLink: ['study-center'],
    disabled: true,
  },
  {
    icon: 'document',
    label: 'Документы',
    routerLink: ['documents'],
  },
];

@Component({
  selector: 'app-nav-menu',
  standalone: true,
  imports: [
    CommonModule,
    NzIconModule,
    RouterModule,
  ],
  templateUrl: './nav-menu.component.html',
  styleUrls: ['./nav-menu.component.less'],
  changeDetection: ChangeDetectionStrategy.OnPush,
  hostDirectives: [
    CardDirective,
  ],
})
export class NavMenuComponent {
  public items: NavItem[] = NAV_ITEMS;
}
