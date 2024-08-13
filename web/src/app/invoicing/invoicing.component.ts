import { Component } from '@angular/core';
import { NavbarComponent } from "../navbar/navbar.component";

@Component({
  selector: 'app-invoicing',
  standalone: true,
  imports: [NavbarComponent],
  templateUrl: './invoicing.component.html',
  styleUrl: './invoicing.component.scss'
})
export class InvoicingComponent {

}
