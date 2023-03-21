import { Injectable, Injector } from '@angular/core';
import { HttpClient, HttpHeaders, HttpInterceptor, HttpRequest, HttpHandler } from '@angular/common/http';
import { Router, CanActivate } from '@angular/router';
import { environment } from 'src/environments/environment';


@Injectable()
export class AuthService {

    API_URL = environment.apiUrl+'/api';
    TOKEN_KEY = 'PAYCALC-JWT';
    LOGIN_KEY = 'PAYCALC-LOGIN';

    constructor(private http: HttpClient, private router: Router) { }

    get token() {
        return localStorage.getItem(this.TOKEN_KEY);
    }

    get isAuthenticated() {
        return !!localStorage.getItem(this.TOKEN_KEY);
    }

    logout() {
        localStorage.removeItem(this.TOKEN_KEY);
        localStorage.removeItem(this.LOGIN_KEY);
        this.router.navigateByUrl('/home');
    }

    login(email: string, pass: string) {
        const headers = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json', 'Cache-Control': 'no-cache' })
        };

        const data = {
            email: email,
            password: pass
        };

        this.http.post(this.API_URL + '/user/token', data, headers).subscribe(
            (res: any) => {
                localStorage.setItem(this.TOKEN_KEY, res.token);
                localStorage.setItem(this.LOGIN_KEY, res.login);
                this.router.navigateByUrl('/records');
            }
        );
    }

    createUser(email: string, pass: string, botid: string) {

        const data = {
            email: email,
            password: pass,
            botid: botid,
        };
        const headers = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json', 'Cache-Control': 'no-cache' })
        };
        this.http.post(this.API_URL + '/user/create', data).subscribe(
            (res: any) => {
                this.http.post(this.API_URL + '/user/token', data, headers).subscribe(
                    (res: any) => {
                        localStorage.setItem(this.TOKEN_KEY, res.token);
        
                        this.router.navigateByUrl('/records');
                    }
                );
            },
        );
    }
}

@Injectable()
export class AuthInterceptorService implements HttpInterceptor {

    constructor(private injector: Injector) { }

    intercept(req: HttpRequest<any>, next: HttpHandler) {
        const authService = this.injector.get(AuthService);
        const authRequest = req.clone({
            headers: req.headers.set('Authorization', '' + authService.token)
        });
        return next.handle(authRequest);
    }
}

@Injectable()
export class CanActivateViaAuthGuard implements CanActivate {
    constructor(private authService: AuthService) {

    }

    canActivate() {
        return this.authService.isAuthenticated;
    }
}