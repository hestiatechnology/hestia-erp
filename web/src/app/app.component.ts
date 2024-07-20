import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { Router } from '@angular/router';
@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit {
  constructor(private router: Router) {
  }

  ngOnInit(): void {
    if (Date.now() >= parseInt(localStorage.getItem("expiry_time") || '')) {
      // Check if token is expired
      this.router.navigate(['login'], { queryParams: { returnUrl: this.router.url, expired: true } });
    } else if (localStorage.getItem("token") === null) {
      this.router.navigate(['login'], { queryParams: { returnUrl: this.router.url } });
    }
  }
}
