import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { RootRoutingModule } from './root-routing.module';
import { RootComponent } from './component/root/root.component';
import { PatchServiceAPIClientModule } from '../lib/api-client/services/patch-service';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    RootComponent
  ],
  imports: [
    BrowserModule,
    RootRoutingModule,
    HttpClientModule,
    PatchServiceAPIClientModule.forRoot()
  ],
  providers: [],
  bootstrap: [RootComponent]
})
export class RootModule {
}
