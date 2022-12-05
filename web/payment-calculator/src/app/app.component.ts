import { Component } from '@angular/core';
import { Chart } from 'angular-highcharts';
import { faCalculator } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from './auth.service';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  constructor(public authService: AuthService) { }
  title = 'routing-app';
  main_icon = faCalculator;
  chart = Chart;
}
