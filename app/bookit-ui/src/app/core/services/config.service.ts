import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, tap } from 'rxjs';

type Config = {
  apiUrl: string;
  authUrl: string;
  clientId: string;
  organizationName: string;
  appName: string;
};

@Injectable({ providedIn: 'root' })
export class ConfigService {
  public get Config(): Config {
    return this.config;
  }

  public get apiUrlWithPostfix(): string {
    return `${this.config.apiUrl}/v1`;
  }

  private config!: Config;

  private configUrl = 'assets/config.json';

  public constructor(private httpClient: HttpClient) {}

  public load(): Observable<Config> {
    return this.httpClient.get<Config>(this.configUrl).pipe(
      tap((config: Config) => {
        this.config = config;
      }),
    );
  }
}
