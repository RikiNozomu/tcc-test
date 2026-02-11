import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { userResponse } from '../../types';
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root',
})
export class User {
  constructor(private http: HttpClient, private cookieService: CookieService) { }

  getMe() {
    // Implement sign-up logic here, e.g., call an API endpoint
    const token = this.cookieService.get('token');
    return this.http.get<userResponse>(environment.NG_APP_API_URL + '/user/me', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
  }
}
