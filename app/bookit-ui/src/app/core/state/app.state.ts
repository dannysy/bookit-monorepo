import { createReducer, on } from '@ngrx/store';
import { User } from '@core/models/user';
import { TokenInfo } from '@core/models/token-info';
import { appUserActions } from './actions/app-user.actions';

export interface AppState {
  tokenInfo: TokenInfo | null;
  currentUser: User | null;
}

const initialState: AppState = {
  tokenInfo:null,
  currentUser: null,
};

export const appReducer = createReducer(
  initialState,
  on(appUserActions.loginSuccess, (state, { tokenInfo }): AppState => ({
    ...state,
    tokenInfo,
  })),
  on(appUserActions.loadUserSuccess, (state, { user }): AppState => ({
    ...state,
    currentUser: user,
  })),
  on(appUserActions.updateTokenInfo, (state, { tokenInfo }): AppState => ({
    ...state,
    tokenInfo,
  })),
  on(appUserActions.logout, (state): AppState => ({
    ...state,
    tokenInfo: null,
    currentUser: null,
  })),

);
