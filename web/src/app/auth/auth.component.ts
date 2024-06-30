import { Component } from '@angular/core';
import { MatIconModule } from "@angular/material/icon";
import { Location } from '@angular/common';
@Component({
  selector: 'app-auth',
  standalone: true,
  imports: [MatIconModule],
  templateUrl: './auth.component.html',
  styleUrl: './auth.component.scss'
})

export class AuthComponent {
  constructor(private location: Location) { }

  goBack() {
    this.location.back();
  }
}
