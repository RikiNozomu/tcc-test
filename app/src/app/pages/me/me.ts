import { Component, OnDestroy, OnInit } from '@angular/core';
import { User } from '../../services/user/user';
import { Router } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { AuthService } from '../../services/auth/auth';
import { Observable } from 'rxjs/internal/Observable';
import { userResponse } from '../../types';
import { AsyncPipe } from '@angular/common';
import { Subscription } from 'rxjs/internal/Subscription';

@Component({
  selector: 'app-me',
  imports: [ButtonModule, AsyncPipe],
  templateUrl: './me.html',
  styleUrl: './me.css',
})
export class MePage implements OnInit, OnDestroy {
  user$!: Observable<userResponse>;
  private subscription: Subscription | undefined;

  constructor(
    private userService: User,
    private authService: AuthService,
  ) {}

  getMe() {
    this.user$ = this.userService.getMe().pipe()
    this.subscription = this.user$.subscribe({
      next: (response) => {
        console.log('User data fetched:', response);  
      },
      error: (error) => {
        console.error('Error fetching user data:', error);
        if (error.status === 401) {
          this.authService.signOut();
        }
      },
    });
  }

  ngOnInit() {
    this.getMe();
  }

  ngOnDestroy() {
    this.subscription?.unsubscribe();
  }
}
