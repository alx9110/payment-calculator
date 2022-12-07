import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RecordsComponent } from './records/records.component';
import { TaxesComponent } from './taxes/taxes.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { CanActivateViaAuthGuard } from './auth.service';
import { CreateUserComponent } from './create-user/create-user.component';

const routes: Routes = [
  { path: '',   redirectTo: '/home', pathMatch: 'full' },
  { path: 'login', component: LoginComponent },
  { path: 'createUser', component: CreateUserComponent },
  { path: 'records', component: RecordsComponent, canActivate: [CanActivateViaAuthGuard]},
  { path: 'taxes', component: TaxesComponent, canActivate: [CanActivateViaAuthGuard]},
  { path: 'home', component: HomeComponent },
  { path: '**', redirectTo: '' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
