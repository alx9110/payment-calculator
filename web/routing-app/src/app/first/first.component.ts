import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-first',
  templateUrl: './first.component.html',
  styleUrls: ['./first.component.css']
})
export class FirstComponent implements OnInit {

  records: any;
  constructor(private apiService: ApiService) { }

  ngOnInit(): void {
    this.apiService.getRecords().subscribe(data => {
      this.records = data;
    });
  }

}
