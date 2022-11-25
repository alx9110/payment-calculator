import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { ChartModule } from 'angular-highcharts';

import { ApiService } from './api.service';
import { HomeComponent } from './home/home.component';
import { RecordsComponent } from './records/records.component';
import { TaxesComponent } from './taxes/taxes.component';

@NgModule({
  declarations: [
    AppComponent,
    RecordsComponent,
    TaxesComponent,
    HomeComponent,
    RecordsComponent,
    TaxesComponent,
  ],
  imports: [
    HttpClientModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    FontAwesomeModule,
    ChartModule,
  ],
  providers: [ApiService],
  bootstrap: [AppComponent]
})
export class AppModule { }
