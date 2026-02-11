import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { loginResponse, userResponse } from '../../types';
import { environment } from '../../../environments/environment.development';
import { Router } from '@angular/router';
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  constructor(private http: HttpClient, private router: Router, private cookieService: CookieService) { }

  isAuthenticated(): boolean {
    const token = this.cookieService.get('token');
    return token !== null && token !== '';
  }

  login(username: string, password: string) {
    // Implement login logic here, e.g., call an API endpoint
    return this.http.post<loginResponse>(environment.NG_APP_API_URL + '/auth/login', {
      username,
      password,
    });
  }

  signUp(username: string, password: string) {
    // Implement sign-up logic here, e.g., call an API endpoint
    return this.http.post<userResponse>(environment.NG_APP_API_URL + '/user', {
      username,
      password,
    });
  }

  signOut() {
    this.cookieService.delete('token');
    this.router.navigate(['/login']);
  }
}
