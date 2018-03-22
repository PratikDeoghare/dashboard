import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterConnectComponent } from './cluster-connect.component';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material';
import {clusterFake} from '../../../testing/fake-data/cluster.fake';
import {BrowserModule} from '@angular/platform-browser';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {SharedModule} from '../../../shared/shared.module';
import {RouterTestingModule} from '../../../testing/router-stubs';
import {HttpClientModule} from '@angular/common/http';
import {NgReduxTestingModule} from '@angular-redux/store/lib/testing/ng-redux-testing.module';
import {SlimLoadingBarModule} from 'ng2-slim-loading-bar';
import {datacentersFake} from '../../../testing/fake-data/datacenter.fake';

const modules: any[] = [
  BrowserModule,
  HttpClientModule,
  BrowserAnimationsModule,
  SlimLoadingBarModule.forRoot(),
  RouterTestingModule,
  NgReduxTestingModule,
  SharedModule
];

describe('ClusterConnectComponent', () => {
  let component: ClusterConnectComponent;
  let fixture: ComponentFixture<ClusterConnectComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [
        ...modules,
      ],
      declarations: [ ClusterConnectComponent ],
      providers: [
        { provide: MAT_DIALOG_DATA, useValue: { cluster: clusterFake } },
        { provide: MatDialogRef, useValue: {} },
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ClusterConnectComponent);
    component = fixture.componentInstance;
    component.datacenter = datacentersFake[0];
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
