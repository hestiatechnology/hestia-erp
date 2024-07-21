import { Component, OnInit } from '@angular/core';
import { AuthComponent } from "../auth.component";
import { FormControl, ReactiveFormsModule, Validators } from "@angular/forms";
import { merge } from "rxjs";
import { takeUntilDestroyed } from "@angular/core/rxjs-interop";
import { MatError, MatFormFieldModule } from "@angular/material/form-field";
import { MatButtonModule } from "@angular/material/button";
import { MatInputModule } from "@angular/material/input";
import { ActivatedRoute } from "@angular/router";

@Component({
  selector: 'app-sso',
  standalone: true,
  imports: [AuthComponent, MatError, MatFormFieldModule, ReactiveFormsModule, MatButtonModule, MatInputModule],
  templateUrl: './sso.component.html',
  styleUrl: './sso.component.scss'
})
export class SsoComponent implements OnInit {
  email = new FormControl('', [Validators.required, Validators.email]);
  emailErrorMessage = '';
  buttonDisabled = true;
  hide = true;

  constructor(private route: ActivatedRoute) {
    merge(this.email.statusChanges, this.email.valueChanges)
      .pipe(takeUntilDestroyed())
      .subscribe(() => {
        this.updateEmailErrorMessage()
        this.updateButtonStatus();
      });
  }

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      if (params["email"]) {
        this.email.setValue(params["email"]);
        // Check if the email is valid and if so call the submit method
        if (this.email.valid) {
          this.submit();
        }
      }
    });
  }

  updateEmailErrorMessage() {
    if (this.email.hasError('required')) {
      this.emailErrorMessage = 'You must enter a value';
    } else if (this.email.hasError('email')) {
      this.emailErrorMessage = 'Not a valid email';
    } else {
      this.emailErrorMessage = '';
    }
  }

  updateButtonStatus() {
    this.buttonDisabled = !(this.email.valid);
  }

  submit() {
    // This is where the SSO magic would
    // happen. For now, we'll just log the email.
    console.log("SSO with email: ", this.email.value);
  }

}
