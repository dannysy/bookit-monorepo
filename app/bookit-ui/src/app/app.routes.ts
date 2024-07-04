import { Routes } from '@angular/router';
import {MainLayoutComponent} from "@core/layouts/main/main.component";
import {authCanActivate} from "@core/guards/auth.guard";
import {tokenCanActivate} from "@core/guards/token.guard";

export const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: '/app' },
  { path: 'callback', pathMatch: 'full', canActivate: [tokenCanActivate], component: MainLayoutComponent },
  {
    path: 'app',
    component: MainLayoutComponent,
    canActivate: [authCanActivate],
    children: [
      {
        path: 'workspace',
        loadChildren: () => import('./features/workspace/workspace.routes').then(m => m.WORKSPACE_ROUTES),
      }
    ]
  },
];
