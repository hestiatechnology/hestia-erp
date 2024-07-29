import { Routes } from '@angular/router';
import { LoginComponent } from './auth/login/login.component';
import { TestComponent } from './test/test.component';
import { ForgotPasswordComponent } from "./auth/forgot-password/forgot-password.component";
import { SsoComponent } from "./auth/sso/sso.component";
import { HomeComponent as TextileHomeComponent } from "./textile/home/home.component";
import { HomeComponent as TechPackHomeComponent } from "./textile/techpack/home/home.component";
import { ViewComponent as TechPackViewComponent } from "./textile/techpack/view/view.component";
import { TextileComponent } from "./textile/textile.component";
import { InvoiceComponent } from "./invoice/invoice.component";
import { HomeComponent as InvoiceHomeComponent } from "./invoice/home/home.component";

export const routes: Routes = [
  { "path": "login", component: LoginComponent },
  { "path": "sso", component: SsoComponent },
  { "path": "forgotpassword", component: ForgotPasswordComponent },

  // /invoice
  {
    path: "invoice",
    component: InvoiceComponent,
    children: [
      { "path": "", component: InvoiceHomeComponent },
    ]
  },

  // /textile
  {
    "path": "textile",
    component: TextileComponent,
    children: [
      { "path": "", component: TextileHomeComponent },
      { "path": "techpack", component: TechPackHomeComponent },
      { "path": "techpack/new", component: TechPackViewComponent }
    ]
  },
  { "path": "test", component: TestComponent },
];
