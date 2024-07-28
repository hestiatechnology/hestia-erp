import { Routes } from '@angular/router';
import { LoginComponent } from './auth/login/login.component';
import { TestComponent } from './test/test.component';
import { ForgotPasswordComponent } from "./auth/forgot-password/forgot-password.component";
import { SsoComponent } from "./auth/sso/sso.component";
import { HomeComponent } from "./textile/home/home.component";
import { HomeComponent as TechPackHomeComponent } from "./textile/techpack/home/home.component";
import { ViewComponent as TechPackViewComponent } from "./textile/techpack/view/view.component";
import { TextileComponent } from "./textile/textile.component";

export const routes: Routes = [
  { "path": "login", component: LoginComponent },
  { "path": "sso", component: SsoComponent },
  { "path": "forgotpassword", component: ForgotPasswordComponent },
  {
    "path": "textile",
    component: TextileComponent,
    children: [
      { "path": "", component: HomeComponent },
      { "path": "techpack", component: TechPackHomeComponent },
      { "path": "techpack/new", component: TechPackViewComponent }
    ]
  },
  { "path": "test", component: TestComponent },
];
