import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { faRubleSign, faAdd } from '@fortawesome/free-solid-svg-icons';
import { ApiService } from '../api.service';
import { Chart } from 'angular-highcharts';

@Component({
  selector: 'app-first',
  templateUrl: './records.component.html',
  styleUrls: ['./records.component.css']
})
export class RecordsComponent implements OnInit {
  price_icon = faRubleSign;
  add_icon = faAdd;
  records: any;
  chart: any;

  constructor(private apiService: ApiService) { }

  getChart() {

  }

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
      
      let hot_values = []
      let cold_values = []
      let energy_values = []
      let drenage_values = []

      for (let record of this.records.data) {
        hot_values.push(record.HotValue);
      }

      for (let record of this.records.data) {
        cold_values.push(record.ColdValue);
      }

      for (let record of this.records.data) {
        energy_values.push(record.EnergyValue);
      }

      for (let record of this.records.data) {
        drenage_values.push(record.DrenageValue);
      }

      this.chart = new Chart({
        chart: {
          type: 'line'
        },
        title: {
          text: 'Аналитика потребления по счетчикам'
        },
        credits: {
          enabled: false
        },
        series: [
          {
            name: 'Горячая вода',
            type: 'line',
            data: hot_values,
          },
          {
            name: 'Холодная вода',
            type: 'line',
            data: cold_values,
          },
          {
            name: 'Электричество',
            type: 'line',
            data: energy_values,
          },
          {
            name: 'Дренаж',
            type: 'line',
            data: drenage_values,
          },
        ]
      });
      
      console.log(this.records)

    });

  }

}
