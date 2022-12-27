import { Component, OnInit } from '@angular/core';
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

  ngOnInit(): void {
    this.apiService.getTaxes().subscribe(data => this.taxes = data);
  }

}