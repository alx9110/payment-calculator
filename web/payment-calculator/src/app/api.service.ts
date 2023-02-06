import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { environment } from 'src/environments/environment';


@Injectable()
export class ApiService {
  API_URL = environment.apiUrl+'/api';
  constructor(
    private http: HttpClient,
    private router: Router,
  ) { }

  getRecords() {
    return this.http.get(this.API_URL + `/records/`);
  }

  createRecord(hot_value: number, cold_value: number, energy_value: number) {
    return this.http.post<Record>(this.API_URL + `/records/`, {HotValue: hot_value, ColdValue: cold_value, EnergyValue: energy_value})
  }

  deleteRecord(id: number) {
    return this.http.delete(this.API_URL + '/records/' + id)
  }

  getTaxes() {
    return this.http.get(this.API_URL + `/taxes/`);
  }

  createTax(hot_price: number, cold_price: number, energy_price: number, drenage_price: number) {
    return this.http.post<Tax>(this.API_URL + `/taxes/`, {HotPrice: hot_price, ColdPrice: cold_price, EnergyPrice: energy_price, DrenagePrice: drenage_price})
  }

  deleteTax(id: number) {
    return this.http.delete(this.API_URL + '/taxes/' + id)
  }

}

export interface Record {
  HotValue: string;
  ColdValue: string;
  EnergyValue: string;
}

export interface Tax {
  HotPrice: string;
  ColdPrice: string;
  EnergyPrice: string;
  DrenagePrice: string;
}
