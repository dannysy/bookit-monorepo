import {CanActivateFn, Router} from '@angular/router';
import {inject} from "@angular/core";
import {HttpClient, HttpParams} from "@angular/common/http";
import {ConfigService} from "@core/services/config.service";
import {UserService} from "@core/services/user.service";
import {TokenInfo} from "@core/models/token-info";
import {Store} from "@ngrx/store";
import {appUserActions} from "@core/state/actions/app-user.actions";

export const tokenCanActivate: CanActivateFn = (route, state) => {
  const cfg = inject(ConfigService)
  const user = inject(UserService)
  const client = inject(HttpClient)
  const router = inject(Router);
  const store = inject(Store);
  if (route.queryParamMap.has('code') && route.queryParamMap.has('state')) {
    const code = route.queryParamMap.get('code') ?? '';
    const state = route.queryParamMap.get('state') ?? '';
    client.get<TokenInfo>(cfg.apiUrlWithPostfix + '/signin',
      {params: new HttpParams().set('code', code).set('state', state)}).
      subscribe((res: TokenInfo) => {
        user.saveInLocalStorage(res)
        store.dispatch(appUserActions.updateTokenInfo({tokenInfo: res}))
        user.getUserByToken().subscribe((res) => {
          store.dispatch(appUserActions.loadUserSuccess({user: res}))
        })
        router.navigate(['app'])
      })
  }
  return false
};
