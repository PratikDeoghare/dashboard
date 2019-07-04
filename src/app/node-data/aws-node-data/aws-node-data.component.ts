import {Component, Input, OnDestroy, OnInit} from '@angular/core';
import {FormArray, FormControl, FormGroup, Validators} from '@angular/forms';
import {Subscription} from 'rxjs';

import {WizardService} from '../../core/services';
import {NodeDataService} from '../../core/services/node-data/node-data.service';
import {CloudSpec} from '../../shared/entity/ClusterEntity';
import {NodeInstanceFlavor, NodeInstanceFlavors} from '../../shared/model/NodeProviderConstants';
import {NodeData, NodeProviderData} from '../../shared/model/NodeSpecChange';

@Component({
  selector: 'kubermatic-aws-node-data',
  templateUrl: './aws-node-data.component.html',
})

export class AWSNodeDataComponent implements OnInit, OnDestroy {
  @Input() cloudSpec: CloudSpec;
  @Input() nodeData: NodeData;
  @Input() clusterId: string;

  instanceTypes: NodeInstanceFlavor[] = NodeInstanceFlavors.AWS;
  diskTypes: string[] = ['standard', 'gp2', 'io1', 'sc1', 'st1'];
  awsNodeForm: FormGroup;
  tags: FormArray;
  hideOptional = true;
  private subscriptions: Subscription[] = [];

  constructor(private addNodeService: NodeDataService, private wizardService: WizardService) {}

  ngOnInit(): void {
    const tagList = new FormArray([]);
    for (const i in this.nodeData.spec.cloud.aws.tags) {
      if (this.nodeData.spec.cloud.aws.tags.hasOwnProperty(i)) {
        tagList.push(new FormGroup({
          key: new FormControl(i),
          value: new FormControl(this.nodeData.spec.cloud.aws.tags[i]),
        }));
      }
    }

    this.awsNodeForm = new FormGroup({
      type: new FormControl(this.nodeData.spec.cloud.aws.instanceType, Validators.required),
      disk_size: new FormControl(this.nodeData.spec.cloud.aws.diskSize, Validators.required),
      disk_type: new FormControl(this.nodeData.spec.cloud.aws.volumeType, Validators.required),
      ami: new FormControl(this.nodeData.spec.cloud.aws.ami),
      tags: tagList,
    });

    if (this.nodeData.spec.cloud.aws.instanceType === '') {
      this.awsNodeForm.controls.type.setValue(this.instanceTypes[0].id);
    }

    this.subscriptions.push(this.awsNodeForm.valueChanges.subscribe(() => {
      this.addNodeService.changeNodeProviderData(this.getNodeProviderData());
    }));

    this.subscriptions.push(this.wizardService.clusterSettingsFormViewChanged$.subscribe((data) => {
      this.hideOptional = data.hideOptional;
    }));

    this.addNodeService.changeNodeProviderData(this.getNodeProviderData());
  }

  isInWizard(): boolean {
    return !this.clusterId || this.clusterId.length === 0;
  }

  getNodeProviderData(): NodeProviderData {
    const tagMap = {};
    for (const i in this.awsNodeForm.controls.tags.value) {
      if (this.awsNodeForm.controls.tags.value[i].key !== '' && this.awsNodeForm.controls.tags.value[i].value !== '') {
        tagMap[this.awsNodeForm.controls.tags.value[i].key] = this.awsNodeForm.controls.tags.value[i].value;
      }
    }

    return {
      spec: {
        aws: {
          instanceType: this.awsNodeForm.controls.type.value,
          diskSize: this.awsNodeForm.controls.disk_size.value,
          ami: this.awsNodeForm.controls.ami.value,
          tags: tagMap,
          volumeType: this.awsNodeForm.controls.disk_type.value,
        },
      },
      valid: this.awsNodeForm.valid,
    };
  }

  getTagForm(form): any {
    return form.get('tags').controls;
  }

  addTag(): void {
    this.tags = this.awsNodeForm.get('tags') as FormArray;
    this.tags.push(new FormGroup({
      key: new FormControl(''),
      value: new FormControl(''),
    }));
  }

  deleteTag(index: number): void {
    const arrayControl = this.awsNodeForm.get('tags') as FormArray;
    arrayControl.removeAt(index);
  }

  ngOnDestroy(): void {
    for (const sub of this.subscriptions) {
      if (sub) {
        sub.unsubscribe();
      }
    }
  }
}
