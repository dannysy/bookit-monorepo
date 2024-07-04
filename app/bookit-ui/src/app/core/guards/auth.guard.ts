import { CanActivateFn, Router } from '@angular/router';
import { Store } from '@ngrx/store';
import { first, mergeMap, of } from 'rxjs';
import { selectCurrentUser } from '@core/state/app.selectors';
import {inject} from "@angular/core";
import {AuthService} from "@core/services/auth.service";

export const authCanActivate: CanActivateFn = (route, state) => {
  const store = inject(Store);
  const auth = inject(AuthService)
  return store.select(selectCurrentUser).pipe(
    first(),
    mergeMap(user => {
      if (!user) {
        window.location.href = auth.Sdk.getSigninUrl();
       return of(false);
      }
      return of(true);
    }),
  );

};
