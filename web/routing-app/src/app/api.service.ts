import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

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

  getTaxes() {
    return this.http.get(this.API_URL + `/taxes/`);
  }

}
