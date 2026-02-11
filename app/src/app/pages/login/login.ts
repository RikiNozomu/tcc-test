import { Component, OnInit, signal } from '@angular/core';
import {
  FormControl,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { InputTextModule } from 'primeng/inputtext';
import { PasswordModule } from 'primeng/password';
import { ButtonModule } from 'primeng/button';
import { Router, RouterLink } from '@angular/router';
import { AuthService } from '../../services/auth/auth';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-login',
  imports: [
    ReactiveFormsModule,
    FormsModule,
    InputTextModule,
    PasswordModule,
    ButtonModule,
    RouterLink,
  ],
  templateUrl: './login.html',
  styleUrl: './login.css',
})
export class LoginPage implements OnInit {
  constructor(
    private authService: AuthService,
    private router: Router,
    private cookieService: CookieService
  ) {}

  loginForm = new FormGroup({
    username: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required]),
  });


  ngOnInit() {
    // If already authenticated, navigate to /me
    if (this.authService.isAuthenticated()) {
      this.router.navigate(['/me']);
    }
  }

  onLogin() {
    if (this.loginForm.invalid) {
      if (
        this.loginForm.hasError('required', 'username') ||
        this.loginForm.hasError('required', 'password')
      ) {
        alert('All fields are required!');
      }
      return;
    }

    // Implement login logic here, e.g., call an API endpoint
    this.authService
      .login(this.loginForm.value.username || '', this.loginForm.value.password || '')
      .subscribe({
        next: (response) => {
          this.cookieService.set('token', response.token);
          this.router.navigate(['/me']);
        },
        error: (error) => {
          // Handle login error
          console.error('Login error:', error);
          alert('Login failed. Please try again.');
        },
      });
  }
}
