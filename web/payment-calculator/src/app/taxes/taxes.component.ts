import { Component, OnInit } from '@angular/core';
import { Subject, switchMap, tap } from 'rxjs';
import { faAdd } from '@fortawesome/free-solid-svg-icons';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-taxes',
  templateUrl: './taxes.component.html',
  styleUrls: ['./taxes.component.css']
})
export class TaxesComponent implements OnInit {

  add_icon = faAdd;
  taxes: any;
  form: FormGroup;
  refresh$ = new Subject<void>();

  constructor(
    private apiService: ApiService,
    private fb: FormBuilder,
  ) { 
    this.form = this.fb.group({
      hot_price: ['', Validators.required],
      cold_price: ['', Validators.required],
      energy_price: ['', Validators.required],
      drenage_price: ['', Validators.required],
    });
  }

  getTaxes() {
    this.apiService.getTaxes().subscribe(data => {
      this.taxes= data;
    });
  }

  postTax() {
    const val = this.form.value;
    if (val.hot_price && val.cold_price && val.energy_price) {
      this.apiService.createTax(val.hot_price, val.cold_price, val.energy_price, val.drenage_price).pipe(tap(() => this.refresh$.next())).subscribe()
    }
  }

  deleteTax(id: number) {
    this.apiService.deleteTax(id).pipe(tap(() => this.refresh$.next())).subscribe()
  };

  ngOnInit(): void {
    this.refresh$.pipe(
      switchMap(() => this.apiService.getTaxes().pipe(
        tap(() => this.getTaxes())
      )
      )).subscribe();
    this.refresh$.next()
  }

}