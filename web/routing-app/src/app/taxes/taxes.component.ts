import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-second',
  templateUrl: './taxes.component.html',
  styleUrls: ['./taxes.component.css']
})
export class TaxesComponent implements OnInit {

  taxes: any;
  constructor(private apiService: ApiService) { }

  ngOnInit(): void {
    this.apiService.getTaxes().subscribe(data => {
      this.taxes = data;
    });
  }

}