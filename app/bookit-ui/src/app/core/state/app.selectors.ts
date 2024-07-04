import { AppState } from './app.state';
import { User } from '@core/models/user';
import { TokenInfo } from '@core/models/token-info';

export const selectCurrentUser = (state: { app: AppState }): User | null => state.app.currentUser;

export const selectTokenInfo = (state: { app: AppState }): TokenInfo | null => state.app.tokenInfo;
