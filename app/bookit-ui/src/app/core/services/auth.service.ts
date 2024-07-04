import { Injectable } from '@angular/core';
import {ConfigService} from "@core/services/config.service";
import Sdk from "casdoor-js-sdk";

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private readonly sdk!: Sdk;

  constructor(private config: ConfigService) {
    this.sdk = new Sdk({
      appName: config.Config.appName,
      clientId: config.Config.clientId,
      organizationName: config.Config.organizationName,
      serverUrl: config.Config.authUrl,
      redirectPath: "/callback",
    })
  }

  public get Sdk(): Sdk {
    return this.sdk;
  }
}
