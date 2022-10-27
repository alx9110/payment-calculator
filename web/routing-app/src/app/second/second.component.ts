import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-second',
  templateUrl: './second.component.html',
  styleUrls: ['./second.component.css']
})
export class SecondComponent implements OnInit {

  taxes: any;
  constructor(private apiService: ApiService) { }

  ngOnInit(): void {
    this.apiService.getTaxes().subscribe(data => {
      this.taxes = data;
    });
  }

}