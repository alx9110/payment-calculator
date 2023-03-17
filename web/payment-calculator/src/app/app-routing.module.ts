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
  { path: 'login', pathMatch: 'full', component: LoginComponent },
  { path: 'createUser', pathMatch: 'full', component: CreateUserComponent },
  { path: 'records', pathMatch: 'full', component: RecordsComponent, canActivate: [CanActivateViaAuthGuard]},
  { path: 'taxes', pathMatch: 'full', component: TaxesComponent, canActivate: [CanActivateViaAuthGuard]},
  { path: 'home', pathMatch: 'full', component: HomeComponent },
  { path: '**', redirectTo: '' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
