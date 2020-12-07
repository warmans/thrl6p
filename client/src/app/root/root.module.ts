import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { RootRoutingModule } from './root-routing.module';
import { RootComponent } from './component/root/root.component';
import { PatchServiceAPIClientModule } from '../lib/api-client/services/patch-service';
import { HttpClientModule } from '@angular/common/http';
import { CoreModule } from '../module/core/core.module';

@NgModule({
  declarations: [
    RootComponent,
  ],
  imports: [
    CoreModule,
    BrowserModule,
    RootRoutingModule,
    HttpClientModule,
    PatchServiceAPIClientModule.forRoot()
  ],
  bootstrap: [RootComponent],
  providers: [],
})
export class RootModule {
}
