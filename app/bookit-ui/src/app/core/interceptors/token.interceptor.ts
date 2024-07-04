import { HttpHandlerFn, HttpInterceptorFn, HttpRequest } from '@angular/common/http';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { appUserActions } from '@core/state/actions/app-user.actions';
import { selectTokenInfo } from '@core/state/app.selectors';
import { Store } from '@ngrx/store';
import { catchError, first, map, mergeMap, of, throwError } from 'rxjs';

const UNAUTHORIZED_STATUS = 401;

export const tokenInterceptor: HttpInterceptorFn = (req: HttpRequest<unknown>, next: HttpHandlerFn) => {
  const store = inject(Store);
  const router = inject(Router);
  return store.select(selectTokenInfo).pipe(
    first(),
    map(tokenInfo => {
      if (!tokenInfo) {
        return req;
      }
      const authReq = req.clone({
        headers: req.headers.set('Authorization', `Bearer ${tokenInfo?.access_token}`),
      });
      return authReq;
    }),
    mergeMap(newReq => next(newReq).pipe(
      catchError(err => {
        if (err.status === UNAUTHORIZED_STATUS) {
          store.dispatch(appUserActions.logout());
          void router.navigate(['/auth']);
        }
        return throwError(() => err);
      }),
    )));
};
