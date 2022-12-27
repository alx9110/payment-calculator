import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';


@Injectable()
export class ApiService {

  API_URL = 'http://localhost:8080/api';
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
  };

  getTaxes() {
    return this.http.get(this.API_URL + `/taxes/`);
  }

}

export interface Record {
  HotValue: string;
  ColdValue: string;
  EnergyValue: string;
}
