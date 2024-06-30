import { Component } from '@angular/core';
import { AuthComponent } from "../auth/auth.component";
import { FormControl, ReactiveFormsModule, Validators } from '@angular/forms';
import { merge } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { MatFormFieldModule } from "@angular/material/form-field";
import { MatButtonModule } from "@angular/material/button";
import { MatInputModule } from "@angular/material/input";
import { RouterModule } from "@angular/router";
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-forgot-password',
  standalone: true,
  imports: [AuthComponent, MatButtonModule, MatInputModule, MatFormFieldModule, ReactiveFormsModule, RouterModule, MatProgressBarModule],
  templateUrl: './forgot-password.component.html',
  styleUrl: './forgot-password.component.scss'
})
export class ForgotPasswordComponent {
  email = new FormControl('', [Validators.required, Validators.email]);
  emailErrorMessage = '';
  buttonDisabled = true;
  hide = true;
  progressBarVisible = false;

   constructor(private _snackBar: MatSnackBar) {
    merge(this.email.statusChanges, this.email.valueChanges)
      .pipe(takeUntilDestroyed())
      .subscribe(() => {
        this.updateEmailErrorMessage()
        this.updateButtonStatus();
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
    this.progressBarVisible = !this.progressBarVisible;
    setTimeout(() => {
      this.progressBarVisible = !this.progressBarVisible;
    }, 3000);
    this._snackBar.open("An email with instructions to reset your password will be sent to you if an account with the provided email exists.", "Close");

  }


}
