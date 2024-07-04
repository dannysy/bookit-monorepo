import { ChangeDetectionStrategy, Component, HostListener, ViewChild, inject, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Store } from '@ngrx/store';
import { selectCurrentUser } from '@core/state/app.selectors';
import { SkeletonDirective } from '@shared/directives/skeleton.directive';
import { NzDropDownModule, NzDropdownMenuComponent } from 'ng-zorro-antd/dropdown';
import { appUserActions } from '@core/state/actions/app-user.actions';
import { Router } from '@angular/router';
import { EmptyPhotoPipe } from '@shared/pipes/empty-photo.pipe';

@Component({
  selector: 'app-user-miniature',
  standalone: true,
  imports: [
    CommonModule,
    SkeletonDirective,
    NzDropDownModule,
    EmptyPhotoPipe,
  ],
  templateUrl: './user-miniature.component.html',
  styleUrls: ['./user-miniature.component.less'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UserMiniatureComponent {
  public store = inject(Store);

  public router = inject(Router);

  public photoLoaded = signal(false);

  public currentUser$ = this.store.select(selectCurrentUser);

  public onLogoutClick(): void {
    this.store.dispatch(appUserActions.logout());
    void this.router.navigate(['/auth']);
  }
}
