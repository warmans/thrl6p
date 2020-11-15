import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UploadComponent } from './page/upload/upload.component';
import { ViewComponent } from './page/view/view.component';

const routes: Routes = [
  { path: 'upload', component: UploadComponent},
  { path: ':id', component: ViewComponent }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UploadRoutingModule {
}
