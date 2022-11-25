import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RecordsComponent } from './records/records.component';
import { TaxesComponent } from './taxes/taxes.component';
import { HomeComponent } from './home/home.component';

const routes: Routes = [
  { path: '',   redirectTo: '/home-component', pathMatch: 'full' },
  { path: 'records-component', component: RecordsComponent },
  { path: 'taxes-component', component: TaxesComponent },
  { path: 'home-component', component: HomeComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
