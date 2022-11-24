import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError, Observable } from 'rxjs';

@Injectable()
export class ApiService {

  API_URL = 'http://localhost:8080/api';
  constructor(private http: HttpClient) { }

  getRecords() {
    // const headers = new HttpHeaders({
    //   'Access-Control-Allow-Origin': '*',
    // })
    return this.http.get(this.API_URL + `/records/`);
  }

  // createRecord(record: Record): Observable<Record> {
  createRecord() {
    return this.http.post<Record>(this.API_URL + `/records/`, {HotValue: 1, ColdValue: 1, EnergyValue: 1, DrenageValue: 1})
  }

  deleteRecord(id: number) {
    return this.http.delete(this.API_URL + `/records/` + id)
  }

  getTaxes() {
    return this.http.get(this.API_URL + `/taxes/`);
  }

}

export interface Record {
  Id: BigInt;
  HotValue: string;
  ColdValue: string;
  EnergyValue: string;
  DrenageValue: string;
}