import { Component, OnInit } from '@angular/core';
import { faRubleSign, faAdd, faHotTub } from '@fortawesome/free-solid-svg-icons';
import { Subject, switchMap, tap } from 'rxjs';
import { ApiService } from '../api.service';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Chart } from 'angular-highcharts';

@Component({
  selector: 'app-records',
  templateUrl: './records.component.html',
  styleUrls: ['./records.component.css']
})

export class RecordsComponent implements OnInit {
  price_icon = faRubleSign;
  add_icon = faAdd;
  hot_icon = faHotTub;
  records: any;
  chart: any;
  refresh$ = new Subject<void>();
  form: FormGroup;

  constructor(
    private apiService: ApiService,
    private fb: FormBuilder,
  ) {
    this.form = this.fb.group({
      hot_value: ['', Validators.required],
      cold_value: ['', Validators.required],
      energy_value: ['', Validators.required],
    });
  }

  getRecords() {
    this.apiService.getRecords().subscribe(data => {
      this.records = data;
      let hot_values = []
      let cold_values = []
      let energy_values = []
      let drenage_values = []
      let months = []
      let error = ""

      if (this.records.error) {
        error = this.records.error
      }

      for (let record of this.records.data) {
        months.push(record.CreatedAt);
      }

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
          type: 'column'
        },
        title: {
          text: 'Аналитика потребления по счетчикам'
        },
        xAxis: {
          categories: months
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
    });
  }

  postRecord() {
    const val = this.form.value;
    if (val.hot_value && val.cold_value && val.energy_value) {
      this.apiService.createRecord(val.hot_value, val.cold_value, val.energy_value).pipe(tap(() => this.refresh$.next())).subscribe()
    }
  }

  deleteRecord(id: number) {
    this.apiService.deleteRecord(id).pipe(tap(() => this.refresh$.next())).subscribe()
  };

  ngOnInit(): void {
    this.refresh$.pipe(
      switchMap(() => this.apiService.getRecords().pipe(
        tap(() => this.getRecords())
      )
      )).subscribe();
    this.refresh$.next()
  }

}
