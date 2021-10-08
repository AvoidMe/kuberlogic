/*
 * CloudLinux Software Inc 2019-2021 All Rights Reserved
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ServicesPageComponent } from '@app/pages/services-page/services-page.component';

const routes: Routes = [
    {
        path: '',
        component: ServicesPageComponent,
        children: [
            {
                path: '',
                loadChildren: () => import('./pages/services-list/services-list.module')
                    .then((mod) => mod.ServicesListModule),
            },
            {
                path: 'create',
                loadChildren: () => import('./pages/create-service/create-service.module')
                    .then((mod) => mod.CreateServiceModule),
            },
            {
                path: ':id',
                loadChildren: () => import('./pages/view-service/view-service.module')
                    .then((mod) => mod.ViewServiceModule),
            }
        ]
    }
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule],
})
export class ServicesPageRoutingModule { }