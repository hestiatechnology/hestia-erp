import { Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { TestComponent } from './test/test.component';
import { ForgotPasswordComponent } from "./forgot-password/forgot-password.component";

export const routes: Routes = [
  {"path": "login", component: LoginComponent},
  {"path": "forgotpassword", component: ForgotPasswordComponent},
  {"path": "test", component: TestComponent},
];
