import { Component } from '@angular/core';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { faHouse, faRubleSign } from '@fortawesome/free-solid-svg-icons';

@Injectable()
export class DataService {
  constructor(private http: HttpClient) {}
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'routing-app';
  main_icon = faHouse;
}
