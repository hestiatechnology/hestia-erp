import { Component } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { ReactiveFormsModule, FormControl, Validators } from "@angular/forms";
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { merge } from 'rxjs';
import { AuthComponent } from "../auth.component";
import { RouterModule } from "@angular/router";
import { MatSnackBar } from '@angular/material/snack-bar';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [AuthComponent, MatButtonModule, MatInputModule, MatFormFieldModule, MatIconModule, ReactiveFormsModule, RouterModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss'
})
export class LoginComponent {
  email = new FormControl('', [Validators.required, Validators.email]);
  password = new FormControl('', [Validators.required]);
  emailErrorMessage = '';
  passwordErrorMessage = '';
  buttonDisabled = true;
  hide = true;

  constructor(private _snackBar: MatSnackBar, private route: ActivatedRoute) {
    merge(this.email.statusChanges, this.email.valueChanges)
      .pipe(takeUntilDestroyed())
      .subscribe(() => {
        this.updateEmailErrorMessage()
        this.updateButtonStatus();
      });
    merge(this.password.statusChanges, this.password.valueChanges)
      .pipe(takeUntilDestroyed())
      .subscribe(() => {
        this.updatePasswordErrorMessage();
        this.updateButtonStatus();
      });
    // Check if the URL Param expired is true
    if (this.route.snapshot.queryParams["expired"]) {
      this._snackBar.open($localize`A sua sessão expirou. Por favor faça login novamente`, $localize`Ignorar`, {
        duration: 10000,
      });
    }

  }

  updateEmailErrorMessage() {
    if (this.email.hasError('required')) {
      this.emailErrorMessage = $localize`Tem de introduzir um valor`;

    } else if (this.email.hasError('email')) {
      this.emailErrorMessage = $localize`Não é um email válido`;
    } else {
      this.emailErrorMessage = '';
    }
  }

  updatePasswordErrorMessage() {
    if (this.password.hasError('required')) {
      this.passwordErrorMessage = $localize`Tem de introduzir um valor`;
    } else {
      this.passwordErrorMessage = '';
    }
  }

  updateButtonStatus() {
    this.buttonDisabled = !(this.email.valid && this.password.valid);
  }
}
