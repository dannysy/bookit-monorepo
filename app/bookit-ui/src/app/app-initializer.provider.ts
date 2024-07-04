import { APP_INITIALIZER, Provider } from '@angular/core';
import { ConfigService } from '@core/services/config.service';
import { forkJoin, mergeMap } from 'rxjs';
import { UserService } from "@core/services/user.service";

export const AppInitializerProvider: Provider = {
  provide: APP_INITIALIZER,
  useFactory: (
    configService: ConfigService,
    userService: UserService,
  ) => () => forkJoin([
    configService.load(),
  ]).pipe(
    mergeMap(() => userService.loadFromLocalStorage()),
  ),
  deps: [ConfigService, UserService],
  multi: true,
};
