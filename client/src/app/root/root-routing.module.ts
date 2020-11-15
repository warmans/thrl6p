import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: 'patch',
    loadChildren: () => import('../module/upload/upload.module').then(m => m.UploadModule),
  },
  {
    path: '',
    redirectTo: 'patch/upload',
    pathMatch: 'full'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class RootRoutingModule {
}
