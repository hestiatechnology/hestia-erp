import { Routes } from '@angular/router';
import { LoginComponent } from './auth/login/login.component';
import { TestComponent } from './test/test.component';
import { ForgotPasswordComponent } from "./auth/forgot-password/forgot-password.component";
import { SsoComponent } from "./auth/sso/sso.component";
import { HomeComponent } from "./textile/home/home.component";
import { NavbarComponent as TextileNavbarComponent } from "./textile/navbar/navbar.component";

export const routes: Routes = [
  { "path": "login", component: LoginComponent },
  { "path": "sso", component: SsoComponent },
  { "path": "forgotpassword", component: ForgotPasswordComponent },
  {
    "path": "textile",
    component: TextileNavbarComponent,
    children: [
      { "path": "", component: HomeComponent }
    ]
  },
  { "path": "test", component: TestComponent },
];
