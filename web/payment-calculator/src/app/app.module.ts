import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { ChartModule } from 'angular-highcharts';

import { ApiService } from './api.service';
import { HomeComponent } from './home/home.component';
import { RecordsComponent } from './records/records.component';
import { TaxesComponent } from './taxes/taxes.component';
import { LoginComponent } from './login/login.component';
import { AuthService, AuthInterceptorService, CanActivateViaAuthGuard } from './auth.service';
import { CreateUserComponent } from './create-user/create-user.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

@NgModule({
  declarations: [
    AppComponent,
    RecordsComponent,
    TaxesComponent,
    HomeComponent,
    RecordsComponent,
    TaxesComponent,
    LoginComponent,
    CreateUserComponent,
  ],
  imports: [
    HttpClientModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    FontAwesomeModule,
    ChartModule,
    FormsModule,
    ReactiveFormsModule,
    RouterModule,
    NgbModule,
  ],
  providers: [
    ApiService,
    AuthService,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptorService,
      multi: true
    },
    CanActivateViaAuthGuard
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}
