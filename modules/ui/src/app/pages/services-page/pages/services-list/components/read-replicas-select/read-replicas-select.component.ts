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

import { Component, Input, OnInit } from '@angular/core';
import { ServiceModel, ServiceModelStatus } from '@models/service.model';
import { MessagesService } from '@services/messages.service';
import { ServicesPageService } from '@services/services-page.service';
import { throwError } from 'rxjs';
import { catchError, tap } from 'rxjs/operators';

@Component({
    selector: 'kl-read-replicas-select',
    templateUrl: './read-replicas-select.component.html',
    styleUrls: ['./read-replicas-select.component.scss'],
})
export class ReadReplicasSelectComponent implements OnInit {
    @Input() service: ServiceModel | undefined;
    selected: number | undefined;
    readonly replicas = [...Array(11).keys()];
    readonly serviceStatus = ServiceModelStatus;
    constructor(
        private messagesService: MessagesService,
        private pageService: ServicesPageService
    ) {}

    ngOnInit(): void {
        this.selected = this.service?.replicas;
    }

    onSubmit(value: any): void {
        if (this.service) {
            this.pageService.editService(`${this.service.ns}:${this.service.name}`, {
                replicas: value,
                name: this.service.name,
                ns: this.service.ns,
                type: this.service.type,
            }).pipe(
                catchError((err) => {
                    this.messagesService.error('An error occurred, please try again later');
                    return throwError(err);
                }),
                tap(() => {
                    this.messagesService.success('Service was successfully updated');
                }),
            ).subscribe();
        }
    }
}