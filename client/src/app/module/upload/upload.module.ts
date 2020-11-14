import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UploadRoutingModule } from './upload-routing.module';
import { UploadComponent } from './page/upload/upload.component';
import { ViewComponent } from './page/view/view.component';
import { ReactiveFormsModule } from '@angular/forms';
import { FileLoaderComponent } from './component/file-loader/file-loader.component';


@NgModule({
  declarations: [UploadComponent, ViewComponent, FileLoaderComponent],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    UploadRoutingModule
  ]
})
export class UploadModule { }
