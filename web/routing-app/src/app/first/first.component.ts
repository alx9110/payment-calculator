import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { faRubleSign } from '@fortawesome/free-solid-svg-icons';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-first',
  templateUrl: './first.component.html',
  styleUrls: ['./first.component.css']
})
export class FirstComponent implements OnInit {
  price_icon = faRubleSign;
  records: any;
  constructor(private apiService: ApiService) { }

  postRecord() {
    this.apiService.createRecord().subscribe(data => {
    });
  }

  deleteRecord(id: number) {
    this.apiService.deleteRecord(id).subscribe(data => {
    })
  };
  

  ngOnInit(): void {
    this.apiService.getRecords().subscribe(data => {
      this.records = data;
    });
  }

}
