import { Component, OnInit, signal } from '@angular/core';
import {
  AbstractControl,
  ReactiveFormsModule,
  ValidationErrors,
  ValidatorFn,
} from '@angular/forms';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { ButtonModule } from 'primeng/button';
import { InputTextModule } from 'primeng/inputtext';
import { PasswordModule } from 'primeng/password';
import { AuthService } from '../../services/auth/auth';
import { Router } from '@angular/router';

@Component({
  selector: 'app-sign-up',
  imports: [ReactiveFormsModule, InputTextModule, PasswordModule, ButtonModule],
  templateUrl: './sign-up.html',
  styleUrl: './sign-up.css',
})
export class SignUpPage implements OnInit {
  constructor(private authService: AuthService, private router: Router) {}

  passwordNotMatchValidator: ValidatorFn = (control: AbstractControl): ValidationErrors | null => {
    const password = control.get('password');
    const confirmPassword = control.get('confirmPassword');
    return password && confirmPassword && password.value === confirmPassword.value ? null : { passwordNotMatch: true };
  };

  signUpForm = new FormGroup(
    {
      username: new FormControl('', [
        Validators.required,
        Validators.minLength(6),
        Validators.maxLength(12),
      ]),
      password: new FormControl('', [
        Validators.required,
        Validators.minLength(6),
        Validators.maxLength(12),
      ]),
      confirmPassword: new FormControl('', [
        Validators.required,
        Validators.minLength(6),
        Validators.maxLength(12),
      ]),
    },
    { validators: this.passwordNotMatchValidator },
  );

  onSignUp() {
    if(this.signUpForm.invalid){
      
      if(this.signUpForm.hasError('required', 'username') || this.signUpForm.hasError('required', 'password') || this.signUpForm.hasError('required', 'confirmPassword')){
        alert('All fields are required!');
      } else if(this.signUpForm.hasError('minlength', 'username') || this.signUpForm.hasError('minlength', 'password') || this.signUpForm.hasError('minlength', 'confirmPassword')){
        alert('Fields must be at least 6 characters long!');
      } else if(this.signUpForm.hasError('maxlength', 'username') || this.signUpForm.hasError('maxlength', 'password') || this.signUpForm.hasError('maxlength', 'confirmPassword')){
        alert('Fields cannot exceed 12 characters!');
      } else if(this.signUpForm.hasError('passwordNotMatch')){
        alert('Passwords do not match!');
      }
      return;
    }
    
    // Implement sign-up logic here, e.g., call an API endpoint
    this.authService.signUp(this.signUpForm.value.username || '', this.signUpForm.value.password || '').subscribe({
      next: (response) => {
        // Handle successful sign-up, e.g., store token, navigate to 'me' page, etc.
        console.log('Sign-up successful:', response);
        alert('Sign-up successful! You can now log in.');
        this.router.navigate(['/login']);
      },
      error: (error) => {
        // Handle sign-up error
        console.error('Sign-up error:', error);
        alert('Sign-up failed. Please try again.');
      },
    });
  }

  ngOnInit(): void {
    console.log('SignUpPage initialized');
    if (this.authService.isAuthenticated()) {
      this.router.navigate(['/me']);
    }
  }
}
