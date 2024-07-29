import { Component } from '@angular/core';
import { NavbarComponent } from "../navbar/navbar.component";

@Component({
  selector: 'app-textile',
  standalone: true,
  imports: [NavbarComponent],
  templateUrl: './textile.component.html',
  styleUrl: './textile.component.scss'
})
export class TextileComponent {
  links = [{ name: "Techpacks", path: "techpack" }];

}
