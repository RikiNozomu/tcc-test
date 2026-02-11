import { Routes } from '@angular/router';
import { LoginPage } from './pages/login/login';
import { MePage } from './pages/me/me';
import { SignUpPage } from './pages/sign-up/sign-up';

export const routes: Routes = [
    { path: 'login', component: LoginPage },
    { path: 'sign-up', component: SignUpPage },
    { path: 'me', component: MePage },
    { path: '**', redirectTo: '/login' },
];
