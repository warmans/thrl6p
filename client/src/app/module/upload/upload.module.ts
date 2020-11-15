import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UploadRoutingModule } from './upload-routing.module';
import { UploadComponent } from './page/upload/upload.component';
import { ViewComponent } from './page/view/view.component';
import { ReactiveFormsModule } from '@angular/forms';
import { FileLoaderComponent } from './component/file-loader/file-loader.component';
import { PatchViewerComponent } from './component/patch-viewer/patch-viewer.component';
import { GaugeModule } from 'angular-gauge';
import { OnIndicatorComponent } from './component/on-indicator/on-indicator.component';


@NgModule({
  declarations: [UploadComponent, ViewComponent, FileLoaderComponent, PatchViewerComponent, OnIndicatorComponent],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    UploadRoutingModule,
    GaugeModule.forRoot(),
  ]
})
export class UploadModule {
}
