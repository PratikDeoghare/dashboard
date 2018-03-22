import { Observable } from 'rxjs/Observable';
import { ApiService } from './../core/services/api/api.service';
import { DatacenterService } from './../core/services/datacenter/datacenter.service';
import { SlimLoadingBarModule } from 'ng2-slim-loading-bar';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ParticlesModule } from 'angular-particle';
import { RouterTestingModule } from '@angular/router/testing';
import { Router, ActivatedRoute } from '@angular/router';

import { By } from '@angular/platform-browser';
import { TestBed, async, ComponentFixture, inject, fakeAsync, tick } from '@angular/core/testing';
import { RouterStub, ActivatedRouteStub } from './../testing/router-stubs';
import { AuthMockService } from '../testing/services/auth-mock.service';

import { Auth } from './../core/services/index';
import { DashboardComponent } from './dashboard.component';
import { ApiMockService } from '../testing/services/api-mock.service';
import { DatacenterMockService } from '../testing/services/datacenter-mock.service';

const modules: any[] = [
    BrowserModule,
    RouterTestingModule,
    BrowserAnimationsModule,
    ParticlesModule,
    SlimLoadingBarModule.forRoot(),
];

describe('DashboardComponent', () => {
    let fixture: ComponentFixture<DashboardComponent>;
    let component: DashboardComponent;
    let authService: AuthMockService;
    let router: Router;
    let apiService: ApiService;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [
                ...modules,
            ],
            declarations: [
                DashboardComponent
            ],
            providers: [
                { provide: Router, useClass: RouterStub },
                { provide: Auth, useClass: AuthMockService },
                { provide: ApiService, useClass: ApiMockService },
                { provide: DatacenterService, useClass: DatacenterMockService },
                { provide: ActivatedRoute, useClass: ActivatedRouteStub }
            ],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(DashboardComponent);
        component = fixture.componentInstance;

        authService = fixture.debugElement.injector.get(Auth) as any;
        router = fixture.debugElement.injector.get(Router);
        apiService = fixture.debugElement.injector.get(ApiService);
    });

    it('should create the cmp', async(() => {
        expect(component).toBeTruthy();
    }));
});
