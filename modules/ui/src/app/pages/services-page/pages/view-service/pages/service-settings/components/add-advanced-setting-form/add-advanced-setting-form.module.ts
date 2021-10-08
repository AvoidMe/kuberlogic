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

import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { AddAdvancedSettingFormComponent } from './add-advanced-setting-form.component';

@NgModule({
    declarations: [AddAdvancedSettingFormComponent],
    exports: [
        AddAdvancedSettingFormComponent
    ],
    imports: [
        CommonModule,
        MatFormFieldModule,
        MatInputModule,
        ReactiveFormsModule,
        MatButtonModule
    ]
})
export class AddAdvancedSettingFormModule { }