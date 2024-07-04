import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, mergeMap, of, tap } from 'rxjs';
import { ConfigService } from './config.service';
import { TokenInfo } from '@core/models/token-info';
import { Store } from '@ngrx/store';
import { appUserActions } from '@core/state/actions/app-user.actions';
import { User } from '@core/models/user';

const TOKEN_LOCAL_STORAGE_KEY = 'token';

const EXPIRED_IN_DISTINCTION = 10 * 1000;

@Injectable({ providedIn: 'root' })
export class UserService {
  private timerId?: number;

  public constructor(
    private readonly _http: HttpClient,
    private readonly _config: ConfigService,
    private readonly store: Store,
  ) {}

  public login(username: string, password: string): Observable<User> {
    return this.auth(username, password).pipe(
      tap(tokenInfo => {
        this.store.dispatch(appUserActions.updateTokenInfo({ tokenInfo }));
        this.setRefreshSchedule(tokenInfo);
        this.saveInLocalStorage(tokenInfo);
      }),
      mergeMap(tokenInfo => this.getUserByToken()),
      tap(user => {
        if (user) {
          this.store.dispatch(appUserActions.loadUserSuccess({ user }));
        }
      }),
    );
  }

  public auth(username: string, password: string): Observable<TokenInfo> {
    return this._http.post<TokenInfo>(`${this._config.apiUrlWithPostfix}/token`, {
      username,
      password,
      grant_type: 'password',
    });
  }

  public refresh(refreshToken: string): Observable<TokenInfo> {
    return this._http.post<TokenInfo>(`${this._config.apiUrlWithPostfix}/token`, {
      refresh_token: refreshToken,
      grant_type: 'refresh_token',
    });
  }

  public loadFromLocalStorage(): Observable<User | null> {
    const tokenRaw = localStorage.getItem(TOKEN_LOCAL_STORAGE_KEY);
    let token: TokenInfo;
    try {
      token = JSON.parse(tokenRaw || '');
    } catch {
      this.clearLocalStorage();
      return of(null);
    }
    if (!token?.access_token) {
      return of(null);
    }
    return this._http.get<User>(`${this._config.apiUrlWithPostfix}/user`,
      { headers: { Authorization: `Bearer ${token.access_token}` } }).pipe(
        catchError(() => of(null)),
        tap(user => {
          if (user) {
            this.store.dispatch(appUserActions.loadUserSuccess({ user }));
          }
        })
      );
    // return this.refresh(token.refresh_token).pipe(
    //   catchError(() => of(null)),
    //   tap(tokenInfo => {
    //     if (tokenInfo) {
    //       this.store.dispatch(appUserActions.updateTokenInfo({ tokenInfo }));
    //       this.setRefreshSchedule(tokenInfo);
    //       this.saveInLocalStorage(tokenInfo);
    //     }
    //   }),
    //   mergeMap(tokenInfo => {
    //     if (tokenInfo) {
    //       return this.getUserByToken(tokenInfo.access_token).pipe(
    //         catchError(() => of(null)),
    //       );
    //     }
    //     return of(null);
    //   }),
    //   tap(user => {
    //     if (user) {
    //       this.store.dispatch(appUserActions.loadUserSuccess({ user }));
    //     }
    //   }),
    // );
  }

  public saveInLocalStorage(tokenInfo: TokenInfo): void {
    localStorage.setItem(TOKEN_LOCAL_STORAGE_KEY, JSON.stringify(tokenInfo));
  }

  public clearLocalStorage(): void {
    localStorage.removeItem(TOKEN_LOCAL_STORAGE_KEY);
  }

  public clearRefreshSchedule(): void {
    if (this.timerId) {
      clearTimeout(this.timerId);
    }
  }

  public setRefreshSchedule(tokenInfo: TokenInfo): void {
    if (tokenInfo.expiry) {
      this.timerId = setTimeout(() => {
        this.store.dispatch(appUserActions.refresh());
      }, tokenInfo.expiry * 1000 - EXPIRED_IN_DISTINCTION) as unknown as number;
    }
  }

  public getUserByToken(): Observable<User> {
    return this._http.get<User>(`${this._config.apiUrlWithPostfix}/user`);
  }
}
