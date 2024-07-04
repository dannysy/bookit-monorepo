import { createActionGroup, emptyProps, props } from '@ngrx/store';
import { User } from '@core/models/user';
import { TokenInfo } from '@core/models/token-info';

export const appUserActions = createActionGroup({
  source: 'App User',
  events: {
    'Login Success': props<{ tokenInfo: TokenInfo }>(),
    'Login Failure': props<{ error: string }>(),
    refresh: emptyProps(),
    updateTokenInfo: props<{ tokenInfo: TokenInfo }>(),
    'Load User Success': props<{ user: User }>(),
    'Load User Failure': props<{ error: string }>(),
    logout: emptyProps(),
  },
});
