import { Routes } from '@angular/router';
import { LoginComponent } from './auth/login/login.component';
import { TestComponent } from './test/test.component';
import { ForgotPasswordComponent } from "./auth/forgot-password/forgot-password.component";
import { SsoComponent } from "./auth/sso/sso.component";
import { HomeComponent as TextileHomeComponent } from "./textile/home/home.component";
import { HomeComponent as TechPackHomeComponent } from "./textile/techpack/home/home.component";
import { ViewComponent as TechPackViewComponent } from "./textile/techpack/view/view.component";
import { TextileComponent } from "./textile/textile.component";
import { InvoicingComponent } from "./invoicing/invoicing.component";
import { InvoicingHomeComponent } from "./invoicing/home/home.component";
import { FaturaHomeComponent } from "./invoicing/ft/home/home.component";
import { FaturaSimplificadaHomeComponent } from "./invoicing/fs/home/home.component";
import { FaturaReciboHomeComponent } from "./invoicing/fr/home/home.component";

export const routes: Routes = [
  { path: "login", component: LoginComponent },
  { path: "sso", component: SsoComponent },
  { path: "forgotpassword", component: ForgotPasswordComponent },

  // /invoicing
  {
    path: "invoicing",
    component: InvoicingComponent,
    children: [
      { path: "", component: InvoicingHomeComponent },
      // FT -  Faturas (Invoices)
      { path: "ft", component: FaturaHomeComponent },
      { path: "ft/:id", loadComponent: () => import('./invoicing/ft/document/document.component').then(m => m.FaturaDocumentComponent) },
      // FS - Faturas Simplificadas (Simplified Invoices)
      { path: "fs", component: FaturaSimplificadaHomeComponent },
      { path: "fs/:id", loadComponent: () => import('./invoicing/fs/document/document.component').then(m => m.FaturaSimplificadaDocumentComponent) },
      // FR - Faturas Recibo (Receipt Invoices)
      { path: "fr", component: FaturaReciboHomeComponent },
      { path: "fr/:id", loadComponent: () => import('./invoicing/fr/document/document.component').then(m => m.FaturaReciboDocumentComponent) }
    ]
  },

  // /textile
  {
    path: "textile",
    component: TextileComponent,
    children: [
      { path: "", component: TextileHomeComponent },
      { path: "techpack", component: TechPackHomeComponent },
      { path: "techpack/new", component: TechPackViewComponent }
    ]
  },
  { path: "test", component: TestComponent },
];
