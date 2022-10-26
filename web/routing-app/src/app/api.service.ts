import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable()
export class ApiService {

  constructor(private http: HttpClient) { }

  getRecords() {
    return this.http.get<Records>(`http://localhost/api/records/`);
  }

}

export interface Records {
  Id: Uint16Array
  Name: string;
  Value: Float32Array;
  Cost: Float32Array;
}