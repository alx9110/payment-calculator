import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-create-user',
  templateUrl: './create-user.component.html',
  styleUrls: ['./create-user.component.css']
})
export class CreateUserComponent implements OnInit {

  form: FormGroup;

  constructor(
    private fb: FormBuilder,
    private authService: AuthService,
    ) {
    this.form = this.fb.group({
        email: ['', Validators.required],
        password: ['', Validators.required],
        botid: [''],
    });
  }

  ngOnInit(): void {
  }

  login() {
    const val = this.form.value;
    if (val.email && val.password) {
      this.authService.login(val.email, val.password);
    }
  }

  createUser() {
    const val = this.form.value;
    if (val.email && val.password) {
      this.authService.createUser(val.email, val.password, val.botid);
    }
  }

}
